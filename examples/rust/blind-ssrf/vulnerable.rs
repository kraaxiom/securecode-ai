// Vulnérable : Blind SSRF — CWE-918 (Server-Side Request Forgery)
// Un job asynchrone vérifie une URL fournie par l'utilisateur en la contactant côté serveur,
// mais ne retourne jamais le contenu ni le statut de la réponse à l'appelant. Comme pour une
// SSRF classique, aucune whitelist ni validation d'IP n'est appliquée : la requête sortante
// reste exploitable même si son résultat n'est jamais visible directement.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct WebhookRegistration {
    url: String,
}

async fn register_webhook(payload: web::Json<WebhookRegistration>) -> impl Responder {
    let url = payload.url.clone();

    // Traitement "fire and forget" : le résultat de la requête n'est jamais renvoyé.
    tokio::spawn(async move {
        let _ = reqwest::get(&url).await;
    });

    HttpResponse::Accepted().body("Webhook en cours de vérification")
}
