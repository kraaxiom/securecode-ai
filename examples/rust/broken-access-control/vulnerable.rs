// Vulnérable : Broken Access Control (catégorie générale) — CWE-284
// (Improper Access Control)
// Aucun contrôle d'autorisation n'est effectué du tout sur la ressource : la route
// vérifie seulement que l'utilisateur est authentifié, puis sert directement le
// fichier de facture, quel que soit son propriétaire réel.

use actix_web::{web, HttpResponse, Responder};
use std::path::PathBuf;

struct AuthUser {
    #[allow(dead_code)]
    id: i64,
}

async fn download_invoice(
    _auth: AuthUser,
    path: web::Path<i64>,
    pool: web::Data<sqlx::PgPool>,
) -> impl Responder {
    let invoice_id = path.into_inner();

    let invoice = sqlx::query!("SELECT file_path, owner_id FROM invoices WHERE id = $1", invoice_id)
        .fetch_optional(pool.get_ref())
        .await;

    match invoice {
        // Aucune vérification que owner_id correspond à l'utilisateur courant.
        Ok(Some(inv)) => {
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
