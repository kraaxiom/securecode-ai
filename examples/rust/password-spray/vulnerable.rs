// Vulnérable : CWE-307 — absence de détection de password spraying
// La limitation des échecs est appliquée uniquement par compte, ce qui ne
// détecte pas un attaquant testant un même mot de passe courant sur de
// nombreux comptes distincts depuis une même source.

use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;
use std::collections::HashMap;
use std::sync::Mutex;

#[derive(Deserialize)]
struct LoginInput {
    email: String,
    password: String,
}

struct AccountAttempts {
    failures: Mutex<HashMap<String, u32>>,
}

async fn authenticate(_email: &str, _password: &str) -> Option<u32> {
    None
}

#[post("/login")]
async fn login(input: web::Json<LoginInput>, tracker: web::Data<AccountAttempts>) -> impl Responder {
    let mut failures = tracker.failures.lock().unwrap();
    // Seule la limite par compte est vérifiée — pas d'agrégation par IP,
    // ce qui laisse passer un spray distribué sur de nombreux comptes.
    let count = failures.entry(input.email.to_lowercase()).or_insert(0);
    if *count >= 5 {
        return HttpResponse::TooManyRequests().finish();
    }

    match authenticate(&input.email, &input.password).await {
        Some(_) => HttpResponse::Ok().json(serde_json::json!({ "ok": true })),
        None => {
            *count += 1;
            HttpResponse::Unauthorized().finish()
        }
    }
}
