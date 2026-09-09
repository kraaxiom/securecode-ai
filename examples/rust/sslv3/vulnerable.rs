// Vulnérable : Protocole SSLv3 activé — CWE-326 (Inadequate Encryption Strength)
// SSLv3 (interdit par la RFC 7568 depuis 2015) est vulnérable à l'attaque POODLE,
// qui permet à un attaquant en position de MITM de déchiffrer des données via un
// padding oracle sur le mode CBC. L'accepter reste une faille critique.

use openssl::ssl::{SslConnector, SslMethod, SslVersion};

/// Construit un connecteur TLS autorisant SSLv3 comme version minimale.
fn construire_connecteur_tls() -> Result<SslConnector, openssl::error::ErrorStack> {
    let mut builder = SslConnector::builder(SslMethod::tls())?;

    // SSLv3 explicitement autorisé comme version minimale : vulnérable à POODLE.
    builder.set_min_proto_version(Some(SslVersion::SSL3))?;

    Ok(builder.build())
}

fn main() -> Result<(), openssl::error::ErrorStack> {
    let _connecteur = construire_connecteur_tls()?;
    println!("Connecteur TLS construit (SSLv3 accepté)");
    Ok(())
}
