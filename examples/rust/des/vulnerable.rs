// Vulnérable : Utilisation de DES / 3DES — CWE-327 (Use of a Broken or Risky Cryptographic Algorithm)
// DES (56 bits effectifs) et 3DES (vulnérable à Sweet32 sur ses blocs de 64 bits)
// n'offrent plus de garantie de confidentialité sérieuse. Leur présence pour
// chiffrer des données sensibles doit être considérée comme une rupture de sécurité.

use des::cipher::{BlockEncrypt, KeyInit};
use des::TdesEde3;

/// Chiffre un bloc de données avec 3DES — algorithme obsolète, cassable via Sweet32.
fn chiffrer_donnees_sensibles(cle: &[u8; 24], bloc: &mut [u8; 8]) {
    // Aucune vérification de fraîcheur de clé, aucun mode authentifié : simple ECB/bloc brut.
    let cipher = TdesEde3::new_from_slice(cle).expect("clé invalide");
    let mut generic_block = (*bloc).into();
    cipher.encrypt_block(&mut generic_block);
    bloc.copy_from_slice(&generic_block);
}

fn main() {
    let cle: [u8; 24] = *b"cledemo1234567890abcdefg";
    let mut donnees = *b"secret!!";
    chiffrer_donnees_sensibles(&cle, &mut donnees);
    println!("Bloc chiffré (3DES) : {:02x?}", donnees);
}
