// CWE-78 — Injection de commande OS (OS Command Injection)
// L'hôte fourni par l'utilisateur est concaténé dans une chaîne passée
// à un shell, sans validation ni séparation des arguments.

use actix_web::{web, HttpResponse, Responder};
use std::process::Command;

pub async fn ping_host(query: web::Query<PingQuery>) -> impl Responder {
    let host = &query.host;

    // Vulnérable : passage par le shell avec concaténation de l'entrée utilisateur
    let cmd = format!("ping -c 3 {}", host);
    let output = Command::new("sh")
        .arg("-c")
        .arg(cmd)
        .output()
        .expect("échec d'exécution");

    HttpResponse::Ok().body(String::from_utf8_lossy(&output.stdout).to_string())
}

#[derive(serde::Deserialize)]
pub struct PingQuery {
    pub host: String,
}
