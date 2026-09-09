// Corrigé : Broken Access Control (catégorie générale) — CWE-284
// (Improper Access Control)
// Une couche d'autorisation centralisée et explicite est appliquée avant de servir
// la ressource, avec un principe "deny by default" : l'accès n'est autorisé que si
// l'utilisateur est propriétaire de la facture ou administrateur.

use actix_web::{web, HttpResponse, Responder};
use std::path::PathBuf;

struct AuthUser {
    id: i64,
    role: String,
}

/// Politique d'autorisation centralisée, appliquée systématiquement.
fn can_view_invoice(user: &AuthUser, owner_id: i64) -> bool {
    user.id == owner_id || user.role == "admin"
}

async fn download_invoice(
    auth: AuthUser,
    path: web::Path<i64>,
    pool: web::Data<sqlx::PgPool>,
) -> impl Responder {
    let invoice_id = path.into_inner();

    let invoice = sqlx::query!("SELECT file_path, owner_id FROM invoices WHERE id = $1", invoice_id)
        .fetch_optional(pool.get_ref())
        .await;

    match invoice {
        Ok(Some(inv)) => {
            // Deny by default : refus explicite si la policy ne l'autorise pas.
            if !can_view_invoice(&auth, inv.owner_id) {
                return HttpResponse::Forbidden().finish();
            }
            let file = PathBuf::from(inv.file_path);
            match actix_files::NamedFile::open(file) {
                Ok(f) => f.into_response(&actix_web::HttpRequest::from_raw(std::ptr::null())),
                Err(_) => HttpResponse::NotFound().finish(),
            }
        }
        Ok(None) => HttpResponse::NotFound().finish(),
        Err(_) => HttpResponse::InternalServerError().finish(),
    }
}
