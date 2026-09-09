// CWE-235 — Pollution de paramètres HTTP (HTTP Parameter Pollution)
// Le rôle est extrait via un désérialiseur qui ne retient silencieusement
// qu'une seule valeur, sans détecter ni rejeter une éventuelle duplication
// du paramètre "role" dans la chaîne de requête.

use actix_web::{web, HttpResponse, Responder};

pub async fn assign_role(query: web::Query<RoleQuery>) -> impl Responder {
    // Vulnérable : aucune vérification du nombre d'occurrences de "role"
    // dans la query string brute avant d'utiliser la valeur désérialisée.
    let role = &query.role;

    HttpResponse::Ok().json(serde_json::json!({ "assigned_role": role }))
}

#[derive(serde::Deserialize)]
pub struct RoleQuery {
    pub role: String,
}
