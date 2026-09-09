// Corrigé : SSRF vers le service de métadonnées GCP — CWE-918 (Server-Side Request Forgery)
// La destination est validée contre une whitelist, les plages privées/link-local sont
// bloquées, et le nom d'hôte de métadonnées `metadata.google.internal` est explicitement
// refusé en plus de son IP link-local. Aucun en-tête fourni par l'utilisateur n'est transmis,
// empêchant l'injection de `Metadata-Flavor: Google`.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use std::net::IpAddr;
use url::Url;

const ALLOWED_HOSTS: &[&str] = &["api.partenaire.example.com"];
// Endpoints sensibles de métadonnées GCP, bloqués explicitement en plus du filtrage link-local.
const METADATA_IP: &str = "169.254.169.254";
const METADATA_HOSTNAME: &str = "metadata.google.internal";

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

    if host == METADATA_IP
        || host.eq_ignore_ascii_case(METADATA_HOSTNAME)
        || parsed.scheme() != "https"
        || !ALLOWED_HOSTS.contains(&host)
    {
        return HttpResponse::BadRequest().body("URL non autorisée");
    }

    let resolved = match tokio::net::lookup_host((host, 443)).await {
        Ok(addrs) => addrs.map(|a| a.ip()).collect::<Vec<_>>(),
        Err(_) => return HttpResponse::BadGateway().finish(),
    };
    if resolved.is_empty() || resolved.iter().any(is_forbidden_ip) {
        return HttpResponse::BadRequest().body("Destination interdite");
    }

    // Aucun en-tête fourni par l'utilisateur n'est transmis à la requête sortante.
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
