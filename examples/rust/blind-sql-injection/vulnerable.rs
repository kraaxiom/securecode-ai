// CWE-89 — Injection SQL aveugle (Blind SQL Injection)
// La valeur utilisateur est concaténée dans la requête SQL sans requête
// préparée, même si seul un booléen "exists" est retourné au client.
// L'absence de paramètre lié permet à un attaquant d'inférer des données
// via des différences de comportement (contenu, code HTTP, timing).

use actix_web::{web, HttpResponse, Responder};
use sqlx::PgPool;

pub async fn check_user(pool: web::Data<PgPool>, query: web::Query<UserQuery>) -> impl Responder {
    let user = &query.user;

    // Vulnérable : concaténation directe dans la requête SQL
    let sql = format!(
        "SELECT 1 FROM users WHERE username = '{}' AND active = true",
        user
    );

    let row = sqlx::query(&sql)
        .fetch_optional(pool.get_ref())
        .await
        .unwrap_or(None);

    HttpResponse::Ok().json(serde_json::json!({ "exists": row.is_some() }))
}

#[derive(serde::Deserialize)]
pub struct UserQuery {
    pub user: String,
}
