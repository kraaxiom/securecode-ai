// CWE-284 : correction — les règles Firestore déployées exigent une
// authentification ET vérifient la propriété de la ressource, avec
// validation de schéma minimale sur l'écriture.

use reqwest::Client;
use serde_json::json;

async fn deployer_regles_firestore(
    client: &Client,
    project_id: &str,
    access_token: &str,
) -> Result<(), reqwest::Error> {
    // Sécurisé : lecture/écriture limitées au propriétaire de la ressource,
    // avec validation minimale du schéma des champs écrits.
    let regles = r#"
        rules_version = '2';
        service cloud.firestore {
          match /databases/{database}/documents {
            match /users/{userId}/documents/{docId} {
              allow read: if request.auth != null && request.auth.uid == userId;
              allow write: if request.auth != null
                            && request.auth.uid == userId
                            && request.resource.data.keys().hasOnly(['title', 'content', 'updatedAt']);
            }
          }
        }
    "#;

    let url = format!(
        "https://firebaserules.googleapis.com/v1/projects/{project_id}/rulesets"
    );

    client
        .post(&url)
        .bearer_auth(access_token)
        .json(&json!({
            "source": {
                "files": [{ "name": "firestore.rules", "content": regles }]
            }
        }))
        .send()
        .await?;

    Ok(())
}
