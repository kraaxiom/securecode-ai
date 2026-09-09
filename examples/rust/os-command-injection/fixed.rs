// CWE-78 : correction — aucune interprétation shell, arguments passés en tableau,
// et validation de l'entrée par liste blanche stricte avant exécution.

use actix_web::{web, HttpResponse};
use std::process::Command;

#[derive(serde::Deserialize)]
struct PingQuery {
    host: String,
}

fn is_valid_host(host: &str) -> bool {
    !host.is_empty()
        && host.len() <= 253
        && host.chars().all(|c| c.is_ascii_alphanumeric() || c == '.' || c == '-')
}

async fn ping_handler(query: web::Query<PingQuery>) -> HttpResponse {
    let host = query.host.trim();

    // Liste blanche stricte : uniquement caractères d'un nom d'hôte/IP valide
    if !is_valid_host(host) {
        return HttpResponse::BadRequest().body("Hôte invalide");
    }

    // Sécurisé : pas de shell, arguments passés séparément (execFile-style)
    let output = Command::new("ping")
        .arg("-c")
        .arg("4")
        .arg(host)
        .output();

    match output {
        Ok(o) => HttpResponse::Ok().body(String::from_utf8_lossy(&o.stdout).to_string()),
        Err(_) => HttpResponse::InternalServerError().body("Erreur d'exécution"),
    }
}
