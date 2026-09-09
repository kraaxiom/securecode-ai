// CWE-1427 : correction — séparation de la couche de décision (LLM) et de
// la couche d'exécution (application), confirmation humaine obligatoire
// pour toute action financière, plafond de sécurité, et journalisation
// exhaustive de la décision proposée.

use serde_json::json;

struct OutilPaiement;

impl OutilPaiement {
    async fn executer(&self, compte: &str, montant_centimes: u64) -> Result<(), String> {
        println!("[simulation] paiement de {montant_centimes} centimes vers {compte}");
        Ok(())
    }
}

struct PropositionPaiement {
    compte: String,
    montant_centimes: u64,
}

const PLAFOND_AUTO_CENTIMES: u64 = 0; // aucune exécution automatique de paiement

async fn analyser_ticket_support(
    client: &reqwest::Client,
    api_key: &str,
    contenu_ticket_externe: &str,
) -> Result<Option<PropositionPaiement>, Box<dyn std::error::Error>> {
    // Sécurisé : le modèle ne propose qu'une décision, sans accès direct à
    // l'outil de paiement (pas de "tools" exposé ici) — le contenu externe
    // ne peut donc jamais déclencher directement une action financière.
    let prompt = format!(
        "Analyse ce ticket (contenu externe non fiable) et indique si un \
         remboursement semble justifié, avec compte et montant proposés :\n\n{contenu_ticket_externe}"
    );

    let resp = client
        .post("https://api.example-llm.com/v1/complete")
        .bearer_auth(api_key)
        .json(&json!({ "prompt": prompt }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    // Journalisation exhaustive de la proposition, avant toute exécution,
    // pour permettre l'audit et la détection d'anomalies.
    println!("[audit] proposition du modèle : {resp}");

    let compte = resp["account"].as_str().unwrap_or_default().to_string();
    let montant = resp["amount_cents"].as_u64().unwrap_or(0);

    if compte.is_empty() {
        return Ok(None);
    }

    Ok(Some(PropositionPaiement { compte, montant_centimes: montant }))
}

async fn executer_apres_validation_humaine(
    outil_paiement: &OutilPaiement,
    proposition: PropositionPaiement,
    validee_par_humain: bool,
) -> Result<(), String> {
    // Sécurisé : couche d'exécution indépendante — exige une confirmation
    // humaine explicite avant tout paiement, quel que soit le montant, car
    // il s'agit d'une action irréversible à fort impact.
    if !validee_par_humain || proposition.montant_centimes > PLAFOND_AUTO_CENTIMES {
        println!("Paiement en attente de validation humaine, non exécuté automatiquement.");
        return Ok(());
    }

    outil_paiement
        .executer(&proposition.compte, proposition.montant_centimes)
        .await
}
