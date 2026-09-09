// CWE-93 — Injection IMAP (IMAP Injection)
// Le critère de recherche fourni par l'utilisateur est concaténé
// directement dans une commande IMAP texte, sans échapper les guillemets
// ni les caractères de contrôle.

use actix_web::{web, HttpResponse, Responder};

pub async fn search_mail(query: web::Query<SearchQuery>) -> impl Responder {
    let criteria = &query.q;

    // Vulnérable : commande IMAP assemblée par concaténation de chaîne brute
    let command = format!("SEARCH SUBJECT \"{}\"", criteria);

    HttpResponse::Ok().body(format!("commande envoyée: {}", command))
}

#[derive(serde::Deserialize)]
pub struct SearchQuery {
    pub q: String,
}
