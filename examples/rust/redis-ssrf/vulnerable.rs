// Vulnérable : SSRF ciblant Redis — CWE-918 (Server-Side Request Forgery)
// L'hôte et le port du cache sont entièrement dérivés d'une entrée utilisateur, sans
// whitelist. Un attaquant peut ainsi rediriger la connexion vers n'importe quelle instance
// Redis interne, généralement dépourvue d'authentification, pour lire ou écrire des clés.

use actix_web::{web, HttpResponse, Responder};
use redis::AsyncCommands;
use serde::Deserialize;

#[derive(Deserialize)]
struct CacheLookupRequest {
    cache_host: String,
    cache_port: u16,
    key: String,
}

async fn cache_lookup(payload: web::Json<CacheLookupRequest>) -> impl Responder {
    // L'hôte/port de connexion Redis provient directement de l'entrée utilisateur.
    let url = format!("redis://{}:{}/", payload.cache_host, payload.cache_port);

    let client = match redis::Client::open(url) {
        Ok(c) => c,
        Err(_) => return HttpResponse::BadRequest().finish(),
    };

    let mut conn = match client.get_multiplexed_async_connection().await {
        Ok(c) => c,
        Err(_) => return HttpResponse::BadGateway().finish(),
    };

    let value: Option<String> = conn.get(&payload.key).await.unwrap_or(None);
    HttpResponse::Ok().json(value)
}
