// Vulnérable : Utilisation de RC4 — CWE-327 (Use of a Broken or Risky Cryptographic Algorithm)
// RC4 présente des biais statistiques connus dans son flux de sortie, exploitables
// pour récupérer du texte en clair. Interdit dans TLS depuis la RFC 7465, il ne
// doit plus être utilisé pour aucun chiffrement applicatif.

use rc4::{KeyInit, Rc4, StreamCipher};

/// Chiffre un flux de données avec RC4 — chiffrement par flux cassé.
fn chiffrer_flux(cle: &[u8], donnees: &mut [u8]) {
    let mut cipher = Rc4::new(cle.into());
    cipher.apply_keystream(donnees);
}

fn main() {
    let cle = b"cledemo123456";
    let mut donnees = *b"donnees a proteger";
    chiffrer_flux(cle, &mut donnees);
    println!("Flux chiffré (RC4) : {:02x?}", donnees);
}
