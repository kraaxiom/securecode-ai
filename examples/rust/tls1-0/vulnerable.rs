// Vulnérable : Protocole TLS 1.0 activé — CWE-326 (Inadequate Encryption Strength)
// TLS 1.0/1.1 sont officiellement dépréciés par la RFC 8996 (2021), exposés à
// l'attaque BEAST et exclus des référentiels de conformité (PCI-DSS) depuis 2018.

use openssl::ssl::{SslConnector, SslMethod, SslVersion};

/// Construit un connecteur TLS autorisant TLS 1.0 comme version minimale.
fn construire_connecteur_tls() -> Result<SslConnector, openssl::error::ErrorStack> {
    let mut builder = SslConnector::builder(SslMethod::tls())?;

    // TLS 1.0 explicitement autorisé : vulnérable à BEAST, non conforme PCI-DSS.
    builder.set_min_proto_version(Some(SslVersion::TLS1))?;

    Ok(builder.build())
}

fn main() -> Result<(), openssl::error::ErrorStack> {
    let _connecteur = construire_connecteur_tls()?;
    println!("Connecteur TLS construit (TLS 1.0 accepté)");
    Ok(())
}
