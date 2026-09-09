// Corrigé : Broken Object Level Authorization (BOLA) — CWE-639
// (Authorization Bypass Through User-Controlled Key)
// La vérification d'appartenance est intégrée directement dans la requête de
// données : seule une commande appartenant à l'utilisateur authentifié peut être
// retournée, quel que soit l'ID demandé.

use actix_web::{web, HttpResponse, Responder};

struct AuthUser {
    id: i64,
}

async fn get_order(
    auth: AuthUser,
    path: web::Path<i64>,
    pool: web::Data<sqlx::PgPool>,
) -> impl Responder {
    let order_id = path.into_inner();

    // Filtrage d'appartenance directement dans la requête (pas en post-traitement).
    let order = sqlx::query!(
        "SELECT id, total_cents, owner_id FROM orders WHERE id = $1 AND owner_id = $2",
        order_id,
        auth.id
    )
    .fetch_optional(pool.get_ref())
    .await;

    match order {
        Ok(Some(o)) => HttpResponse::Ok().json(serde_json::json!({
            "id": o.id,
            "total_cents": o.total_cents,
        })),
        // 404 générique : ne révèle pas si la ressource existe pour un autre utilisateur.
        Ok(None) => HttpResponse::NotFound().finish(),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}
