// Vulnérable : CWE-287 — contournement du MFA
// Un token pleinement privilégié est émis dès la validation du mot de passe,
// avant toute vérification du second facteur, qui n'est qu'une étape
// facultative en apparence côté client.

use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct LoginInput {
    email: String,
    password: String,
}

struct User {
    id: u32,
    mfa_enabled: bool,
}

async fn authenticate(_email: &str, _password: &str) -> Option<User> {
    Some(User { id: 1, mfa_enabled: true })
}

fn issue_full_token(user_id: u32) -> String {
    format!("full-token-{}", user_id)
}

#[post("/login")]
async fn login(input: web::Json<LoginInput>) -> impl Responder {
    match authenticate(&input.email, &input.password).await {
        // Token complet émis immédiatement, indépendamment de mfa_enabled.
        Some(user) => HttpResponse::Ok().json(serde_json::json!({ "token": issue_full_token(user.id) })),
        None => HttpResponse::Unauthorized().finish(),
    }
}
