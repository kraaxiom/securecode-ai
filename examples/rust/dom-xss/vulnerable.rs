// Vulnérable : CWE-79 — DOM XSS côté serveur (angle Rust/axum)
// Le paramètre de requête `q` est injecté directement dans un bloc <script>
// inline de la réponse HTML. Le navigateur exécute ce script avec la donnée
// utilisateur non échappée, ouvrant la voie à une rupture du contexte JS.

use axum::{extract::Query, response::Html, routing::get, Router};
use std::collections::HashMap;

async fn search(Query(params): Query<HashMap<String, String>>) -> Html<String> {
    let q = params.get("q").cloned().unwrap_or_default();

    // Concaténation directe de la donnée utilisateur dans un script inline.
    let body = format!(
        "<html><body><script>var q = '{}';</script></body></html>",
        q
    );
    Html(body)
}

pub fn app() -> Router {
    Router::new().route("/search", get(search))
}
