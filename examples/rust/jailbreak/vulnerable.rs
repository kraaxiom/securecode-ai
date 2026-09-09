// CWE-1427 : Improper Neutralization of Input Used for LLM Prompting
// Faiblesse structurelle : le chatbot public s'appuie uniquement sur le
// system prompt du modèle comme garde-fou, sans couche de modération
// indépendante, sans limite de tours de conversation, et sans supervision
// des schémas d'usage anormaux. Aucun payload de contournement n'est
// démontré ici — seule l'absence de défense en profondeur est illustrée.

use serde_json::json;

const SYSTEM_PROMPT: &str = "Tu es un assistant. Tu ne dois jamais donner \
d'instructions dangereuses.";

async fn repondre(
    client: &reqwest::Client,
    api_key: &str,
    historique: &[(String, String)], // (role, contenu), sans limite de taille
) -> Result<String, Box<dyn std::error::Error>> {
    // Vulnérable : aucune limite sur le nombre de tours de conversation —
    // un historique long permet une érosion progressive des restrictions
    // sans qu'aucun contrôle applicatif n'intervienne.
    let mut messages = vec![json!({ "role": "system", "content": SYSTEM_PROMPT })];
    for (role, contenu) in historique {
        messages.push(json!({ "role": role, "content": contenu }));
    }

    let resp = client
        .post("https://api.example-llm.com/v1/chat/completions")
        .bearer_auth(api_key)
        .json(&json!({ "messages": messages }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    let reponse = resp["choices"][0]["message"]["content"]
        .as_str()
        .unwrap_or_default()
        .to_string();

    // Vulnérable : la réponse du modèle est renvoyée telle quelle, sans
    // classifieur de sécurité indépendant en sortie.
    Ok(reponse)
}
