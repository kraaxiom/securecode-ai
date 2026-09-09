// Corrigé : Générateur aléatoire non cryptographique — CWE-338 (Use of Cryptographically Weak PRNG)
// Le token est désormais généré via `OsRng`, un générateur cryptographiquement sûr
// (CSPRNG) puisant son entropie dans le système d'exploitation, sans graine fixe.

use rand::rngs::OsRng;
use rand::RngCore;

/// Génère un token de réinitialisation de mot de passe avec un CSPRNG (128 bits d'entropie).
fn generer_token_reinitialisation() -> String {
    let mut octets = [0u8; 16]; // 16 octets = 128 bits d'entropie minimale
    OsRng.fill_bytes(&mut octets);
    octets.iter().map(|b| format!("{:02x}", b)).collect()
}

fn main() {
    let token = generer_token_reinitialisation();
    println!("Token de réinitialisation : {}", token);
}
