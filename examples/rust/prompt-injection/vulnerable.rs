// CWE-1427 : Improper Neutralization of Input Used for LLM Prompting
// Le prompt système et l'entrée utilisateur sont concaténés dans une seule
// chaîne de texte brute envoyée au modèle, sans séparation structurelle des
// rôles. Le modèle ne peut donc pas distinguer "instruction développeur" et
// "instruction utilisateur", et la sortie déclenche une action sensible
// sans validation applicative intermédiaire.

use serde_json::json;

const SYSTEM_PROMPT: &str = "Tu es un assistant support client. \
Résume la demande et réponds poliment.";

async fn traiter_message_utilisateur(
    client: &reqwest::Client,
    api_key: &str,
    message_utilisateur: &str,
) -> Result<String, reqwest::Error> {
    // Vulnérable : simple concaténation de chaînes au lieu d'utiliser les
    // rôles structurés de l'API (system/user/assistant).
    let prompt_complet = format!("{SYSTEM_PROMPT}\n\nMessage client : {message_utilisateur}");

    let resp = client
        .post("https://api.example-llm.com/v1/complete")
        .bearer_auth(api_key)
        .json(&json!({ "prompt": prompt_complet }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    let reponse = resp["completion"].as_str().unwrap_or_default().to_string();

    // Vulnérable : la sortie du modèle est utilisée directement pour
    // déclencher une action (envoi au client) sans aucune validation.
    Ok(reponse)
}
