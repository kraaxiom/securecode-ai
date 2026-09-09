// Vulnérable : CWE-326 — JWT signé avec un secret faible
// Le secret de signature HMAC est court et codé en dur dans le source,
// permettant une attaque par dictionnaire/force brute hors ligne.

use actix_web::{post, HttpResponse, Responder};
use jsonwebtoken::{encode, EncodingKey, Header};
use serde::Serialize;

#[derive(Serialize)]
struct Claims {
    sub: String,
    role: String,
}

// Secret court, en dur dans le code source versionné.
const JWT_SECRET: &str = "secret";

#[post("/login/issue-token")]
async fn issue_token() -> impl Responder {
    let claims = Claims { sub: "user-1".to_string(), role: "user".to_string() };
    let token = encode(
        &Header::default(),
        &claims,
        &EncodingKey::from_secret(JWT_SECRET.as_bytes()),
    )
    .unwrap();

    HttpResponse::Ok().json(serde_json::json!({ "token": token }))
}
