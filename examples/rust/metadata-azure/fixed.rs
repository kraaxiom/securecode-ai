// Corrigé : SSRF vers le service de métadonnées Azure (IMDS) — CWE-918 (Server-Side Request Forgery)
// La destination est validée contre une whitelist et les plages privées/link-local (dont
// l'IMDS Azure) sont bloquées explicitement. Les en-têtes HTTP transmis à la requête
// sortante ne sont plus dérivés d'une entrée utilisateur : l'utilisateur ne peut donc plus
// injecter l'en-tête `Metadata: true` attendu par l'IMDS.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use std::net::IpAddr;
use url::Url;

const ALLOWED_HOSTS: &[&str] = &["api.partenaire.example.com"];
// Endpoint sensible IMDS Azure, bloqué explicitement en plus du filtrage des plages link-local.
const METADATA_IP: &str = "169.254.169.254";

#[derive(Deserialize)]
struct FetchRequest {
    url: String,
}

fn is_forbidden_ip(ip: &IpAddr) -> bool {
    match ip {
        IpAddr::V4(v4) => {
            v4.is_private() || v4.is_loopback() || v4.is_link_local() || v4.is_broadcast()
        }
        IpAddr::V6(v6) => v6.is_loopback() || v6.is_unspecified(),
    }
}

async fn fetch_resource(payload: web::Json<FetchRequest>) -> impl Responder {
    let parsed = match Url::parse(&payload.url) {
        Ok(u) => u,
        Err(_) => return HttpResponse::BadRequest().finish(),
    };
    let host = match parsed.host_str() {
        Some(h) => h,
        None => return HttpResponse::BadRequest().finish(),
    };

    if host == METADATA_IP || parsed.scheme() != "https" || !ALLOWED_HOSTS.contains(&host) {
        return HttpResponse::BadRequest().body("URL non autorisée");
    }

    let resolved = match tokio::net::lookup_host((host, 443)).await {
        Ok(addrs) => addrs.map(|a| a.ip()).collect::<Vec<_>>(),
        Err(_) => return HttpResponse::BadGateway().finish(),
    };
    if resolved.is_empty() || resolved.iter().any(is_forbidden_ip) {
        return HttpResponse::BadRequest().body("Destination interdite");
    }

    // Aucun en-tête fourni par l'utilisateur n'est transmis à la requête sortante :
    // impossible d'injecter `Metadata: true` ou tout autre en-tête sensible.
    let client = match reqwest::Client::builder()
        .redirect(reqwest::redirect::Policy::none())
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
