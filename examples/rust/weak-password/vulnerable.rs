// Vulnérable : CWE-521 — politique de mot de passe faible
// L'inscription accepte un mot de passe dès 6 caractères, sans vérification
// contre une liste de mots de passe compromis connus.

use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct RegisterInput {
    email: String,
    password: String,
}

#[post("/register")]
async fn register(input: web::Json<RegisterInput>) -> impl Responder {
    // Seule contrainte : longueur minimale insuffisante, aucun contrôle
    // contre les mots de passe compromis connus.
    if input.password.len() < 6 {
        return HttpResponse::BadRequest()
            .json(serde_json::json!({ "error": "Mot de passe trop court" }));
    }

    create_user(&input.email, &input.password);
    HttpResponse::Ok().json(serde_json::json!({ "ok": true }))
}

fn create_user(_email: &str, _password: &str) {}
