// Corrigé : CWE-347 — vérification de signature active via jsonwebtoken avec
// une liste explicite d'algorithmes forts, excluant structurellement `none`,
// conformément à rules/remediation/jwt-none.md.

use actix_web::{get, HttpRequest, HttpResponse, Responder};
use jsonwebtoken::{decode, Algorithm, DecodingKey, Validation};
use serde::Deserialize;

#[derive(Deserialize)]
struct Claims {
    role: String,
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

    let secret = std::env::var("JWT_SECRET").expect("JWT_SECRET manquant");
    let key = DecodingKey::from_secret(secret.as_bytes());

    // La bibliothèque jsonwebtoken exige une vérification de signature et
    // une liste explicite d'algorithmes : `none` ne peut structurellement
    // pas être passé ici.
    let validation = Validation::new(Algorithm::HS256);

    match decode::<Claims>(&token, &key, &validation) {
        Ok(data) if data.claims.role == "admin" => HttpResponse::Ok().json(serde_json::json!({ "ok": true })),
        _ => HttpResponse::Forbidden().finish(),
    }
}
