// Vulnérable : CWE-79 — XSS réfléchi
// Le paramètre de recherche `q` issu de la requête HTTP est réinséré tel quel
// dans le HTML de la page de résultats, sans encodage contextuel.

use actix_web::{get, web, HttpResponse, Responder};
use std::collections::HashMap;

#[get("/search")]
async fn search(query: web::Query<HashMap<String, String>>) -> impl Responder {
    let term = query.get("q").cloned().unwrap_or_default();

    // Concaténation directe du paramètre de requête dans le HTML de réponse.
    let body = format!("<html><body><h1>Résultats pour : {}</h1></body></html>", term);

    HttpResponse::Ok().content_type("text/html; charset=utf-8").body(body)
}
