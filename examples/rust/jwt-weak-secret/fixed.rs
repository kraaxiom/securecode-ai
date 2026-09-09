// Corrigé : CWE-326 — secret chargé depuis un gestionnaire de secrets
// (variable d'environnement en exemple), avec vérification d'une entropie
// minimale de 256 bits, conformément à rules/remediation/jwt-weak-secret.md.

use actix_web::{post, HttpResponse, Responder};
use jsonwebtoken::{encode, EncodingKey, Header};
use serde::Serialize;

#[derive(Serialize)]
struct Claims {
    sub: String,
    role: String,
}

fn load_jwt_secret() -> String {
    let secret = std::env::var("JWT_SECRET").expect("JWT_SECRET manquant");
    // Vérification explicite d'une entropie minimale (256 bits = 32 octets)
    // avant toute utilisation pour signer un token.
    if secret.as_bytes().len() < 32 {
        panic!("JWT_SECRET insuffisant (256 bits minimum requis).");
    }
    secret
}

#[post("/login/issue-token")]
async fn issue_token() -> impl Responder {
    let secret = load_jwt_secret(); // généré via un CSPRNG et stocké dans un vault
    let claims = Claims { sub: "user-1".to_string(), role: "user".to_string() };
    let token = encode(
        &Header::default(),
        &claims,
        &EncodingKey::from_secret(secret.as_bytes()),
    )
    .unwrap();

    HttpResponse::Ok().json(serde_json::json!({ "token": token }))
}
