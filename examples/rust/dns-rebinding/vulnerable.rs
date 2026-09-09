// Vulnérable : DNS Rebinding (contournement de filtre SSRF) — CWE-918 (Server-Side Request Forgery)
// L'hôte est résolu une première fois pour valider qu'il ne pointe pas vers une IP privée,
// mais la requête HTTP réelle est ensuite effectuée séparément : le client HTTP re-résout
// le DNS à ce moment-là. Si l'attaquant contrôle un domaine dont la réponse DNS change entre
// les deux résolutions (TOCTOU), la validation initiale peut être contournée.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use std::net::IpAddr;

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
        Some(h) => h,
        None => return HttpResponse::BadRequest().finish(),
    };

    // Résolution DNS n°1 : utilisée uniquement pour la validation.
    let resolved = tokio::net::lookup_host((host, 443)).await;
    let ok = matches!(resolved, Ok(mut it) if it.all(|a| !is_private_ip(&a.ip())));
    if !ok {
        return HttpResponse::BadRequest().body("Destination interdite");
    }

    // Résolution DNS n°2 (implicite) : reqwest re-résout `host` au moment de la connexion.
    // Si le TTL DNS a expiré et que la réponse a changé entre-temps, la connexion peut
    // aboutir vers une IP interne malgré la validation précédente.
    match reqwest::get(parsed.as_str()).await {
        Ok(resp) => {
            let body = resp.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(_) => HttpResponse::BadGateway().finish(),
    }
}
