// CWE-284 : Improper Access Control
// Le service backend déploie des règles de sécurité Firestore en mode
// "test" (`allow read, write: if true;`) vers l'API Firebase, ce qui rend
// toute la base de données lisible et modifiable sans authentification.

use reqwest::Client;
use serde_json::json;

async fn deployer_regles_firestore(
    client: &Client,
    project_id: &str,
    access_token: &str,
) -> Result<(), reqwest::Error> {
    // Vulnérable : règle "mode test" laissée telle quelle en production,
    // sans vérification d'authentification ni de propriété des données.
    let regles = r#"
        rules_version = '2';
        service cloud.firestore {
          match /databases/{database}/documents {
            match /{document=**} {
              allow read, write: if true;
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
