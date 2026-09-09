// Vulnérable : Configuration TLS faible — CWE-326 (Inadequate Encryption Strength)
// La vérification de certificat est désactivée et les suites de chiffrement
// faibles (sans forward secrecy) ne sont pas exclues, exposant le canal à
// l'interception et à la falsification de trafic (MITM).

use reqwest::Client;

/// Construit un client HTTP qui désactive la vérification de certificat TLS.
fn construire_client_http() -> Client {
    Client::builder()
        .danger_accept_invalid_certs(true) // vérification de certificat désactivée
        .danger_accept_invalid_hostnames(true) // vérification du hostname désactivée
        .build()
        .expect("échec de construction du client HTTP")
}

#[tokio::main]
async fn main() {
    let client = construire_client_http();
    let _ = client.get("https://api.exemple.local/donnees").send().await;
    println!("Requête envoyée sans vérification TLS");
}
