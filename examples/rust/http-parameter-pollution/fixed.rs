// Correction CWE-235 — vérification explicite du nombre d'occurrences du
// paramètre "role" dans la query string brute ; toute duplication est
// rejetée plutôt que résolue silencieusement.

use actix_web::{web, HttpResponse, HttpRequest, Responder};

pub async fn assign_role(req: HttpRequest, query: web::Query<RoleQuery>) -> impl Responder {
    // Sécurisé : rejet explicite si le paramètre "role" apparaît plusieurs fois
    let occurrences = req.query_string().matches("role=").count();
    if occurrences > 1 {
        return HttpResponse::BadRequest().body("Paramètre dupliqué non autorisé");
    }

    let role = &query.role;
    HttpResponse::Ok().json(serde_json::json!({ "assigned_role": role }))
}

#[derive(serde::Deserialize)]
pub struct RoleQuery {
    pub role: String,
}
