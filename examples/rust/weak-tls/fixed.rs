// Corrigé : Configuration TLS faible — CWE-326 (Inadequate Encryption Strength)
// La vérification du certificat et du hostname reste active, et la version TLS
// minimale est fixée à 1.2, conformément à rules/remediation/weak-tls.md.

use reqwest::{tls::Version, Client};

/// Construit un client HTTP avec vérification stricte du certificat et TLS 1.2 minimum.
fn construire_client_http() -> Client {
    Client::builder()
        // La vérification de certificat et de hostname reste active (comportement par défaut,
        // jamais désactivée même en environnement de test).
        .min_tls_version(Version::TLS_1_2)
        .build()
        .expect("échec de construction du client HTTP")
}

#[tokio::main]
async fn main() {
    let client = construire_client_http();
    let _ = client.get("https://api.exemple.local/donnees").send().await;
    println!("Requête envoyée avec vérification TLS stricte (TLS 1.2 minimum)");
}
