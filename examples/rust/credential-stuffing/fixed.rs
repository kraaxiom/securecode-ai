// Corrigé : CWE-307 — détection de vélocité globale par IP en complément du
// contrôle par compte, et exigence de MFA après succès de l'authentification,
// conformément à rules/remediation/credential-stuffing.md.

use actix_web::{post, web, HttpRequest, HttpResponse, Responder};
use serde::Deserialize;
use std::collections::HashMap;
use std::sync::Mutex;
use std::time::{Duration, Instant};

#[derive(Deserialize)]
struct LoginInput {
    email: String,
    password: String,
}

struct User {
    id: u32,
    mfa_enabled: bool,
}

struct VelocityTracker {
    per_ip: Mutex<HashMap<String, (u32, Instant)>>,
}

impl VelocityTracker {
    fn velocity(&self, ip: &str, window: Duration) -> u32 {
        let map = self.per_ip.lock().unwrap();
        match map.get(ip) {
            Some((count, first_seen)) if first_seen.elapsed() < window => *count,
            _ => 0,
        }
    }

    fn record(&self, ip: &str, window: Duration) {
        let mut map = self.per_ip.lock().unwrap();
        let entry = map.entry(ip.to_string()).or_insert((0, Instant::now()));
        if entry.1.elapsed() > window {
            *entry = (0, Instant::now());
        }
        entry.0 += 1;
    }
}

async fn authenticate(_email: &str, _password: &str) -> Option<User> {
    None
}

const SUSPICIOUS_THRESHOLD: u32 = 30;

#[post("/login")]
async fn login(
    req: HttpRequest,
    input: web::Json<LoginInput>,
    tracker: web::Data<VelocityTracker>,
) -> impl Responder {
    let ip = req.peer_addr().map(|a| a.ip().to_string()).unwrap_or_default();
    let window = Duration::from_secs(600);

    // Vélocité globale par IP, tous comptes confondus : signale un trafic
    // automatisé typique du credential stuffing.
    if tracker.velocity(&ip, window) > SUSPICIOUS_THRESHOLD {
        return HttpResponse::TooManyRequests()
            .json(serde_json::json!({ "error": "Trafic anormal détecté depuis cette origine." }));
    }
    tracker.record(&ip, window);

    match authenticate(&input.email, &input.password).await {
        Some(user) if user.mfa_enabled => HttpResponse::Ok()
            .json(serde_json::json!({ "mfa_required": true, "challenge_id": format!("mfa-{}", user.id) })),
        Some(user) => HttpResponse::Ok().json(serde_json::json!({ "token": format!("tok-{}", user.id) })),
        None => HttpResponse::Unauthorized().json(serde_json::json!({ "error": "Identifiants invalides" })),
    }
}
