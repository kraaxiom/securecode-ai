// Vulnérable : CWE-307 — absence de limitation des tentatives d'authentification
// L'endpoint de connexion n'applique aucun rate limiting ni verrouillage
// progressif, permettant un nombre illimité de tentatives par compte/IP.

use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct LoginInput {
    email: String,
    password: String,
}

async fn authenticate(_email: &str, _password: &str) -> Option<u32> {
    // Simule une vérification de mot de passe.
    None
}

#[post("/login")]
async fn login(input: web::Json<LoginInput>) -> impl Responder {
    // Aucune vérification de compteur d'échecs ni de rate limiting avant
    // de tenter l'authentification.
    match authenticate(&input.email, &input.password).await {
        Some(_user_id) => HttpResponse::Ok().json(serde_json::json!({ "ok": true })),
        None => HttpResponse::Unauthorized().json(serde_json::json!({ "error": "Identifiants invalides" })),
    }
}
