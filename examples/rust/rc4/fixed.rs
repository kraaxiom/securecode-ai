// Corrigé : Utilisation de RC4 — CWE-327 (Use of a Broken or Risky Cryptographic Algorithm)
// RC4 est remplacé par AES-256-GCM, un chiffrement authentifié avec nonce unique
// généré aléatoirement à chaque opération, conformément à rules/remediation/rc4.md.

use aes_gcm::aead::{Aead, KeyInit, OsRng};
use aes_gcm::{Aes256Gcm, Nonce};
use rand::RngCore;

/// Chiffre un flux de données avec AES-256-GCM (chiffrement authentifié).
fn chiffrer_flux(cle: &[u8; 32], donnees: &[u8]) -> (Vec<u8>, [u8; 12]) {
    let cipher = Aes256Gcm::new_from_slice(cle).expect("clé invalide");

    let mut nonce_bytes = [0u8; 12];
    OsRng.fill_bytes(&mut nonce_bytes);
    let nonce = Nonce::from_slice(&nonce_bytes);

    let ciphertext = cipher
        .encrypt(nonce, donnees)
        .expect("échec du chiffrement AES-256-GCM");

    (ciphertext, nonce_bytes)
}

fn main() {
    let cle: [u8; 32] = rand::random();
    let (chiffre, nonce) = chiffrer_flux(&cle, b"donnees a proteger");
    println!("Flux chiffré (AES-256-GCM) : {:02x?}, nonce : {:02x?}", chiffre, nonce);
}
