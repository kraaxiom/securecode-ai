// Corrigé : CWE-330 — identifiant de session généré exclusivement via un
// générateur cryptographiquement sûr, avec au moins 128 bits d'entropie,
// conformément à rules/remediation/session-prediction.md.

use actix_web::cookie::{Cookie, SameSite};
use actix_web::{post, web, HttpResponse, Responder};
use rand::RngCore;

fn generate_session_id() -> String {
    // CSPRNG dédié : 256 bits d'entropie, aucune dérivation de valeur prévisible.
    let mut bytes = [0u8; 32];
    rand::thread_rng().fill_bytes(&mut bytes);
    hex::encode(bytes)
}

#[post("/login/{user_id}")]
async fn login(_path: web::Path<u32>) -> impl Responder {
    let session_id = generate_session_id();

    // Régénération systématique à chaque élévation de privilège (ici, connexion),
    // et cookie marqué HttpOnly/Secure/SameSite.
    let cookie = Cookie::build("session_id", session_id)
        .http_only(true)
        .secure(true)
        .same_site(SameSite::Strict)
        .finish();

    HttpResponse::Ok().cookie(cookie).finish()
}
