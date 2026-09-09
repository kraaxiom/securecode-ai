// CWE-78 : Improper Neutralization of Special Elements used in an OS Command
// L'entrée utilisateur `host` est concaténée directement dans une chaîne de
// commande shell, puis exécutée via `sh -c`. Un attaquant peut injecter des
// méta-caractères shell pour exécuter des commandes arbitraires.

use actix_web::{web, HttpResponse};
use std::process::Command;

#[derive(serde::Deserialize)]
struct PingQuery {
    host: String,
}

async fn ping_handler(query: web::Query<PingQuery>) -> HttpResponse {
    let host = query.host.clone();

    // Vulnérable : concaténation directe dans une commande interprétée par un shell
    let cmd = format!("ping -c 4 {}", host);
    let output = Command::new("sh")
        .arg("-c")
        .arg(cmd)
        .output();

    match output {
        Ok(o) => HttpResponse::Ok().body(String::from_utf8_lossy(&o.stdout).to_string()),
        Err(_) => HttpResponse::InternalServerError().body("Erreur d'exécution"),
    }
}
