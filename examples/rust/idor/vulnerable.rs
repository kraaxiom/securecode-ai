// Vulnérable : Insecure Direct Object Reference (IDOR) — CWE-639
// (Authorization Bypass Through User-Controlled Key)
// Le document est récupéré uniquement par l'ID fourni par le client. La seule
// vérification effectuée est que l'utilisateur est connecté ; rien ne garantit
// que le document appartient bien à cet utilisateur.

use axum::{extract::{Path, State}, response::IntoResponse, Json};
use sqlx::PgPool;

struct AuthUser {
    #[allow(dead_code)]
    id: i64,
}

async fn get_document(
    _auth: AuthUser,
    Path(doc_id): Path<i64>,
    State(pool): State<PgPool>,
) -> impl IntoResponse {
    // Requête basée uniquement sur l'ID transmis, sans clause d'appartenance.
    let doc = sqlx::query!("SELECT id, title, content FROM documents WHERE id = $1", doc_id)
        .fetch_optional(&pool)
        .await;

    match doc {
        Ok(Some(d)) => Json(serde_json::json!({
            "id": d.id, "title": d.title, "content": d.content
        }))
        .into_response(),
        Ok(None) => axum::http::StatusCode::NOT_FOUND.into_response(),
        Err(_) => axum::http::StatusCode::INTERNAL_SERVER_ERROR.into_response(),
    }
}
