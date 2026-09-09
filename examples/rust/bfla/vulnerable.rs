// Vulnérable : Broken Function Level Authorization (BFLA) — CWE-862 (Missing Authorization)
// L'endpoint effectue une action privilégiée (suppression d'un utilisateur) en ne
// vérifiant que l'authentification (`AuthUser`), sans jamais contrôler que l'appelant
// dispose du rôle "admin" requis pour cette fonction. N'importe quel utilisateur
// authentifié peut donc appeler cette route et supprimer n'importe quel compte.

use actix_web::{web, HttpResponse, Responder};

struct AuthUser {
    id: i64,
    #[allow(dead_code)]
    role: String,
}

async fn delete_user(
    auth: AuthUser,
    path: web::Path<i64>,
    pool: web::Data<sqlx::PgPool>,
) -> impl Responder {
    let target_id = path.into_inner();

    // Aucune vérification de rôle : seule la présence d'un utilisateur authentifié compte.
    let result = sqlx::query!("DELETE FROM users WHERE id = $1", target_id)
        .execute(pool.get_ref())
        .await;

    match result {
        Ok(_) => HttpResponse::NoContent().finish(),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}
