// Corrigé : CWE-79 — le paramètre de requête est encodé contextuellement
// (échappement HTML) avant insertion dans la réponse, conformément à
// rules/remediation/reflected-xss.md.

use actix_web::{get, web, HttpResponse, Responder};
use std::collections::HashMap;

fn escape_html(input: &str) -> String {
    input
        .replace('&', "&amp;")
        .replace('<', "&lt;")
        .replace('>', "&gt;")
        .replace('"', "&quot;")
        .replace('\'', "&#39;")
}

#[get("/search")]
async fn search(query: web::Query<HashMap<String, String>>) -> impl Responder {
    let term = query.get("q").cloned().unwrap_or_default();

    // Échappement systématique de la donnée issue de la requête.
    let body = format!(
        "<html><body><h1>Résultats pour : {}</h1></body></html>",
        escape_html(&term)
    );

    HttpResponse::Ok().content_type("text/html; charset=utf-8").body(body)
}
