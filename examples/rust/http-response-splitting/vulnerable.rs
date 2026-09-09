// CWE-113 — Fractionnement de réponse HTTP (HTTP Response Splitting)
// La valeur "next" fournie par l'utilisateur est utilisée telle quelle
// pour construire une redirection, sans liste blanche ni filtrage des
// caractères CR/LF, permettant l'injection d'une seconde réponse.

use actix_web::{web, HttpResponse, Responder};

pub async fn redirect(query: web::Query<RedirectQuery>) -> impl Responder {
    let next = query.next.clone().unwrap_or_else(|| "/dashboard".to_string());

    // Vulnérable : redirection construite directement depuis l'entrée utilisateur
    HttpResponse::Found()
        .append_header(("Location", next))
        .finish()
}

#[derive(serde::Deserialize)]
pub struct RedirectQuery {
    pub next: Option<String>,
}
