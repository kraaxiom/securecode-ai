// Corrigé : Blind SSRF — CWE-918 (Server-Side Request Forgery)
// Les mêmes contrôles qu'une SSRF classique sont appliqués avant de lancer la requête
// asynchrone : whitelist de domaines, résolution DNS et validation de l'IP obtenue,
// redirections désactivées. La requête sortante émise par le worker est en outre
// journalisée pour permettre la détection a posteriori d'un scan interne.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use std::net::IpAddr;
use url::Url;

const ALLOWED_HOSTS: &[&str] = &["api.partenaire.example.com"];

#[derive(Deserialize)]
struct WebhookRegistration {
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

async fn validate_destination(raw_url: &str) -> Result<Url, &'static str> {
    let parsed = Url::parse(raw_url).map_err(|_| "URL invalide")?;
    let host = parsed.host_str().ok_or("Hôte manquant")?;

    if parsed.scheme() != "https" || !ALLOWED_HOSTS.contains(&host) {
        return Err("URL non autorisée");
    }

    let resolved = tokio::net::lookup_host((host, 443))
        .await
        .map_err(|_| "Résolution DNS impossible")?
        .map(|a| a.ip())
        .collect::<Vec<_>>();

    if resolved.is_empty() || resolved.iter().any(is_forbidden_ip) {
        return Err("Destination interdite");
    }

    Ok(parsed)
}

async fn register_webhook(payload: web::Json<WebhookRegistration>) -> impl Responder {
    let raw_url = payload.url.clone();

    match validate_destination(&raw_url).await {
        Ok(url) => {
            tokio::spawn(async move {
                // Requête journalisée pour permettre la détection a posteriori d'un scan interne.
                log::info!("Requête sortante webhook validée vers {}", url);
                let client = match reqwest::Client::builder()
                    .redirect(reqwest::redirect::Policy::none())
                    .timeout(std::time::Duration::from_secs(5))
                    .build()
                {
                    Ok(c) => c,
                    Err(_) => return,
                };
                let _ = client.get(url.as_str()).send().await;
            });
            HttpResponse::Accepted().body("Webhook en cours de vérification")
        }
        Err(reason) => HttpResponse::BadRequest().body(reason),
    }
}
