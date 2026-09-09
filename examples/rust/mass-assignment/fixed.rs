// Corrigé : Mass Assignment — CWE-915
// (Improperly Controlled Modification of Dynamically-Determined Object Attributes)
// Un DTO d'entrée distinct du modèle interne est utilisé : il ne contient que les
// champs légitimement modifiables par l'utilisateur (`name`, `email`). Les champs
// sensibles (`role`, `is_admin`) sont totalement absents du binding et ne peuvent
// donc jamais être définis par la requête entrante.

use axum::{extract::{Path, State}, response::IntoResponse, Json};
use serde::Deserialize;
use sqlx::PgPool;

// DTO d'entrée : liste blanche explicite des champs autorisés, séparée du modèle interne.
#[derive(Deserialize)]
struct UpdateUserDto {
    name: String,
    email: String,
    // 'role' et 'is_admin' n'existent pas dans ce DTO : impossibles à injecter.
}

struct AuthUser {
    #[allow(dead_code)]
    id: i64,
}

async fn update_user(
    _auth: AuthUser,
    Path(user_id): Path<i64>,
    State(pool): State<PgPool>,
    Json(payload): Json<UpdateUserDto>,
) -> impl IntoResponse {
    // Seuls les champs autorisés sont écrits ; 'role'/'is_admin' restent inchangés.
    let result = sqlx::query!(
        "UPDATE users SET name = $1, email = $2 WHERE id = $3",
        payload.name, payload.email, user_id
    )
    .execute(&pool)
    .await;

    match result {
        Ok(_) => axum::http::StatusCode::OK.into_response(),
        Err(_) => axum::http::StatusCode::INTERNAL_SERVER_ERROR.into_response(),
    }
}
