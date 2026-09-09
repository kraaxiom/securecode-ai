// Corrigé : Protocole SSLv2 activé — CWE-326 (Inadequate Encryption Strength)
// La configuration TLS impose désormais TLS 1.2 comme version minimale (idéalement
// TLS 1.3), SSLv2 n'étant plus jamais négociable, et la vérification de certificat
// reste active.

use openssl::ssl::{SslConnector, SslMethod, SslVerifyMode, SslVersion};

/// Construit un connecteur TLS n'acceptant que TLS 1.2 et TLS 1.3.
fn construire_connecteur_tls() -> Result<SslConnector, openssl::error::ErrorStack> {
    let mut builder = SslConnector::builder(SslMethod::tls())?;

    // Version minimale forcée à TLS 1.2 : SSLv2/SSLv3/TLS1.0/1.1 sont exclus.
    builder.set_min_proto_version(Some(SslVersion::TLS1_2))?;
    builder.set_max_proto_version(Some(SslVersion::TLS1_3))?;

    // La vérification du certificat serveur reste impérativement active.
    builder.set_verify(SslVerifyMode::PEER);

    Ok(builder.build())
}

fn main() -> Result<(), openssl::error::ErrorStack> {
    let _connecteur = construire_connecteur_tls()?;
    println!("Connecteur TLS construit (TLS 1.2 minimum, vérification active)");
    Ok(())
}
