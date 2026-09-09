// Vulnérable : Broken Object Level Authorization (BOLA) — CWE-639
// (Authorization Bypass Through User-Controlled Key)
// L'endpoint récupère une commande par son ID transmis par le client sans jamais
// vérifier que cette commande appartient bien à l'utilisateur authentifié.
// Il suffit de faire varier l'ID dans l'URL pour lire la commande de quelqu'un d'autre.

use actix_web::{web, HttpResponse, Responder};

struct AuthUser {
    id: i64,
}

async fn get_order(
    _auth: AuthUser,
    path: web::Path<i64>,
    pool: web::Data<sqlx::PgPool>,
) -> impl Responder {
    let order_id = path.into_inner();

    // Requête filtrée uniquement par ID, aucune clause sur le propriétaire.
    let order = sqlx::query!("SELECT id, total_cents, owner_id FROM orders WHERE id = $1", order_id)
        .fetch_optional(pool.get_ref())
        .await;

    match order {
        Ok(Some(o)) => HttpResponse::Ok().json(serde_json::json!({
            "id": o.id,
            "total_cents": o.total_cents,
        })),
        Ok(None) => HttpResponse::NotFound().finish(),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}
