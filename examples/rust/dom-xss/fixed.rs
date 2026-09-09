// Corrigé : CWE-79 — la donnée utilisateur n'est jamais concaténée directement
// dans du JavaScript inline. Elle est sérialisée en JSON (échappement des
// caractères dangereux pour un contexte script) avant insertion, conformément
// à rules/remediation/dom-xss.md.

use axum::{extract::Query, response::Html, routing::get, Router};
use serde_json::json;
use std::collections::HashMap;

async fn search(Query(params): Query<HashMap<String, String>>) -> Html<String> {
    let q = params.get("q").cloned().unwrap_or_default();

    // Sérialisation JSON sûre : les guillemets, apostrophes et balises
    // dans `q` sont correctement échappés pour un contexte <script>.
    let safe_json = json!(q).to_string();
    let body = format!(
        "<html><body><script>var q = {};</script></body></html>",
        safe_json
    );
    Html(body)
}

pub fn app() -> Router {
    Router::new().route("/search", get(search))
}
