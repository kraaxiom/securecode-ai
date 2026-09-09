// Vulnérable : Forced Browsing — CWE-425 (Direct Request ('Forced Browsing'))
// La page de rapports internes n'est liée nulle part dans l'interface utilisateur,
// mais reste accessible sans aucun contrôle d'autorisation propre : la "sécurité"
// repose uniquement sur le fait que l'URL n'est pas découvrable, pas sur un
// contrôle d'accès réel côté serveur.

use axum::{extract::State, response::IntoResponse, Json};
use sqlx::PgPool;

// Route enregistrée ailleurs : .route("/internal-reports", get(internal_reports))
// Elle n'exige même pas d'authentification, encore moins un rôle spécifique.
async fn internal_reports(State(pool): State<PgPool>) -> impl IntoResponse {
    let reports = sqlx::query!("SELECT id, title, revenue_cents FROM internal_reports")
        .fetch_all(&pool)
        .await
        .unwrap_or_default();

    let payload: Vec<_> = reports
        .into_iter()
        .map(|r| serde_json::json!({ "id": r.id, "title": r.title, "revenue_cents": r.revenue_cents }))
        .collect();

    Json(payload)
}
