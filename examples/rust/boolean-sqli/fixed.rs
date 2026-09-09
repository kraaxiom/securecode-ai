// Correction CWE-89 — requête préparée avec paramètre lié pour toute
// condition impliquant une donnée utilisateur ; plus aucune comparaison
// SQL n'est construite par concaténation de chaîne.

use actix_web::{web, HttpResponse, Responder};
use sqlx::PgPool;

pub async fn search_products(
    pool: web::Data<PgPool>,
    query: web::Query<SearchQuery>,
) -> impl Responder {
    let name = &query.name;

    // Sécurisé : paramètre lié, aucune interprétation de la valeur comme code SQL
    let rows = sqlx::query_as::<_, (i32, String)>("SELECT * FROM products WHERE name = $1")
        .bind(name)
        .fetch_all(pool.get_ref())
        .await
        .unwrap_or_default();

    HttpResponse::Ok().json(rows)
}

#[derive(serde::Deserialize)]
pub struct SearchQuery {
    pub name: String,
}
