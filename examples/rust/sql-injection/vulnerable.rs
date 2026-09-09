// CWE-89 : Improper Neutralization of Special Elements used in an SQL Command
// Le paramètre `name` est concaténé directement dans le texte de la requête
// SQL, permettant à un attaquant de modifier la structure de la requête.

use actix_web::{web, HttpResponse};
use sqlx::PgPool;

#[derive(serde::Deserialize)]
struct SearchQuery {
    name: String,
}

async fn search_users_handler(
    pool: web::Data<PgPool>,
    query: web::Query<SearchQuery>,
) -> HttpResponse {
    // Vulnérable : concaténation directe de l'entrée utilisateur dans le SQL
    let sql = format!("SELECT id, name FROM users WHERE name = '{}'", query.name);

    let rows = sqlx::query(&sql).fetch_all(pool.get_ref()).await;

    match rows {
        Ok(r) => HttpResponse::Ok().body(format!("{} résultat(s)", r.len())),
        Err(_) => HttpResponse::InternalServerError().body("Erreur base de données"),
    }
}
