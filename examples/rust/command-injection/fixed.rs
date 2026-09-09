// Correction CWE-78 — exécution directe du binaire avec arguments passés
// en tableau distinct (pas de shell), et validation stricte de l'entrée
// comme adresse IP avant tout passage à l'exécutable.

use actix_web::{web, HttpResponse, Responder};
use std::net::IpAddr;
use std::process::Command;
use std::str::FromStr;

pub async fn ping_host(query: web::Query<PingQuery>) -> impl Responder {
    let host = &query.host;

    // Sécurisé : validation liste blanche (format IP) avant exécution
    let ip: IpAddr = match IpAddr::from_str(host) {
        Ok(ip) => ip,
        Err(_) => return HttpResponse::BadRequest().body("Hôte invalide"),
    };

    // Sécurisé : pas de shell, arguments passés en tableau distinct
    let output = Command::new("ping")
        .arg("-c")
        .arg("3")
        .arg(ip.to_string())
        .output()
        .expect("échec d'exécution");

    HttpResponse::Ok().body(String::from_utf8_lossy(&output.stdout).to_string())
}

#[derive(serde::Deserialize)]
pub struct PingQuery {
    pub host: String,
}
