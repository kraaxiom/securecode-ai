// Vulnérable : CWE-347 — acceptation de l'algorithme JWT `none`
// Le token est simplement décodé sans vérification de signature active pour
// prendre une décision d'autorisation, ouvrant la voie à un token forgé
// avec `alg: none` et une signature vide.

use actix_web::{get, HttpRequest, HttpResponse, Responder};
use base64::Engine;
use serde::Deserialize;
use serde_json::Value;

#[derive(Deserialize)]
struct Claims {
    role: String,
}

/// Décodage naïf du payload JWT sans vérification cryptographique de la
/// signature — équivalent à jwt.decode() sans verify en JS/Python.
fn decode_unverified(token: &str) -> Option<Claims> {
    let parts: Vec<&str> = token.split('.').collect();
    if parts.len() != 3 {
        return None;
    }
    let payload_bytes = base64::engine::general_purpose::URL_SAFE_NO_PAD
        .decode(parts[1])
        .ok()?;
    let value: Value = serde_json::from_slice(&payload_bytes).ok()?;
    serde_json::from_value(value).ok()
}

#[get("/admin/data")]
async fn admin_data(req: HttpRequest) -> impl Responder {
    let token = req
        .headers()
        .get("Authorization")
        .and_then(|h| h.to_str().ok())
        .unwrap_or("")
        .trim_start_matches("Bearer ")
        .to_string();

    // AUCUNE vérification de signature : un token alg:none avec signature
    // vide passe cette étape sans problème.
    match decode_unverified(&token) {
        Some(claims) if claims.role == "admin" => HttpResponse::Ok().json(serde_json::json!({ "ok": true })),
        _ => HttpResponse::Forbidden().finish(),
    }
}
