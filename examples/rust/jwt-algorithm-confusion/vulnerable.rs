// Vulnérable : CWE-347 — confusion d'algorithme JWT
// Le décodage du token laisse la bibliothèque déduire l'algorithme depuis
// l'en-tête `alg` du token, sans imposer côté serveur l'algorithme attendu
// (RS256). Un attaquant peut forger un token HS256 signé avec la clé
// publique RSA réutilisée comme secret HMAC.

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

    // Validation par défaut : accepte plusieurs familles d'algorithmes,
    // la clé publique RSA peut être détournée comme secret HMAC.
    let key = DecodingKey::from_rsa_pem(RSA_PUBLIC_KEY_PEM.as_bytes()).unwrap();
    let validation = Validation::default(); // n'impose pas explicitement RS256

    match decode::<Claims>(&token, &key, &validation) {
        Ok(data) if data.claims.role == "admin" => HttpResponse::Ok().json(serde_json::json!({ "ok": true })),
        _ => HttpResponse::Forbidden().finish(),
    }
}
