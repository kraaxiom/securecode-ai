// CWE-90 — Injection LDAP (LDAP Injection)
// L'identifiant utilisateur est concaténé directement dans le filtre de
// recherche LDAP, sans échappement des caractères spéciaux LDAP.

use actix_web::{web, HttpResponse, Responder};

pub async fn find_user(query: web::Query<LdapQuery>) -> impl Responder {
    let uid = &query.uid;

    // Vulnérable : filtre LDAP construit par concaténation directe
    let filter = format!("(uid={})", uid);

    HttpResponse::Ok().body(format!("filtre utilisé: {}", filter))
}

#[derive(serde::Deserialize)]
pub struct LdapQuery {
    pub uid: String,
}
