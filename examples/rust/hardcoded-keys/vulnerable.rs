// Vulnérable : Clés cryptographiques codées en dur — CWE-798 (Use of Hard-coded Credentials)
// La clé de chiffrement est écrite en clair dans le code source. Quiconque a accès
// au dépôt, à l'historique Git ou au binaire compilé peut l'extraire et déchiffrer
// toutes les données protégées par cette clé.

use aes_gcm::aead::{Aead, KeyInit};
use aes_gcm::{Aes256Gcm, Nonce};

// Clé codée en dur : identique en dev, staging et production, jamais tournée.
const ENCRYPTION_KEY: &[u8; 32] = b"changeme1234567890changeme123456";

fn chiffrer(donnees: &[u8], nonce: &[u8; 12]) -> Vec<u8> {
    let cipher = Aes256Gcm::new_from_slice(ENCRYPTION_KEY).expect("clé invalide");
    cipher
        .encrypt(Nonce::from_slice(nonce), donnees)
        .expect("échec du chiffrement")
}

fn main() {
    let chiffre = chiffrer(b"donnees confidentielles", &[0u8; 12]);
    println!("Chiffré : {:02x?}", chiffre);
}
