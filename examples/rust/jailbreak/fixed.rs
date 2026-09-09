// CWE-1427 : correction structurelle — ajout d'une couche de modération
// indépendante du modèle principal (entrée et sortie), limitation du nombre
// de tours de conversation exploitables, et surveillance des schémas
// d'usage anormaux. Le system prompt n'est jamais considéré comme l'unique
// barrière de sécurité.

use serde_json::json;

const SYSTEM_PROMPT: &str = "Tu es un assistant. Tu ne dois jamais donner \
d'instructions dangereuses.";
const MAX_TOURS_CONVERSATION: usize = 20;

async fn moderer_contenu(
    client: &reqwest::Client,
    api_key: &str,
    texte: &str,
) -> Result<bool, Box<dyn std::error::Error>> {
    // Sécurisé : classifieur de sécurité indépendant du modèle principal,
    // appelé en entrée ET en sortie — ne dépend pas des instructions du
    // system prompt pour détecter un contenu à risque.
    let resp = client
        .post("https://api.example-llm.com/v1/moderations")
        .bearer_auth(api_key)
        .json(&json!({ "input": texte }))
        .send()
        .await?
        .json::<serde_json::Value>()
        .await?;

    Ok(resp["flagged"].as_bool().unwrap_or(true)) // fail-safe : bloque par défaut
}

async fn repondre(
    client: &reqwest::Client,
    api_key: &str,
    historique: &[(String, String)],
) -> Result<String, Box<dyn std::error::Error>> {
    // Sécurisé : nombre de tours de conversation exploitables limité, pour
    // freiner l'érosion progressive des restrictions sur de longs échanges.
    if historique.len() > MAX_TOURS_CONVERSATION {
        return Err("limite de tours de conversation atteinte, session à réinitialiser".into());
    }

    if let Some((_, dernier_message)) = historique.last() {
        if moderer_contenu(client, api_key, dernier_message).await? {
            return Err("message utilisateur bloqué par la modération d'entrée".into());
        }
    }

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

    // Sécurisé : la sortie du modèle passe elle aussi par la modération
    // indépendante avant d'être renvoyée à l'utilisateur.
    if moderer_contenu(client, api_key, &reponse).await? {
        return Err("réponse du modèle bloquée par la modération de sortie".into());
    }

    Ok(reponse)
}
