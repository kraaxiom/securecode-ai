// Corrigé : Insecure Direct Object Reference (IDOR) — CWE-639
// (Authorization Bypass Through User-Controlled Key)
// La clause de vérification d'appartenance (`owner_id = $2`) est intégrée
// directement dans la requête de base de données : impossible de récupérer un
// document appartenant à un autre utilisateur, même en devinant son ID.

use axum::{extract::{Path, State}, response::IntoResponse, Json};
use sqlx::PgPool;

struct AuthUser {
    id: i64,
}

async fn get_document(
    auth: AuthUser,
    Path(doc_id): Path<i64>,
    State(pool): State<PgPool>,
) -> impl IntoResponse {
    // Clause d'appartenance directement dans la requête, pas en post-traitement.
    let doc = sqlx::query!(
        "SELECT id, title, content FROM documents WHERE id = $1 AND owner_id = $2",
        doc_id,
        auth.id
    )
    .fetch_optional(&pool)
    .await;

    match doc {
        Ok(Some(d)) => Json(serde_json::json!({
            "id": d.id, "title": d.title, "content": d.content
        }))
        .into_response(),
        // 404 générique, qu'il n'existe pas ou qu'il appartienne à quelqu'un d'autre.
        Ok(None) => axum::http::StatusCode::NOT_FOUND.into_response(),
        Err(_) => axum::http::StatusCode::INTERNAL_SERVER_ERROR.into_response(),
    }
}
