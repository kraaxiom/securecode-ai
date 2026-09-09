// CWE-444 — Contrebande de requêtes HTTP (HTTP Request Smuggling)
// L'application fait confiance aux en-têtes Content-Length et
// Transfer-Encoding transmis par un proxy en amont sans vérifier leur
// cohérence, ouvrant la porte à une divergence d'interprétation.

use actix_web::{web, HttpRequest, HttpResponse, Responder};

pub async fn handle_api(req: HttpRequest, body: web::Bytes) -> impl Responder {
    // Vulnérable : aucune vérification de la présence conjointe des deux en-têtes
    let has_length = req.headers().contains_key("content-length");
    let has_encoding = req.headers().contains_key("transfer-encoding");

    // Traitement de la requête sans tenir compte de l'ambiguïté potentielle
    let _ = (has_length, has_encoding);
    HttpResponse::Ok().body(body)
}
