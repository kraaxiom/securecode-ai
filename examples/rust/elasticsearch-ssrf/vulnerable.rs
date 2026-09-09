// Vulnérable : SSRF ciblant Elasticsearch — CWE-918 (Server-Side Request Forgery)
// Le service applicatif construit l'URL du cluster Elasticsearch à partir de l'hôte fourni
// par l'utilisateur, sans whitelist. Un attaquant peut ainsi rediriger le proxy de recherche
// vers n'importe quel hôte interne, y compris une instance Elasticsearch non authentifiée,
// pour lister les index ou exfiltrer des données.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct SearchProxyRequest {
    es_host: String,
    index: String,
    query: String,
}

async fn proxy_search(payload: web::Json<SearchProxyRequest>) -> impl Responder {
    // L'hôte du cluster de recherche est entièrement déterminé par l'utilisateur.
    let url = format!(
        "http://{}/{}/_search?q={}",
        payload.es_host, payload.index, payload.query
    );

    match reqwest::get(&url).await {
        Ok(resp) => {
            let body = resp.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(_) => HttpResponse::BadGateway().finish(),
    }
}
