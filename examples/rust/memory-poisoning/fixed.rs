// CWE-349 : correction — écriture en mémoire persistante uniquement après
// confirmation explicite de l'utilisateur, mémoire cloisonnée par
// utilisateur, contenu réinjecté revalidé, et interface de consultation /
// suppression offerte à l'utilisateur.

use serde_json::json;
use std::collections::HashMap;

#[derive(Clone)]
struct EntreeMemoire {
    contenu: String,
    confirmee_par_utilisateur: bool,
}

struct MemoireAgent {
    // Sécurisé : cloisonnement strict par utilisateur, aucune structure
    // partagée entre tenants.
    faits_par_utilisateur: HashMap<String, Vec<EntreeMemoire>>,
}

impl MemoireAgent {
    fn proposer_fait(&mut self, utilisateur_id: &str, fait_deduit: String) -> usize {
        // Sécurisé : le fait déduit est stocké comme "proposition" non
        // confirmée, jamais réinjecté tant qu'il n'est pas validé.
        let entrees = self.faits_par_utilisateur.entry(utilisateur_id.to_string()).or_default();
        entrees.push(EntreeMemoire { contenu: fait_deduit, confirmee_par_utilisateur: false });
        entrees.len() - 1
    }

    fn confirmer_fait(&mut self, utilisateur_id: &str, index: usize) {
        // Sécurisé : confirmation explicite requise avant que le fait ne
        // devienne durable et réinjectable.
        if let Some(entrees) = self.faits_par_utilisateur.get_mut(utilisateur_id) {
            if let Some(entree) = entrees.get_mut(index) {
                entree.confirmee_par_utilisateur = true;
            }
        }
    }

    fn supprimer_fait(&mut self, utilisateur_id: &str, index: usize) {
        // Sécurisé : l'utilisateur peut consulter et supprimer ses propres
        // entrées mémoire.
        if let Some(entrees) = self.faits_par_utilisateur.get_mut(utilisateur_id) {
            if index < entrees.len() {
                entrees.remove(index);
            }
        }
    }

    fn contexte_pour_session(&self, utilisateur_id: &str) -> String {
        // Sécurisé : seules les entrées confirmées par l'utilisateur sont
        // réinjectées comme contexte.
        self.faits_par_utilisateur
            .get(utilisateur_id)
            .map(|entrees| {
                entrees
                    .iter()
                    .filter(|e| e.confirmee_par_utilisateur)
                    .map(|e| e.contenu.clone())
                    .collect::<Vec<_>>()
                    .join("\n")
            })
            .unwrap_or_default()
    }
}

async fn traiter_tour_conversation(
    client: &reqwest::Client,
    api_key: &str,
    memoire: &mut MemoireAgent,
    utilisateur_id: &str,
    message: &str,
) -> Result<(String, Option<usize>), Box<dyn std::error::Error>> {
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

    let mut index_proposition = None;
    if let Some(nouveau_fait) = resp["extracted_fact"].as_str() {
        // Sécurisé : proposition enregistrée mais non confirmée — l'appelant
        // doit demander une confirmation explicite à l'utilisateur avant
        // d'appeler `confirmer_fait`.
        index_proposition = Some(memoire.proposer_fait(utilisateur_id, nouveau_fait.to_string()));
    }

    Ok((resp["completion"].as_str().unwrap_or_default().to_string(), index_proposition))
}
