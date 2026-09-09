// Vulnérable : Utilisation de MD5 — CWE-328 (Use of Weak Hash)
// MD5 est cassé depuis 2004 (collisions générables en quelques secondes) et
// n'offre aucune résistance à la force brute moderne. L'utiliser pour hacher
// un mot de passe avant stockage supprime toute garantie de sécurité.

use md5::{Digest, Md5};

/// Hache le mot de passe utilisateur avec MD5 avant stockage en base — sans sel.
fn hacher_mot_de_passe(mot_de_passe: &str) -> String {
    let mut hasher = Md5::new();
    hasher.update(mot_de_passe.as_bytes());
    format!("{:x}", hasher.finalize())
}

fn main() {
    let hash = hacher_mot_de_passe("motdepasse_utilisateur");
    println!("Hash MD5 stocké en base : {}", hash);
}
