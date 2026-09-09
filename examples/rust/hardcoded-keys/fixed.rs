// Corrigé : Clés cryptographiques codées en dur — CWE-798 (Use of Hard-coded Credentials)
// La clé est désormais chargée depuis une variable d'environnement injectée de façon
// sécurisée (ou un gestionnaire de secrets en production), et jamais depuis une
// valeur littérale du code source.

use aes_gcm::aead::{Aead, KeyInit};
use aes_gcm::{Aes256Gcm, Nonce};
use std::env;

/// Charge la clé de chiffrement depuis l'environnement.
/// Échoue explicitement si le secret est absent — jamais de valeur par défaut.
fn charger_cle_chiffrement() -> Result<[u8; 32], String> {
    let cle_b64 = env::var("ENCRYPTION_KEY")
        .map_err(|_| "ENCRYPTION_KEY manquant dans l'environnement".to_string())?;

    let cle = base64::decode(&cle_b64).map_err(|_| "ENCRYPTION_KEY mal formé".to_string())?;
    cle.try_into()
        .map_err(|_| "ENCRYPTION_KEY doit faire 32 octets (AES-256)".to_string())
}

fn chiffrer(cle: &[u8; 32], donnees: &[u8], nonce: &[u8; 12]) -> Result<Vec<u8>, String> {
    let cipher = Aes256Gcm::new_from_slice(cle).map_err(|_| "clé invalide".to_string())?;
    cipher
        .encrypt(Nonce::from_slice(nonce), donnees)
        .map_err(|_| "échec du chiffrement".to_string())
}

fn main() -> Result<(), String> {
    let cle = charger_cle_chiffrement()?;
    let chiffre = chiffrer(&cle, b"donnees confidentielles", &[0u8; 12])?;
    println!("Chiffré : {:02x?}", chiffre);
    Ok(())
}
