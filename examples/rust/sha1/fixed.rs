// Corrigé : Utilisation de SHA-1 — CWE-328 (Use of Weak Hash)
// Le hachage de mot de passe utilise désormais Argon2id (sel automatique, coût
// adaptatif) ; SHA-1 aurait sinon pu être remplacé par SHA-256 pour un usage
// d'intégrité générale non lié aux mots de passe.

use argon2::password_hash::{PasswordHasher, PasswordVerifier, SaltString};
use argon2::{Argon2, PasswordHash};
use rand::rngs::OsRng;

/// Hache le mot de passe utilisateur avec Argon2id avant stockage en base.
fn hacher_mot_de_passe(mot_de_passe: &str) -> Result<String, argon2::password_hash::Error> {
    let sel = SaltString::generate(&mut OsRng);
    let argon2 = Argon2::default(); // Argon2id par défaut
    let hash = argon2.hash_password(mot_de_passe.as_bytes(), &sel)?;
    Ok(hash.to_string())
}

/// Vérifie un mot de passe candidat face au hash stocké.
fn verifier_mot_de_passe(mot_de_passe: &str, hash_stocke: &str) -> Result<bool, argon2::password_hash::Error> {
    let hash = PasswordHash::new(hash_stocke)?;
    Ok(Argon2::default()
        .verify_password(mot_de_passe.as_bytes(), &hash)
        .is_ok())
}

fn main() -> Result<(), argon2::password_hash::Error> {
    let hash = hacher_mot_de_passe("motdepasse_utilisateur")?;
    println!("Hash Argon2id stocké en base : {}", hash);
    println!("Vérification : {}", verifier_mot_de_passe("motdepasse_utilisateur", &hash)?);
    Ok(())
}
