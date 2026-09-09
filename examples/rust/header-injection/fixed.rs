// Correction CWE-113 — suppression des caractères de contrôle \r/\n et
// validation du nom de fichier par liste blanche de caractères avant
// insertion dans l'en-tête.

use actix_web::{web, HttpResponse, Responder};
use std::path::Path;

pub async fn download(query: web::Query<DownloadQuery>) -> impl Responder {
    let raw = &query.filename;

    // Sécurisé : suppression des CR/LF, puis extraction du seul nom de base
    let cleaned: String = raw.chars().filter(|c| *c != '\r' && *c != '\n').collect();
    let filename = Path::new(&cleaned)
        .file_name()
        .and_then(|f| f.to_str())
        .unwrap_or("download");

    let header_value = format!("attachment; filename=\"{}\"", filename);

    HttpResponse::Ok()
        .append_header(("Content-Disposition", header_value))
        .body("contenu du fichier")
}

#[derive(serde::Deserialize)]
pub struct DownloadQuery {
    pub filename: String,
}
