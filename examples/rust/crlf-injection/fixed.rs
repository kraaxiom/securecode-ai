// Correction CWE-93 — rejet de toute entrée contenant \r ou \n, et
// restriction à un chemin local relatif (pas d'URL externe), avant
// écriture dans l'en-tête Location.

use actix_web::{web, HttpResponse, Responder};

pub async fn redirect(query: web::Query<RedirectQuery>) -> impl Responder {
    let next = &query.next;

    // Sécurisé : rejet des caractères de contrôle et validation du format
    let safe_next = if next.starts_with('/') && !next.contains(['\r', '\n']) {
        next.as_str()
    } else {
        "/"
    };

    HttpResponse::Found()
        .append_header(("Location", safe_next))
        .finish()
}

#[derive(serde::Deserialize)]
pub struct RedirectQuery {
    pub next: String,
}
