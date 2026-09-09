// Corrigé : Protocole TLS 1.0 activé — CWE-326 (Inadequate Encryption Strength)
// La version minimale acceptée est désormais TLS 1.2, avec TLS 1.3 en cible
// préférentielle, conformément à rules/remediation/tls1-0.md.

use openssl::ssl::{SslConnector, SslMethod, SslVerifyMode, SslVersion};

/// Construit un connecteur TLS n'acceptant que TLS 1.2 et TLS 1.3.
fn construire_connecteur_tls() -> Result<SslConnector, openssl::error::ErrorStack> {
    let mut builder = SslConnector::builder(SslMethod::tls())?;

    builder.set_min_proto_version(Some(SslVersion::TLS1_2))?;
    builder.set_max_proto_version(Some(SslVersion::TLS1_3))?;
    builder.set_verify(SslVerifyMode::PEER);

    Ok(builder.build())
}

fn main() -> Result<(), openssl::error::ErrorStack> {
    let _connecteur = construire_connecteur_tls()?;
    println!("Connecteur TLS construit (TLS 1.2 minimum, TLS 1.0/1.1 exclus)");
    Ok(())
}
