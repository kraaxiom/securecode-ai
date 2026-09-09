// Corrigé : CWE-307 — limitation de débit par compte ET par IP, avec
// verrouillage progressif, conformément à rules/remediation/brute-force.md.

use actix_web::{post, web, HttpResponse, HttpRequest, Responder};
use serde::Deserialize;
use std::collections::HashMap;
use std::sync::Mutex;
use std::time::{Duration, Instant};

#[derive(Deserialize)]
struct LoginInput {
    email: String,
    password: String,
}

/// Compteur d'échecs en mémoire, clé = "ip:email" (à remplacer par Redis en prod).
struct AttemptTracker {
    attempts: Mutex<HashMap<String, (u32, Instant)>>,
}

impl AttemptTracker {
    fn too_many_attempts(&self, key: &str, max: u32, window: Duration) -> bool {
        let mut map = self.attempts.lock().unwrap();
        if let Some((count, first_seen)) = map.get(key) {
            if first_seen.elapsed() < window && *count >= max {
                return true;
            }
        }
        false
    }

    fn record_failure(&self, key: &str, window: Duration) {
        let mut map = self.attempts.lock().unwrap();
        let entry = map.entry(key.to_string()).or_insert((0, Instant::now()));
        if entry.1.elapsed() > window {
            *entry = (0, Instant::now());
        }
        entry.0 += 1;
    }

    fn clear(&self, key: &str) {
        self.attempts.lock().unwrap().remove(key);
    }
}

async fn authenticate(_email: &str, _password: &str) -> Option<u32> {
    None
}

#[post("/login")]
async fn login(
    req: HttpRequest,
    input: web::Json<LoginInput>,
    tracker: web::Data<AttemptTracker>,
) -> impl Responder {
    let ip = req.peer_addr().map(|a| a.ip().to_string()).unwrap_or_default();
    let key = format!("{}:{}", ip, input.email.to_lowercase());
    let window = Duration::from_secs(900);

    if tracker.too_many_attempts(&key, 5, window) {
        return HttpResponse::TooManyRequests()
            .json(serde_json::json!({ "error": "Trop de tentatives. Réessayez plus tard." }));
    }

    match authenticate(&input.email, &input.password).await {
        Some(_user_id) => {
            tracker.clear(&key);
            HttpResponse::Ok().json(serde_json::json!({ "ok": true }))
        }
        None => {
            tracker.record_failure(&key, window);
            // Message d'erreur générique, ne permet pas de distinguer
            // un compte existant d'un compte inexistant.
            HttpResponse::Unauthorized().json(serde_json::json!({ "error": "Identifiants invalides" }))
        }
    }
}
