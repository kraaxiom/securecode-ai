// Corrigé : Server-Side Request Forgery (SSRF) classique — CWE-918 (Server-Side Request Forgery)
// La destination est validée contre une whitelist explicite de domaines métier, le schéma
// est restreint à HTTPS, le DNS est résolu puis l'adresse IP obtenue est vérifiée comme
// non privée/loopback/link-local avant l'émission de la requête, et les redirections
// automatiques sont désactivées.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use std::net::IpAddr;
use url::Url;

const ALLOWED_HOSTS: &[&str] = &["api.partenaire.example.com"];

#[derive(Deserialize)]
struct CallbackRequest {
    callback_url: String,
}

fn is_forbidden_ip(ip: &IpAddr) -> bool {
    match ip {
        IpAddr::V4(v4) => {
            v4.is_private() || v4.is_loopback() || v4.is_link_local() || v4.is_broadcast()
        }
        IpAddr::V6(v6) => v6.is_loopback() || v6.is_unspecified(),
    }
}

async fn fetch_callback(payload: web::Json<CallbackRequest>) -> impl Responder {
    let parsed = match Url::parse(&payload.callback_url) {
        Ok(u) => u,
        Err(_) => return HttpResponse::BadRequest().body("URL invalide"),
    };

    let host = match parsed.host_str() {
        Some(h) => h,
        None => return HttpResponse::BadRequest().body("Hôte manquant"),
    };

    // Whitelist stricte de destinations autorisées + schéma restreint à HTTPS.
    if parsed.scheme() != "https" || !ALLOWED_HOSTS.contains(&host) {
        return HttpResponse::BadRequest().body("URL non autorisée");
    }

    // Résolution DNS puis validation de l'IP résultante avant toute connexion.
    let resolved = match tokio::net::lookup_host((host, 443)).await {
        Ok(addrs) => addrs.map(|a| a.ip()).collect::<Vec<_>>(),
        Err(_) => return HttpResponse::BadGateway().finish(),
    };
    if resolved.is_empty() || resolved.iter().any(is_forbidden_ip) {
        return HttpResponse::BadRequest().body("Destination interdite");
    }

    let client = match reqwest::Client::builder()
        .redirect(reqwest::redirect::Policy::none()) // pas de redirection automatique
        .timeout(std::time::Duration::from_secs(5))
        .build()
    {
        Ok(c) => c,
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
