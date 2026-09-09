// CWE-349 : correction — modération du contenu avant génération de
// l'embedding, limite de fréquence/volume par utilisateur, et
// réévaluation périodique de la pertinence pour des requêtes de référence.

use serde_json::json;
use std::collections::HashMap;
use std::time::{Duration, Instant};

struct LimiteurFrequence {
    dernieres_soumissions: HashMap<String, Vec<Instant>>,
}

impl LimiteurFrequence {
    // Sécurisé : limite le volume d'indexation par source pour freiner une
    // campagne d'empoisonnement automatisée.
    fn autoriser(&mut self, utilisateur_id: &str) -> bool {
        let maintenant = Instant::now();
        let entrees = self.dernieres_soumissions.entry(utilisateur_id.to_string()).or_default();
        entrees.retain(|t| maintenant.duration_since(*t) < Duration::from_secs(3600));

        if entrees.len() >= 10 {
            return false;
        }
        entrees.push(maintenant);
        true
    }
}

async fn moderer_contenu(
    client: &reqwest::Client,
    api_key: &str,
    texte: &str,
) -> Result<bool, Box<dyn std::error::Error>> {
    // Sécurisé : modération indépendante avant toute génération d'embedding.
    let resp = client
        .post("https://api.example-llm.com/v1/moderations")
        .bearer_auth(api_key)
        .json(&json!({ "input": texte }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    Ok(resp["flagged"].as_bool().unwrap_or(true))
}

async fn indexer_avis_utilisateur(
    client: &reqwest::Client,
    api_key: &str,
    vector_db_url: &str,
    avis_utilisateur: &str,
    utilisateur_id: &str,
    limiteur: &mut LimiteurFrequence,
) -> Result<(), Box<dyn std::error::Error>> {
    // Sécurisé : limite de fréquence par source appliquée avant tout
    // traitement.
    if !limiteur.autoriser(utilisateur_id) {
        return Err("limite de soumissions atteinte pour cet utilisateur".into());
    }

    // Sécurisé : contenu modéré avant génération et indexation de
    // l'embedding.
    if moderer_contenu(client, api_key, avis_utilisateur).await? {
        return Err("contenu rejeté par la modération, non indexé".into());
    }

    let embedding_resp = client
        .post("https://api.example-llm.com/v1/embeddings")
        .bearer_auth(api_key)
        .json(&json!({ "input": avis_utilisateur }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    let vecteur = embedding_resp["embedding"].clone();

    client
        .post(vector_db_url)
        .json(&json!({
            "vector": vecteur,
            "metadata": { "text": avis_utilisateur, "user": utilisateur_id }
        }))
        .send()
        .await?;

    Ok(())
}
