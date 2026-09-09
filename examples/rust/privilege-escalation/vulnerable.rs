// Vulnérable : Privilege Escalation — CWE-269 (Improper Privilege Management)
// L'endpoint d'attribution de rôle accepte n'importe quelle valeur de `role`
// envoyée par le client et l'applique sans vérifier que l'appelant a lui-même le
// droit d'accorder ce niveau de privilège. Un utilisateur standard peut ainsi
// s'attribuer (ou attribuer à un tiers) le rôle "admin".

use axum::{extract::{Path, State}, response::IntoResponse, Json};
use serde::Deserialize;
use sqlx::PgPool;

#[derive(Deserialize)]
struct RoleUpdate {
    role: String, // valeur entièrement contrôlée par le client, non validée par rapport à l'appelant
}

struct AuthUser {
    #[allow(dead_code)]
    id: i64,
    #[allow(dead_code)]
    role: String,
}

async fn assign_role(
    _auth: AuthUser,
    Path(user_id): Path<i64>,
    State(pool): State<PgPool>,
    Json(payload): Json<RoleUpdate>,
) -> impl IntoResponse {
    // Aucune vérification que l'appelant peut accorder ce rôle précis.
    let result = sqlx::query!("UPDATE users SET role = $1 WHERE id = $2", payload.role, user_id)
        .execute(&pool)
        .await;

    match result {
        Ok(_) => axum::http::StatusCode::OK.into_response(),
        Err(_) => axum::http::StatusCode::INTERNAL_SERVER_ERROR.into_response(),
    }
}
