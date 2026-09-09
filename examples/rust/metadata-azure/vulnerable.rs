// Vulnérable : SSRF vers le service de métadonnées Azure (IMDS) — CWE-918 (Server-Side Request Forgery)
// La fonction de récupération d'URL laisse l'utilisateur contrôler à la fois la destination
// et les en-têtes transmis. Hébergée sur une ressource Azure avec identité managée, elle peut
// être détournée pour atteindre l'IMDS Azure (169.254.169.254) en injectant l'en-tête
// `Metadata: true` requis, et récupérer un jeton Azure AD.

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

    // Les en-têtes de la requête sortante sont librement définis par l'utilisateur,
    // ce qui permet d'injecter `Metadata: true` attendu par l'IMDS Azure.
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
