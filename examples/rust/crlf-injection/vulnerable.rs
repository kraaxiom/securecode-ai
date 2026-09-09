// CWE-93 — Injection CRLF (CRLF Injection)
// La valeur "next" fournie par l'utilisateur est écrite telle quelle dans
// l'en-tête Location, sans filtrage des caractères \r et \n, permettant
// l'ajout de lignes/en-têtes non prévus.

use actix_web::{web, HttpResponse, Responder};

pub async fn redirect(query: web::Query<RedirectQuery>) -> impl Responder {
    let next = &query.next;

    // Vulnérable : valeur utilisateur injectée brute dans l'en-tête
    HttpResponse::Found()
        .append_header(("Location", next.as_str()))
        .finish()
}

#[derive(serde::Deserialize)]
pub struct RedirectQuery {
    pub next: String,
}
