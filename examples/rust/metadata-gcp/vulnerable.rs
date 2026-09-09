// Vulnérable : SSRF vers le service de métadonnées GCP — CWE-918 (Server-Side Request Forgery)
// La fonction de récupération d'URL ne filtre ni la plage link-local ni le nom d'hôte de
// métadonnées, et laisse l'utilisateur définir les en-têtes transmis. Hébergée sur GCE/Cloud
// Run avec un compte de service attaché, elle peut être détournée pour atteindre
// `metadata.google.internal` en injectant l'en-tête `Metadata-Flavor: Google` requis.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;
use std::collections::HashMap;

#[derive(Deserialize)]
struct FetchRequest {
    url: String,
    headers: HashMap<String, String>,
}

async fn fetch_resource(payload: web::Json<FetchRequest>) -> impl Responder {
    let client = reqwest::Client::new();
    let mut request = client.get(&payload.url);

    for (name, value) in &payload.headers {
        request = request.header(name, value);
    }

    match request.send().await {
        Ok(resp) => {
            let body = resp.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(_) => HttpResponse::BadGateway().finish(),
    }
}
