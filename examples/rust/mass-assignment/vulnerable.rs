// Vulnérable : Mass Assignment — CWE-915
// (Improperly Controlled Modification of Dynamically-Determined Object Attributes)
// Le corps JSON entier est désérialisé directement dans le modèle de base de
// données `User`, qui contient des champs sensibles (`role`, `is_admin`). Un
// attaquant peut donc injecter ces champs dans la requête et s'auto-promouvoir.

use axum::{extract::{Path, State}, response::IntoResponse, Json};
use serde::Deserialize;
use sqlx::PgPool;

#[derive(Deserialize)]
struct User {
    #[allow(dead_code)]
    id: i64,
    name: String,
    email: String,
    role: String,      // champ sensible exposé au binding direct
    is_admin: bool,     // champ sensible exposé au binding direct
}

struct AuthUser {
    #[allow(dead_code)]
    id: i64,
}

// Le body entier (potentiellement { "name": ..., "role": "admin", "is_admin": true })
// est désérialisé tel quel dans le modèle interne complet.
async fn update_user(
    _auth: AuthUser,
    Path(user_id): Path<i64>,
    State(pool): State<PgPool>,
    Json(payload): Json<User>,
) -> impl IntoResponse {
    let result = sqlx::query!(
        "UPDATE users SET name = $1, email = $2, role = $3, is_admin = $4 WHERE id = $5",
        payload.name, payload.email, payload.role, payload.is_admin, user_id
    )
    .execute(&pool)
    .await;

    match result {
        Ok(_) => axum::http::StatusCode::OK.into_response(),
        Err(_) => axum::http::StatusCode::INTERNAL_SERVER_ERROR.into_response(),
    }
}
