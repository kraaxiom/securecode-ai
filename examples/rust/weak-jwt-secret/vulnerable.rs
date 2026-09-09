// Vulnérable : Secret JWT faible ou codé en dur — CWE-1391 (Use of Weak Credentials)
// Le secret HMAC utilisé pour signer les JWT est une chaîne courte et prévisible,
// codée en dur dans le source. Un attaquant peut le retrouver par force brute
// hors ligne ou simplement en lisant le code, puis forger des tokens valides.

use jsonwebtoken::{encode, EncodingKey, Header};
use serde::Serialize;

#[derive(Serialize)]
struct Claims {
    sub: String,
    exp: usize,
}

// Secret court et prévisible, identique en dev/staging/prod.
const JWT_SECRET: &str = "secret123";

fn emettre_token(utilisateur_id: &str) -> String {
    let claims = Claims {
        sub: utilisateur_id.to_string(),
        exp: 9999999999,
    };
    encode(
        &Header::default(), // HS256 par défaut
        &claims,
        &EncodingKey::from_secret(JWT_SECRET.as_bytes()),
    )
    .expect("échec de signature du token")
}

fn main() {
    let token = emettre_token("user_42");
    println!("Token émis : {}", token);
}
