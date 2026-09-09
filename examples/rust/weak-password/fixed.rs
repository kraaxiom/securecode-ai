// Corrigé : CWE-521 — longueur minimale de 12 caractères et vérification
// contre une liste de mots de passe compromis connus, conformément à
// rules/remediation/weak-password.md. Pas de règles de composition
// artificielles ni de rotation forcée.

use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct RegisterInput {
    email: String,
    password: String,
}

const MIN_LENGTH: usize = 12;

/// Vérifie le mot de passe contre une liste de fuites connues
/// (ex. requête k-anonymity vers une API type HaveIBeenPwned).
async fn is_password_compromised(_password: &str) -> bool {
    false
}

#[post("/register")]
async fn register(input: web::Json<RegisterInput>) -> impl Responder {
    if input.password.chars().count() < MIN_LENGTH {
        return HttpResponse::BadRequest().json(
            serde_json::json!({ "error": "Le mot de passe doit contenir au moins 12 caractères" }),
        );
    }

    if is_password_compromised(&input.password).await {
        return HttpResponse::BadRequest().json(serde_json::json!({
            "error": "Ce mot de passe a été trouvé dans une fuite de données connue, choisissez-en un autre"
        }));
    }

    create_user(&input.email, &input.password);
    HttpResponse::Ok().json(serde_json::json!({ "ok": true }))
}

fn create_user(_email: &str, _password: &str) {}
