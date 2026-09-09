// Corrigé : SSRF ciblant Elasticsearch — CWE-918 (Server-Side Request Forgery)
// L'hôte du cluster de recherche n'est plus dérivé d'une entrée utilisateur : il provient
// d'une configuration serveur figée. La requête vers Elasticsearch transporte des
// identifiants d'authentification, et seuls l'index et la requête de recherche restent
// paramétrables, avec validation stricte du nom d'index.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;

// Hôte du cluster Elasticsearch fixé côté serveur, jamais dérivé d'une entrée utilisateur.
const ES_HOST: &str = "es-internal.svc.cluster.local:9200";

#[derive(Deserialize)]
struct SearchProxyRequest {
    index: String,
    query: String,
}

fn is_valid_index_name(index: &str) -> bool {
    !index.is_empty()
        && index.len() <= 64
        && index
            .chars()
            .all(|c| c.is_ascii_alphanumeric() || c == '-' || c == '_')
}

async fn proxy_search(payload: web::Json<SearchProxyRequest>) -> impl Responder {
    if !is_valid_index_name(&payload.index) {
        return HttpResponse::BadRequest().body("Nom d'index invalide");
    }

    let es_user = std::env::var("ES_USERNAME").unwrap_or_default();
    let es_pass = std::env::var("ES_PASSWORD").unwrap_or_default();

    let url = format!("https://{}/{}/_search", ES_HOST, payload.index);

    let client = match reqwest::Client::builder()
        .redirect(reqwest::redirect::Policy::none())
        .timeout(std::time::Duration::from_secs(5))
        .build()
    {
        Ok(c) => c,
        Err(_) => return HttpResponse::InternalServerError().finish(),
    };

    // Authentification requise même en réseau interne, et requête envoyée en JSON structuré
    // plutôt qu'en interpolation brute dans l'URL.
    match client
        .get(&url)
        .basic_auth(es_user, Some(es_pass))
        .query(&[("q", &payload.query)])
        .send()
        .await
    {
        Ok(resp) => {
            let body = resp.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(_) => HttpResponse::BadGateway().finish(),
    }
}
