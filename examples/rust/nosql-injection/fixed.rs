// Correction CWE-943 — validation stricte du schéma : username/password
// doivent être des chaînes, jamais des objets/opérateurs MongoDB. Le
// filtre est reconstruit explicitement à partir de valeurs typées.

use actix_web::{web, HttpResponse, Responder};
use mongodb::{bson::doc, Collection};

#[derive(serde::Deserialize)]
pub struct LoginRequest {
    pub username: String,
    pub password: String,
}

pub async fn login(
    collection: web::Data<Collection<mongodb::bson::Document>>,
    body: web::Json<LoginRequest>,
) -> impl Responder {
    // Sécurisé : le type serde force username/password à être des chaînes ;
    // aucun objet/opérateur ($ne, $where...) ne peut être injecté ici.
    let filter = doc! {
        "username": &body.username,
        "password": &body.password,
    };

    match collection.find_one(filter, None).await {
        Ok(Some(_user)) => HttpResponse::Ok().body("connecté"),
        _ => HttpResponse::Unauthorized().body("échec"),
    }
}
