// Corrigé : Utilisation de DES / 3DES — CWE-327 (Use of a Broken or Risky Cryptographic Algorithm)
// DES/3DES sont remplacés par AES-256 en mode authentifié (AES-256-GCM), avec un
// nonce unique et aléatoire généré à chaque chiffrement, comme recommandé par
// rules/remediation/des.md.

use aes_gcm::aead::{Aead, KeyInit, OsRng};
use aes_gcm::{Aes256Gcm, Nonce};
use rand::RngCore;

/// Chiffre des données sensibles avec AES-256-GCM (chiffrement authentifié).
fn chiffrer_donnees_sensibles(cle: &[u8; 32], donnees: &[u8]) -> (Vec<u8>, [u8; 12]) {
    let cipher = Aes256Gcm::new_from_slice(cle).expect("clé invalide");

    // Nonce unique et aléatoire à chaque appel : ne jamais le réutiliser avec la même clé.
    let mut nonce_bytes = [0u8; 12];
    OsRng.fill_bytes(&mut nonce_bytes);
    let nonce = Nonce::from_slice(&nonce_bytes);

    let ciphertext = cipher
        .encrypt(nonce, donnees)
        .expect("échec du chiffrement AES-256-GCM");

    (ciphertext, nonce_bytes)
    // Le nonce doit être stocké/transmis avec le texte chiffré : il est nécessaire au déchiffrement.
}

fn main() {
    let cle: [u8; 32] = rand::random();
    let (chiffre, nonce) = chiffrer_donnees_sensibles(&cle, b"donnees sensibles");
    println!("Chiffré (AES-256-GCM) : {:02x?}, nonce : {:02x?}", chiffre, nonce);
}
