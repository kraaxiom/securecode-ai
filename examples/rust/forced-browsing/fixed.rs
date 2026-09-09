// Corrigé : Forced Browsing — CWE-425 (Direct Request ('Forced Browsing'))
// Un contrôle d'authentification et de rôle explicite est imposé sur la route,
// indépendamment de sa visibilité dans l'interface. La découvrabilité de l'URL
// n'est plus le seul rempart : le serveur refuse l'accès à quiconque n'a pas le
// rôle requis, exactement comme pour une route visible dans l'UI.

use axum::{extract::State, response::IntoResponse, Json};
use sqlx::PgPool;

struct AuthUser {
    role: String,
}

// Extracteur simplifié : AuthUser est résolu à partir de la session/JWT en amont.
async fn internal_reports(
    auth: AuthUser,
    State(pool): State<PgPool>,
) -> impl IntoResponse {
    // Contrôle d'autorisation explicite, appliqué même si la route n'est pas dans l'UI.
    if auth.role != "analyst" && auth.role != "admin" {
        return (
            axum::http::StatusCode::FORBIDDEN,
            Json(serde_json::json!({ "error": "Accès refusé" })),
        )
            .into_response();
    }

    let reports = sqlx::query!("SELECT id, title, revenue_cents FROM internal_reports")
        .fetch_all(&pool)
        .await
        .unwrap_or_default();

    let payload: Vec<_> = reports
        .into_iter()
        .map(|r| serde_json::json!({ "id": r.id, "title": r.title, "revenue_cents": r.revenue_cents }))
        .collect();

    Json(payload).into_response()
}
