// CWE-1427 : Improper Neutralization of Input Used for LLM Prompting
// L'agent dispose d'un outil à fort impact (paiement) directement
// invocable par le modèle, sans validation applicative intermédiaire ni
// confirmation humaine, y compris lorsqu'il vient de traiter du contenu
// externe non fiable dans la même session.

use serde_json::json;

struct OutilPaiement;

impl OutilPaiement {
    async fn executer(&self, compte: &str, montant_centimes: u64) -> Result<(), String> {
        println!("[simulation] paiement de {montant_centimes} centimes vers {compte}");
        Ok(())
    }
}

async fn traiter_ticket_support(
    client: &reqwest::Client,
    api_key: &str,
    contenu_ticket_externe: &str, // provient d'un ticket soumis par un tiers
    outil_paiement: &OutilPaiement,
) -> Result<(), Box<dyn std::error::Error>> {
    let prompt = format!(
        "Analyse ce ticket et effectue un remboursement si justifié :\n\n{contenu_ticket_externe}"
    );

    let resp = client
        .post("https://api.example-llm.com/v1/agent/run")
        .bearer_auth(api_key)
        .json(&json!({ "prompt": prompt, "tools": ["process_payment"] }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    if let Some(action) = resp["tool_call"]["process_payment"].as_object() {
        let compte = action["account"].as_str().unwrap_or_default();
        let montant = action["amount_cents"].as_u64().unwrap_or(0);

        // Vulnérable : action financière irréversible exécutée directement
        // sur simple décision du modèle, sans confirmation humaine ni
        // plafond, alors que la décision découle d'un contenu externe.
        outil_paiement.executer(compte, montant).await.ok();
    }

    Ok(())
}
