// CWE-349 : Acceptance of Extraneous Untrusted Data With Trust
// L'agent écrit automatiquement en mémoire persistante tout "fait" qu'il
// déduit de la conversation, sans confirmation utilisateur, et réinjecte
// ce contenu comme contexte de confiance dans les sessions futures.

use serde_json::json;
use std::collections::HashMap;

struct MemoireAgent {
    faits_par_utilisateur: HashMap<String, Vec<String>>,
}

impl MemoireAgent {
    fn ecrire_fait(&mut self, utilisateur_id: &str, fait_deduit: String) {
        // Vulnérable : écriture automatique en mémoire persistante sans
        // validation ni confirmation de l'utilisateur — un fait halluciné
        // ou injecté via la conversation devient durable.
        self.faits_par_utilisateur
            .entry(utilisateur_id.to_string())
            .or_default()
            .push(fait_deduit);
    }

    fn contexte_pour_session(&self, utilisateur_id: &str) -> String {
        // Vulnérable : la mémoire est réinjectée telle quelle comme contexte
        // de confiance, sans revalidation, dans le prompt système suivant.
        self.faits_par_utilisateur
            .get(utilisateur_id)
            .cloned()
            .unwrap_or_default()
            .join("\n")
    }
}

async fn traiter_tour_conversation(
    client: &reqwest::Client,
    api_key: &str,
    memoire: &mut MemoireAgent,
    utilisateur_id: &str,
    message: &str,
) -> Result<String, Box<dyn std::error::Error>> {
    let contexte = memoire.contexte_pour_session(utilisateur_id);
    let prompt = format!("Contexte connu :\n{contexte}\n\nMessage : {message}\n\nExtrait un nouveau fait à mémoriser si pertinent.");

    let resp = client
        .post("https://api.example-llm.com/v1/complete")
        .bearer_auth(api_key)
        .json(&json!({ "prompt": prompt }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    if let Some(nouveau_fait) = resp["extracted_fact"].as_str() {
        memoire.ecrire_fait(utilisateur_id, nouveau_fait.to_string());
    }

    Ok(resp["completion"].as_str().unwrap_or_default().to_string())
}
