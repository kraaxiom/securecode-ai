// Correction CWE-90 — échappement dédié des caractères spéciaux LDAP
// (*, (, ), \, NUL) avant insertion dans le filtre de recherche.

use actix_web::{web, HttpResponse, Responder};

fn escape_ldap_filter(input: &str) -> String {
    let mut out = String::with_capacity(input.len());
    for c in input.chars() {
        match c {
            '*' => out.push_str("\\2a"),
            '(' => out.push_str("\\28"),
            ')' => out.push_str("\\29"),
            '\\' => out.push_str("\\5c"),
            '\0' => out.push_str("\\00"),
            other => out.push(other),
        }
    }
    out
}

pub async fn find_user(query: web::Query<LdapQuery>) -> impl Responder {
    // Sécurisé : échappement dédié avant construction du filtre
    let uid = escape_ldap_filter(&query.uid);
    let filter = format!("(uid={})", uid);

    HttpResponse::Ok().body(format!("filtre utilisé: {}", filter))
}

#[derive(serde::Deserialize)]
pub struct LdapQuery {
    pub uid: String,
}
