// Correction CWE-93 — les guillemets et caractères de contrôle sont
// neutralisés avant assemblage, et la longueur du critère est bornée ;
// dans une intégration réelle, on privilégierait l'API structurée du
// client IMAP plutôt qu'une commande texte assemblée à la main.

use actix_web::{web, HttpResponse, Responder};

fn sanitize_imap_criteria(input: &str) -> String {
    input
        .replace(['"', '\r', '\n'], "")
        .chars()
        .take(200)
        .collect()
}

pub async fn search_mail(query: web::Query<SearchQuery>) -> impl Responder {
    let criteria = sanitize_imap_criteria(&query.q);

    // Sécurisé : critère neutralisé et borné avant assemblage de la commande
    let command = format!("SEARCH SUBJECT \"{}\"", criteria);

    HttpResponse::Ok().body(format!("commande envoyée: {}", command))
}

#[derive(serde::Deserialize)]
pub struct SearchQuery {
    pub q: String,
}
