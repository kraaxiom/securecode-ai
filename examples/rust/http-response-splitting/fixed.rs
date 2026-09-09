// Correction CWE-113 — la redirection est restreinte à une liste blanche
// de chemins connus ; toute valeur hors liste (y compris contenant des
// caractères CR/LF) retombe sur une destination par défaut sûre.

use actix_web::{web, HttpResponse, Responder};
use std::collections::HashSet;

pub async fn redirect(query: web::Query<RedirectQuery>) -> impl Responder {
    let allowed: HashSet<&str> = ["/dashboard", "/profile", "/account"].into_iter().collect();

    // Sécurisé : liste blanche stricte, aucune valeur brute n'atteint l'en-tête
    let next = query
        .next
        .as_deref()
        .filter(|n| allowed.contains(n))
        .unwrap_or("/dashboard");

    HttpResponse::Found()
        .append_header(("Location", next))
        .finish()
}

#[derive(serde::Deserialize)]
pub struct RedirectQuery {
    pub next: Option<String>,
}
