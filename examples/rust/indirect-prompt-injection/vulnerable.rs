// CWE-1427 : Improper Neutralization of Input Used for LLM Prompting
// L'agent récupère le contenu d'une page web fournie par l'utilisateur et
// l'injecte tel quel dans le contexte du modèle, sans marquage de
// provenance ni délimitation, alors qu'il dispose en parallèle d'outils
// actifs (envoi d'e-mail).

use serde_json::json;

struct OutilEnvoiEmail;

impl OutilEnvoiEmail {
    async fn envoyer(&self, destinataire: &str, corps: &str) {
        println!("[simulation] e-mail envoyé à {destinataire} : {corps}");
    }
}

async fn resumer_page_et_repondre(
    client: &reqwest::Client,
    api_key: &str,
    url: &str,
    outil_email: &OutilEnvoiEmail,
) -> Result<(), Box<dyn std::error::Error>> {
    let contenu_page = client.get(url).send().await?.text().await?;

    // Vulnérable : contenu externe non fiable injecté directement dans le
    // contexte, sans balise de provenance, alors que l'agent dispose
    // simultanément d'un outil à fort impact (envoi d'e-mail).
    let prompt = format!(
        "Résume cette page et envoie le résumé par e-mail si pertinent :\n\n{contenu_page}"
    );

    let resp = client
        .post("https://api.example-llm.com/v1/agent/run")
        .bearer_auth(api_key)
        .json(&json!({ "prompt": prompt, "tools": ["send_email"] }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    if let Some(email_a_envoyer) = resp["tool_call"]["send_email"].as_object() {
        // Vulnérable : l'action déclenchée par le modèle après lecture du
        // contenu externe est exécutée directement, sans validation
        // applicative ni confirmation humaine.
        let dest = email_a_envoyer["to"].as_str().unwrap_or_default();
        let corps = email_a_envoyer["body"].as_str().unwrap_or_default();
        outil_email.envoyer(dest, corps).await;
    }

    Ok(())
}
