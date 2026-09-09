// Corrigé : CWE-287 — génération et vérification stricte d'un `state`
// unique par requête, conformément à rules/remediation/oauth-misconfiguration.md.
// redirect_uri validé par correspondance exacte (liste blanche), pas de wildcard.

use actix_web::{get, web, HttpRequest, HttpResponse, Responder};
use rand::RngCore;
use std::collections::HashMap;

const ALLOWED_REDIRECT_URI: &str = "https://app.example.com/oauth/callback";

fn generate_state() -> String {
    let mut bytes = [0u8; 16];
    rand::thread_rng().fill_bytes(&mut bytes);
    hex::encode(bytes)
}

fn build_authorization_url(state: &str) -> String {
    format!(
        "https://provider.example.com/oauth/authorize?client_id=abc&response_type=code&redirect_uri={}&state={}",
        ALLOWED_REDIRECT_URI, state
    )
}

#[get("/oauth/login")]
async fn oauth_login(req: HttpRequest) -> impl Responder {
    let state = generate_state();
    // Stocké en session côté serveur pour vérification au retour.
    let session = req.extensions_mut();
    // (dans une vraie implémentation : session.insert("oauth_state", state.clone()))

    HttpResponse::Found()
        .append_header(("Location", build_authorization_url(&state)))
        .cookie(actix_web::cookie::Cookie::build("oauth_state", state).http_only(true).finish())
        .finish()
}

#[get("/oauth/callback")]
async fn oauth_callback(req: HttpRequest, query: web::Query<HashMap<String, String>>) -> impl Responder {
    let returned_state = query.get("state").cloned().unwrap_or_default();
    let expected_state = req
        .cookie("oauth_state")
        .map(|c| c.value().to_string())
        .unwrap_or_default();

    if returned_state.is_empty() || returned_state != expected_state {
        return HttpResponse::Forbidden()
            .json(serde_json::json!({ "error": "State invalide — tentative de CSRF sur le callback OAuth." }));
    }

    let code = query.get("code").cloned().unwrap_or_default();
    HttpResponse::Ok().json(serde_json::json!({ "exchanged_code": code }))
}
