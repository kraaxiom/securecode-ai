// Vulnérable : Server-Side Request Forgery (SSRF) classique — CWE-918 (Server-Side Request Forgery)
// L'URL de callback fournie par l'utilisateur est transmise telle quelle au client HTTP,
// sans whitelist de destinations ni validation de l'adresse IP résolue. Le serveur peut
// ainsi être détourné pour émettre des requêtes vers des ressources internes normalement
// inaccessibles depuis l'extérieur (services internes, interfaces d'administration).

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct CallbackRequest {
    callback_url: String,
}

async fn fetch_callback(payload: web::Json<CallbackRequest>) -> impl Responder {
    // Aucune whitelist, aucune validation d'hôte/IP, redirections suivies par défaut.
    match reqwest::get(&payload.callback_url).await {
        Ok(resp) => {
            let body = resp.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(_) => HttpResponse::BadGateway().finish(),
    }
}
