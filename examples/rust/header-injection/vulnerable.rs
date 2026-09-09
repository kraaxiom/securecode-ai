// CWE-113 — Injection d'en-têtes HTTP (Header Injection)
// Le nom de fichier fourni par l'utilisateur est concaténé directement
// dans l'en-tête Content-Disposition, sans filtrage des caractères \r/\n.

use actix_web::{web, HttpResponse, Responder};

pub async fn download(query: web::Query<DownloadQuery>) -> impl Responder {
    let filename = &query.filename;

    // Vulnérable : concaténation brute dans la valeur d'en-tête
    let header_value = format!("attachment; filename={}", filename);

    HttpResponse::Ok()
        .append_header(("Content-Disposition", header_value))
        .body("contenu du fichier")
}

#[derive(serde::Deserialize)]
pub struct DownloadQuery {
    pub filename: String,
}
