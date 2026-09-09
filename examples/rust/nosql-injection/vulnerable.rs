// CWE-943 — Injection NoSQL (NoSQL Injection)
// Le corps JSON de la requête est transmis tel quel comme filtre de
// requête MongoDB, sans valider que les champs sont bien des chaînes
// scalaires (un objet comme {"$ne": null} serait accepté).

use actix_web::{web, HttpResponse, Responder};
use mongodb::{bson::Document, Collection};

pub async fn login(
    collection: web::Data<Collection<Document>>,
    body: web::Json<serde_json::Value>,
) -> impl Responder {
    // Vulnérable : le JSON brut est converti et utilisé directement comme filtre
    let filter: Document = mongodb::bson::to_document(&body.into_inner()).unwrap_or_default();

    match collection.find_one(filter, None).await {
        Ok(Some(_user)) => HttpResponse::Ok().body("connecté"),
        _ => HttpResponse::Unauthorized().body("échec"),
    }
}
