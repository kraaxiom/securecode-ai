// Corrigé : CWE-307 — agrégation globale des échecs par IP en complément de
// la limite par compte, permettant de détecter un spray distribué sur de
// nombreux comptes depuis une même source, conformément à
// rules/remediation/password-spray.md.

use actix_web::{post, web, HttpRequest, HttpResponse, Responder};
use serde::Deserialize;
use std::collections::HashMap;
use std::sync::Mutex;

#[derive(Deserialize)]
struct LoginInput {
    email: String,
    password: String,
}

struct AttemptTracker {
    per_account: Mutex<HashMap<String, u32>>,
    per_ip: Mutex<HashMap<String, u32>>,
}

async fn authenticate(_email: &str, _password: &str) -> Option<u32> {
    None
}

const ACCOUNT_LIMIT: u32 = 5;
const IP_LIMIT: u32 = 50; // volume tous comptes confondus, révélateur d'un spray

#[post("/login")]
async fn login(
    req: HttpRequest,
    input: web::Json<LoginInput>,
    tracker: web::Data<AttemptTracker>,
) -> impl Responder {
    let ip = req.peer_addr().map(|a| a.ip().to_string()).unwrap_or_default();

    let account_attempts = *tracker
        .per_account
        .lock()
        .unwrap()
        .get(&input.email.to_lowercase())
        .unwrap_or(&0);
    let ip_attempts = *tracker.per_ip.lock().unwrap().get(&ip).unwrap_or(&0);

    if account_attempts >= ACCOUNT_LIMIT || ip_attempts >= IP_LIMIT {
        if ip_attempts >= IP_LIMIT {
            flag_suspicious_source(&ip);
        }
        return HttpResponse::TooManyRequests().finish();
    }

    match authenticate(&input.email, &input.password).await {
        Some(_) => HttpResponse::Ok().json(serde_json::json!({ "ok": true })),
        None => {
            *tracker.per_account.lock().unwrap().entry(input.email.to_lowercase()).or_insert(0) += 1;
            *tracker.per_ip.lock().unwrap().entry(ip).or_insert(0) += 1;
            HttpResponse::Unauthorized().finish()
        }
    }
}

fn flag_suspicious_source(_ip: &str) {
    // Route l'alerte vers la supervision de sécurité.
}
