// CWE-89 : Improper Neutralization of Special Elements used in an SQL Command
// (variante UNION-based). L'identifiant `id` est concaténé sans validation ni
// paramètre lié, ouvrant la porte à l'ajout d'une clause UNION SELECT pour
// exfiltrer des données d'autres tables.

use actix_web::{web, HttpResponse};
use sqlx::PgPool;

#[derive(serde::Deserialize)]
struct ProductQuery {
    id: String,
}

async fn product_handler(
    pool: web::Data<PgPool>,
    query: web::Query<ProductQuery>,
) -> HttpResponse {
    // Vulnérable : la valeur brute non typée est insérée directement dans le SQL
    let sql = format!("SELECT id, title, price FROM products WHERE id = {}", query.id);

    let rows = sqlx::query(&sql).fetch_all(pool.get_ref()).await;

    match rows {
        Ok(r) => HttpResponse::Ok().body(format!("{} produit(s)", r.len())),
        Err(_) => HttpResponse::InternalServerError().body("Erreur base de données"),
    }
}
