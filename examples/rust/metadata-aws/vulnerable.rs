// Vulnérable : SSRF vers le service de métadonnées AWS (IMDS) — CWE-918 (Server-Side Request Forgery)
// La fonction de récupération d'URL ne filtre pas la plage d'adresses link-local. Hébergée
// sur une instance EC2/ECS, elle peut donc être détournée pour interroger l'IMDS
// (169.254.169.254) et, si l'infrastructure utilise encore IMDSv1, récupérer les
// identifiants temporaires du rôle IAM attaché sans jeton de session.

use actix_web::{web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct FetchRequest {
    url: String,
}

async fn fetch_resource(payload: web::Json<FetchRequest>) -> impl Responder {
    // Aucun filtrage des plages privées/link-local : 169.254.169.254 est atteignable.
    match reqwest::get(&payload.url).await {
        Ok(resp) => {
            let body = resp.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(_) => HttpResponse::BadGateway().finish(),
    }
}
