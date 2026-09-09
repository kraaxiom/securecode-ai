// Vulnérable : Générateur aléatoire non cryptographique — CWE-338 (Use of Cryptographically Weak PRNG)
// `rand::thread_rng` combiné à des distributions simples reste un PRNG statistique
// prévisible dans certains contextes ; ici l'exemple utilise une graine fixe et un
// générateur non-CSPRNG explicite, inadapté à la génération de tokens de sécurité.

use rand::rngs::SmallRng;
use rand::{Rng, SeedableRng};

/// Génère un token de réinitialisation de mot de passe — PRNG non cryptographique, graine fixe.
fn generer_token_reinitialisation() -> String {
    // Graine fixe et générateur statistique : la séquence est reproductible/prévisible.
    let mut rng = SmallRng::seed_from_u64(42);
    (0..32)
        .map(|_| format!("{:x}", rng.gen_range(0..16)))
        .collect()
}

fn main() {
    let token = generer_token_reinitialisation();
    println!("Token de réinitialisation : {}", token);
}
