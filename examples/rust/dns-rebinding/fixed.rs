// Corrigé : DNS Rebinding (contournement de filtre SSRF) — CWE-918 (Server-Side Request Forgery)
// Le DNS est résolu une seule fois, l'IP obtenue est validée, puis la connexion est établie
// directement sur cette IP épinglée (DNS pinning) au lieu de laisser le client HTTP
// re-résoudre le nom séparément. Le nom d'hôte d'origine reste utilisé pour SNI/TLS et
// l'en-tête Host, mais la connexion réseau cible l'adresse déjà validée.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use std::net::{IpAddr, SocketAddr};
use std::sync::Arc;

#[derive(Deserialize)]
struct UrlCheckRequest {
    url: String,
}

fn is_private_ip(ip: &IpAddr) -> bool {
    match ip {
        IpAddr::V4(v4) => v4.is_private() || v4.is_loopback() || v4.is_link_local(),
        IpAddr::V6(v6) => v6.is_loopback(),
    }
}

async fn check_and_fetch(payload: web::Json<UrlCheckRequest>) -> impl Responder {
    let parsed = match url::Url::parse(&payload.url) {
        Ok(u) => u,
        Err(_) => return HttpResponse::BadRequest().finish(),
    };
    let host = match parsed.host_str() {
        Some(h) => h.to_string(),
        None => return HttpResponse::BadRequest().finish(),
    };
    if parsed.scheme() != "https" {
        return HttpResponse::BadRequest().body("Schéma non autorisé");
    }
    let port = parsed.port_or_known_default().unwrap_or(443);

    // Résolution DNS unique : l'IP retenue ici est celle réellement utilisée pour la connexion.
    let addrs: Vec<IpAddr> = match tokio::net::lookup_host((host.as_str(), port)).await {
        Ok(it) => it.map(|a| a.ip()).collect(),
        Err(_) => return HttpResponse::BadGateway().finish(),
    };
    if addrs.is_empty() || addrs.iter().any(is_private_ip) {
        return HttpResponse::BadRequest().body("Destination interdite");
    }
    let pinned_ip = addrs[0];
    let pinned_addr = SocketAddr::new(pinned_ip, port);

    // Épinglage DNS : `resolve` force reqwest à utiliser l'IP déjà validée pour ce host,
    // au lieu de re-résoudre le nom au moment de la connexion.
    let client = match reqwest::Client::builder()
        .resolve(&host, pinned_addr)
        .redirect(reqwest::redirect::Policy::none())
        .timeout(std::time::Duration::from_secs(5))
        .build()
    {
        Ok(c) => Arc::new(c),
        Err(_) => return HttpResponse::InternalServerError().finish(),
    };

    match client.get(parsed.as_str()).send().await {
        Ok(resp) => {
            let body = resp.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(_) => HttpResponse::BadGateway().finish(),
    }
}
