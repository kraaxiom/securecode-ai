// Vulnérable : Utilisation de SHA-1 — CWE-328 (Use of Weak Hash)
// SHA-1 est cryptographiquement cassé pour la résistance aux collisions depuis
// l'attaque "SHAttered" (2017). L'utiliser pour hacher un mot de passe avant
// stockage ne fournit plus aucune garantie de sécurité.

use sha1::{Digest, Sha1};

/// Hache le mot de passe utilisateur avec SHA-1 avant stockage en base — sans sel.
fn hacher_mot_de_passe(mot_de_passe: &str) -> String {
    let mut hasher = Sha1::new();
    hasher.update(mot_de_passe.as_bytes());
    format!("{:x}", hasher.finalize())
}

fn main() {
    let hash = hacher_mot_de_passe("motdepasse_utilisateur");
    println!("Hash SHA-1 stocké en base : {}", hash);
}
