// Correction CWE-89 — requête préparée avec paramètre lié et typage strict
// de l'identifiant, et message générique renvoyé au client tandis que le
// détail de l'erreur est journalisé côté serveur uniquement.

use actix_web::{web, HttpResponse, Responder};
use sqlx::PgPool;

pub async fn get_order(pool: web::Data<PgPool>, path: web::Path<i32>) -> impl Responder {
    let id = path.into_inner();

    // Sécurisé : requête préparée, id déjà typé i32 par l'extracteur Path
    match sqlx::query_as::<_, (i32, f64)>("SELECT * FROM orders WHERE id = $1")
        .bind(id)
        .fetch_all(pool.get_ref())
        .await
    {
        Ok(rows) => HttpResponse::Ok().json(rows),
        Err(e) => {
            // Sécurisé : journalisation serveur, message générique au client
            log::error!("erreur requête orders: {}", e);
            HttpResponse::InternalServerError().body("Une erreur est survenue.")
        }
    }
}
