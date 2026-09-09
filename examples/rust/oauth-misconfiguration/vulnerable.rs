// Vulnérable : CWE-287 — mauvaise configuration OAuth2
// Le flux d'autorisation ne génère ni ne vérifie de paramètre `state`,
// exposant le callback à une CSRF permettant de lier le compte de
// l'attaquant à la session de la victime.

use actix_web::{get, web, HttpResponse, Responder};
use std::collections::HashMap;

fn build_authorization_url() -> String {
    // Aucun state généré/stocké avant redirection vers le fournisseur.
    "https://provider.example.com/oauth/authorize?client_id=abc&response_type=code".to_string()
}

#[get("/oauth/login")]
async fn oauth_login() -> impl Responder {
    HttpResponse::Found()
        .append_header(("Location", build_authorization_url()))
        .finish()
}

#[get("/oauth/callback")]
async fn oauth_callback(query: web::Query<HashMap<String, String>>) -> impl Responder {
    let code = query.get("code").cloned().unwrap_or_default();
    // Le code est échangé directement, sans jamais vérifier de `state`.
    HttpResponse::Ok().json(serde_json::json!({ "exchanged_code": code }))
}
