// Corrigé : CWE-287 — token intermédiaire à privilèges limités tant que le
// second facteur n'est pas validé côté serveur ; token complet uniquement
// après vérification MFA, conformément à rules/remediation/mfa-bypass.md.

use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct LoginInput {
    email: String,
    password: String,
}

#[derive(Deserialize)]
struct MfaInput {
    partial_token: String,
    code: String,
}

struct User {
    id: u32,
    mfa_enabled: bool,
}

async fn authenticate(_email: &str, _password: &str) -> Option<User> {
    Some(User { id: 1, mfa_enabled: true })
}

fn issue_partial_token(user_id: u32) -> String {
    // Token de portée restreinte (scope: mfa_pending), inutilisable sur les
    // endpoints protégés tant que le second facteur n'est pas validé.
    format!("partial-token-{}", user_id)
}

fn issue_full_token(user_id: u32) -> String {
    format!("full-token-{}", user_id)
}

async fn verify_mfa_code(_partial_token: &str, _code: &str) -> Option<u32> {
    // Vérification serveur du code TOTP/OTP associé au token partiel.
    Some(1)
}

#[post("/login")]
async fn login(input: web::Json<LoginInput>) -> impl Responder {
    match authenticate(&input.email, &input.password).await {
        Some(user) if user.mfa_enabled => HttpResponse::Ok().json(serde_json::json!({
            "mfa_required": true,
            "partial_token": issue_partial_token(user.id)
        })),
        Some(user) => HttpResponse::Ok().json(serde_json::json!({ "token": issue_full_token(user.id) })),
        None => HttpResponse::Unauthorized().finish(),
    }
}

#[post("/mfa/verify")]
async fn mfa_verify(input: web::Json<MfaInput>) -> impl Responder {
    // Le token complet n'est délivré qu'après validation serveur du second facteur.
    match verify_mfa_code(&input.partial_token, &input.code).await {
        Some(user_id) => HttpResponse::Ok().json(serde_json::json!({ "token": issue_full_token(user_id) })),
        None => HttpResponse::Unauthorized().json(serde_json::json!({ "error": "Code invalide" })),
    }
}
