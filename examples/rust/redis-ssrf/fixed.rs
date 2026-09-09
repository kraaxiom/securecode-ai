// Corrigé : SSRF ciblant Redis — CWE-918 (Server-Side Request Forgery)
// L'hôte et le port de l'instance Redis sont fixés côté serveur, jamais dérivés d'une
// entrée utilisateur. La connexion transporte un mot de passe (`requirepass`/ACL) et n'est
// possible qu'avec une clé de cache validée, ce qui empêche toute redirection vers une
// instance Redis interne arbitraire.

use actix_web::{web, HttpResponse, Responder};
use redis::AsyncCommands;
use serde::Deserialize;

// Hôte/port Redis fixés côté serveur, jamais dérivés d'une entrée utilisateur, avec
// authentification obligatoire même en réseau interne.
const REDIS_HOST: &str = "redis-internal.svc.cluster.local";
const REDIS_PORT: u16 = 6379;

#[derive(Deserialize)]
struct CacheLookupRequest {
    key: String,
}

fn is_valid_cache_key(key: &str) -> bool {
    !key.is_empty()
        && key.len() <= 128
        && key
            .chars()
            .all(|c| c.is_ascii_alphanumeric() || c == ':' || c == '-' || c == '_')
}

async fn cache_lookup(payload: web::Json<CacheLookupRequest>) -> impl Responder {
    if !is_valid_cache_key(&payload.key) {
        return HttpResponse::BadRequest().body("Clé de cache invalide");
    }

    let redis_password = std::env::var("REDIS_PASSWORD").unwrap_or_default();
    let url = format!(
        "redis://:{}@{}:{}/",
        redis_password, REDIS_HOST, REDIS_PORT
    );

    let client = match redis::Client::open(url) {
        Ok(c) => c,
        Err(_) => return HttpResponse::InternalServerError().finish(),
    };

    let mut conn = match client.get_multiplexed_async_connection().await {
        Ok(c) => c,
        Err(_) => return HttpResponse::BadGateway().finish(),
    };

    let value: Option<String> = conn.get(&payload.key).await.unwrap_or(None);
    HttpResponse::Ok().json(value)
}
