// Vulnérable : CWE-330 — identifiant de session prévisible
// L'identifiant de session est dérivé de valeurs connues (ID utilisateur,
// horodatage) via un hash faible, le rendant devinable ou reconstructible
// par un attaquant sans connaître le mot de passe.

use actix_web::{post, web, HttpResponse, Responder};
use std::time::{SystemTime, UNIX_EPOCH};

fn generate_session_id(user_id: u32) -> String {
    let timestamp = SystemTime::now().duration_since(UNIX_EPOCH).unwrap().as_secs();
    // Concaténation/hash de valeurs prévisibles — pas un CSPRNG dédié.
    format!("{:x}", md5_like_hash(&format!("{}-{}", user_id, timestamp)))
}

fn md5_like_hash(input: &str) -> u64 {
    // Simule un hash faible et déterministe (illustration du pattern vulnérable).
    input.bytes().fold(0u64, |acc, b| acc.wrapping_mul(31).wrapping_add(b as u64))
}

#[post("/login/{user_id}")]
async fn login(path: web::Path<u32>) -> impl Responder {
    let session_id = generate_session_id(path.into_inner());
    HttpResponse::Ok()
        .cookie(actix_web::cookie::Cookie::new("session_id", session_id))
        .finish()
}
