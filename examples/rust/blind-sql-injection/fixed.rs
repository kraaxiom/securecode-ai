// Correction CWE-89 — utilisation d'une requête préparée avec paramètre lié,
// même pour une condition dont seul le résultat booléen est renvoyé.
// Le temps de réponse et le message restent uniformes pour limiter
// les canaux d'inférence (pas de branche conditionnelle sur l'erreur SQL).

use actix_web::{web, HttpResponse, Responder};
use sqlx::PgPool;

pub async fn check_user(pool: web::Data<PgPool>, query: web::Query<UserQuery>) -> impl Responder {
    let user = &query.user;

    // Sécurisé : requête préparée, paramètre lié, pas de concaténation
    let row = sqlx::query("SELECT 1 FROM users WHERE username = $1 AND active = true")
        .bind(user)
        .fetch_optional(pool.get_ref())
        .await
        .unwrap_or(None);

    // Réponse et code HTTP uniformes quel que soit le résultat
    HttpResponse::Ok().json(serde_json::json!({ "exists": row.is_some() }))
}

#[derive(serde::Deserialize)]
pub struct UserQuery {
    pub user: String,
}
