// CWE-89 — Injection SQL booléenne (Boolean-based SQL Injection)
// Le nom de produit est concaténé directement dans la clause WHERE,
// permettant à un attaquant de modifier la valeur de vérité de la
// condition et d'observer une différence binaire dans la réponse.

use actix_web::{web, HttpResponse, Responder};
use sqlx::PgPool;

pub async fn search_products(
    pool: web::Data<PgPool>,
    query: web::Query<SearchQuery>,
) -> impl Responder {
    let name = &query.name;

    // Vulnérable : comparaison SQL construite par concaténation
    let sql = format!("SELECT * FROM products WHERE name = '{}'", name);

    let rows = sqlx::query_as::<_, (i32, String)>(&sql)
        .fetch_all(pool.get_ref())
        .await
        .unwrap_or_default();

    HttpResponse::Ok().json(rows)
}

#[derive(serde::Deserialize)]
pub struct SearchQuery {
    pub name: String,
}
