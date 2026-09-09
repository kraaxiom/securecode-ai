// Vulnérable : Protocole SSLv2 activé — CWE-326 (Inadequate Encryption Strength)
// SSLv2 (interdit par la RFC 6176 depuis 2011) présente des failles structurelles
// graves (négociation non authentifiée, MAC faible) exploitées notamment par
// l'attaque DROWN. L'accepter dans la configuration TLS expose à un downgrade.

use openssl::ssl::{SslConnector, SslMethod, SslVerifyMode};

/// Construit un connecteur TLS acceptant SSLv2 — protocole totalement cassé.
fn construire_connecteur_tls() -> SslConnector {
    let mut builder = SslConnector::builder(SslMethod::tls()).expect("init SSL");

    // Aucune restriction de version minimale : SSLv2 reste négociable.
    builder.set_verify(SslVerifyMode::NONE); // la vérification de certificat est aussi désactivée

    builder.build()
}

fn main() {
    let _connecteur = construire_connecteur_tls();
    println!("Connecteur TLS construit (SSLv2 potentiellement négociable)");
}
