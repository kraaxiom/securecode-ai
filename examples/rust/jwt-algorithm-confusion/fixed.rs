// Corrigé : CWE-347 — l'algorithme attendu est figé explicitement à RS256,
// jamais déduit du header du token, empêchant la réutilisation de la clé
// publique RSA comme secret HMAC, conformément à
// rules/remediation/jwt-algorithm-confusion.md.

use actix_web::{get, web, HttpRequest, HttpResponse, Responder};
use jsonwebtoken::{decode, Algorithm, DecodingKey, Validation};
use serde::Deserialize;

#[derive(Deserialize)]
struct Claims {
    sub: String,
    role: String,
}

const RSA_PUBLIC_KEY_PEM: &str = "-----BEGIN PUBLIC KEY-----\n...\n-----END PUBLIC KEY-----";

#[get("/admin/data")]
async fn admin_data(req: HttpRequest) -> impl Responder {
    let token = req
        .headers()
        .get("Authorization")
        .and_then(|h| h.to_str().ok())
        .unwrap_or("")
        .trim_start_matches("Bearer ")
        .to_string();

    let key = DecodingKey::from_rsa_pem(RSA_PUBLIC_KEY_PEM.as_bytes()).unwrap();

    // Liste fermée et explicite : uniquement RS256, exclut HS256 et toute
    // autre famille, quel que soit le header `alg` fourni par le client.
    let mut validation = Validation::new(Algorithm::RS256);
    validation.algorithms = vec![Algorithm::RS256];

    match decode::<Claims>(&token, &key, &validation) {
        Ok(data) if data.claims.role == "admin" => HttpResponse::Ok().json(serde_json::json!({ "ok": true })),
        _ => HttpResponse::Forbidden().finish(),
    }
}
