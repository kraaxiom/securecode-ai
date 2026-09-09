// Corrigé : Privilege Escalation — CWE-269 (Improper Privilege Management)
// L'attribution de rôle vérifie explicitement que l'appelant est lui-même
// autorisé à accorder le rôle demandé (garde en handler), puis invalide les
// sessions actives de la cible pour éviter la persistance des anciens privilèges.

use axum::{extract::{Path, State}, response::IntoResponse, Json};
use serde::Deserialize;
use sqlx::PgPool;

#[derive(Deserialize)]
struct RoleUpdate {
    role: String,
}

struct AuthUser {
    #[allow(dead_code)]
    id: i64,
    role: String,
}

impl AuthUser {
    /// Un acteur ne peut accorder qu'un rôle strictement inférieur au sien,
    /// et seul un admin peut accorder le rôle "admin".
    fn can_grant_role(&self, requested: &str) -> bool {
        match (self.role.as_str(), requested) {
            ("admin", _) => true,
            ("manager", "member") => true,
            _ => false,
        }
    }
}

async fn assign_role(
    auth: AuthUser,
    Path(user_id): Path<i64>,
    State(pool): State<PgPool>,
    Json(payload): Json<RoleUpdate>,
) -> impl IntoResponse {
    // Vérification explicite du droit d'attribution avant toute écriture.
    if !auth.can_grant_role(&payload.role) {
        return (
            axum::http::StatusCode::FORBIDDEN,
            Json(serde_json::json!({ "error": "Attribution de rôle non autorisée" })),
        )
            .into_response();
    }

    let result = sqlx::query!("UPDATE users SET role = $1 WHERE id = $2", payload.role, user_id)
        .execute(&pool)
        .await;

    if result.is_err() {
        return axum::http::StatusCode::INTERNAL_SERVER_ERROR.into_response();
    }

    // Invalide les sessions actives de la cible pour forcer une réauthentification.
    let _ = sqlx::query!("DELETE FROM sessions WHERE user_id = $1", user_id)
        .execute(&pool)
        .await;

    axum::http::StatusCode::OK.into_response()
}
