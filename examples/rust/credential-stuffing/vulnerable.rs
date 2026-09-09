// Vulnérable : CWE-307 — absence de détection de credential stuffing
// L'endpoint de connexion vérifie uniquement les identifiants sans corréler
// le volume de tentatives par source, ni proposer de MFA, laissant passer
// un trafic automatisé testant des couples identifiant/mot de passe fuités.

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
    None
}

#[post("/login")]
async fn login(input: web::Json<LoginInput>) -> impl Responder {
    // Aucune vélocité globale par IP, aucune vérification MFA après succès.
    match authenticate(&input.email, &input.password).await {
        Some(user) => HttpResponse::Ok().json(serde_json::json!({ "token": format!("tok-{}", user.id) })),
        None => HttpResponse::Unauthorized().json(serde_json::json!({ "error": "Identifiants invalides" })),
    }
}
