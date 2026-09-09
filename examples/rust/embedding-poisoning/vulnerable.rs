// CWE-349 : Acceptance of Extraneous Untrusted Data With Trust
// Tout contenu soumis par un utilisateur (description de produit, avis)
// est directement transformé en embedding et indexé dans la base
// vectorielle, sans modération ni limite de fréquence, permettant à un
// attaquant de polluer les résultats de recherche sémantique.

use serde_json::json;

async fn indexer_avis_utilisateur(
    client: &reqwest::Client,
    api_key: &str,
    vector_db_url: &str,
    avis_utilisateur: &str,
    utilisateur_id: &str,
) -> Result<(), Box<dyn std::error::Error>> {
    // Vulnérable : le contenu est directement transformé en embedding et
    // indexé, sans modération préalable ni limite de volume par source.
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
