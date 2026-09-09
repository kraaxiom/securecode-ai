// CWE-89 — Injection SQL basée sur les erreurs (Error-based SQL Injection)
// La requête est construite par concaténation, et le message d'erreur SQL
// détaillé est renvoyé directement au client en cas d'échec.

use actix_web::{web, HttpResponse, Responder};
use sqlx::PgPool;

pub async fn get_order(pool: web::Data<PgPool>, path: web::Path<String>) -> impl Responder {
    let id = path.into_inner();

    // Vulnérable : concaténation directe de l'identifiant dans la requête
    let sql = format!("SELECT * FROM orders WHERE id = {}", id);

    match sqlx::query_as::<_, (i32, f64)>(&sql)
        .fetch_all(pool.get_ref())
        .await
    {
        Ok(rows) => HttpResponse::Ok().json(rows),
        // Vulnérable : le message d'erreur natif du driver est relayé au client
        Err(e) => HttpResponse::InternalServerError().body(format!("Erreur SQL: {}", e)),
    }
}
