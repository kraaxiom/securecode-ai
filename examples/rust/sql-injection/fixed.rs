// CWE-89 : correction — requête préparée avec paramètre lié via sqlx (`bind`),
// aucune concaténation de valeur utilisateur dans le texte SQL.

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
    // Sécurisé : paramètre lié, la valeur n'est jamais interpolée dans le texte SQL
    let rows = sqlx::query("SELECT id, name FROM users WHERE name = $1")
        .bind(&query.name)
        .fetch_all(pool.get_ref())
        .await;

    match rows {
        Ok(r) => HttpResponse::Ok().body(format!("{} résultat(s)", r.len())),
        Err(_) => HttpResponse::InternalServerError().body("Erreur base de données"),
    }
}
