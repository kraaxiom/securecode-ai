// Corrigé : Broken Function Level Authorization (BFLA) — CWE-862 (Missing Authorization)
// Ajout d'une vérification de rôle explicite et centralisée avant toute action
// privilégiée. Le contrôle est effectué côté serveur, indépendamment de ce que
// l'interface utilisateur affiche ou masque.

use actix_web::{web, HttpResponse, Responder};

struct AuthUser {
    id: i64,
    role: String,
}

impl AuthUser {
    /// Vérification de rôle centralisée, réutilisable sur tout endpoint sensible.
    fn require_role(&self, required: &str) -> Result<(), HttpResponse> {
        if self.role != required {
            return Err(HttpResponse::Forbidden().json(serde_json::json!({
                "error": "Accès refusé : privilège insuffisant"
            })));
        }
        Ok(())
    }
}

async fn delete_user(
    auth: AuthUser,
    path: web::Path<i64>,
    pool: web::Data<sqlx::PgPool>,
) -> impl Responder {
    // Vérification de rôle explicite avant toute action d'administration.
    if let Err(forbidden) = auth.require_role("admin") {
        return forbidden;
    }

    let target_id = path.into_inner();

    let result = sqlx::query!("DELETE FROM users WHERE id = $1", target_id)
        .execute(pool.get_ref())
        .await;

    match result {
        Ok(_) => HttpResponse::NoContent().finish(),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}
