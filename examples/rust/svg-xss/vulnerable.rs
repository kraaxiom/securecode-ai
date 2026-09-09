// Vulnérable : CWE-79 — XSS via SVG
// Un fichier SVG uploadé comme avatar est stocké puis servi tel quel avec
// Content-Type image/svg+xml depuis la même origine que l'application,
// sans sanitisation du contenu XML (balises script, gestionnaires onload/onerror).

use actix_web::{get, web, HttpResponse, Responder};
use std::fs;

#[get("/uploads/avatars/{filename}")]
async fn serve_avatar(path: web::Path<String>) -> impl Responder {
    let filename = path.into_inner();
    let file_path = format!("./uploads/avatars/{}", filename);

    // Le fichier SVG est lu et renvoyé sans aucune sanitisation, affiché
    // inline dans le navigateur, dans le contexte d'origine de l'application.
    match fs::read(&file_path) {
        Ok(content) => HttpResponse::Ok()
            .content_type("image/svg+xml")
            .body(content),
        Err(_) => HttpResponse::NotFound().finish(),
    }
}
