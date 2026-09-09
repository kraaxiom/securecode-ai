// Corrigé : Secret JWT faible ou codé en dur — CWE-1391 (Use of Weak Credentials)
// Le secret est chargé depuis un gestionnaire de secrets/variable d'environnement,
// avec une vérification de longueur minimale (256 bits), l'algorithme attendu est
// épinglé explicitement à la vérification, et l'expiration reste courte.

use jsonwebtoken::{
    decode, encode, Algorithm, DecodingKey, EncodingKey, Header, Validation,
};
use serde::{Deserialize, Serialize};
use std::env;
use std::time::{SystemTime, UNIX_EPOCH};

#[derive(Serialize, Deserialize)]
struct Claims {
    sub: String,
    exp: usize,
}

/// Charge le secret JWT depuis l'environnement et vérifie son entropie minimale.
fn charger_secret_jwt() -> Result<String, String> {
    let secret = env::var("JWT_SECRET").map_err(|_| "JWT_SECRET manquant".to_string())?;
    if secret.len() < 32 {
        // 32 octets = 256 bits d'entropie minimale exigée.
        return Err("JWT_SECRET doit contenir au moins 256 bits d'entropie".to_string());
    }
    Ok(secret)
}

fn emettre_token(secret: &str, utilisateur_id: &str) -> Result<String, String> {
    let expiration = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs() as usize
        + 15 * 60; // expiration courte : 15 minutes

    let claims = Claims {
        sub: utilisateur_id.to_string(),
        exp: expiration,
    };

    encode(
        &Header::new(Algorithm::HS256),
        &claims,
        &EncodingKey::from_secret(secret.as_bytes()),
    )
    .map_err(|_| "échec de signature du token".to_string())
}

/// Vérifie le token en épinglant strictement l'algorithme attendu (rejette `alg: none`).
fn verifier_token(secret: &str, token: &str) -> Result<Claims, String> {
    let mut validation = Validation::new(Algorithm::HS256);
    validation.algorithms = vec![Algorithm::HS256]; // aucun autre algorithme accepté

    decode::<Claims>(token, &DecodingKey::from_secret(secret.as_bytes()), &validation)
        .map(|data| data.claims)
        .map_err(|_| "token invalide ou expiré".to_string())
}

fn main() -> Result<(), String> {
    let secret = charger_secret_jwt()?;
    let token = emettre_token(&secret, "user_42")?;
    let claims = verifier_token(&secret, &token)?;
    println!("Token émis pour {}", claims.sub);
    Ok(())
}
