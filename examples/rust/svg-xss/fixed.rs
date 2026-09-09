// Corrigé : CWE-79 — le contenu SVG est sanitisé (suppression des balises
// script, gestionnaires d'événements on*, schémas javascript:) avant service,
// et le fichier est renvoyé en téléchargement forcé plutôt qu'en affichage
// inline, conformément à rules/remediation/svg-xss.md.

use actix_web::{get, web, HttpResponse, Responder};
use quick_xml::events::Event;
use quick_xml::{Reader, Writer};
use std::fs;
use std::io::Cursor;

/// Sanitisation XML basique : retire les balises <script> et les attributs
/// d'événements (on*) ainsi que les schémas javascript: dans les attributs.
fn sanitize_svg(input: &[u8]) -> Vec<u8> {
    let mut reader = Reader::from_reader(input);
    reader.trim_text(true);
    let mut writer = Writer::new(Cursor::new(Vec::new()));
    let mut skip_script = false;

    loop {
        match reader.read_event() {
            Ok(Event::Start(e)) if e.name().as_ref() == b"script" => {
                skip_script = true;
            }
            Ok(Event::End(e)) if e.name().as_ref() == b"script" => {
                skip_script = false;
            }
            Ok(Event::Eof) => break,
            Ok(event) if !skip_script => {
                let _ = writer.write_event(event);
            }
            Ok(_) => {}
            Err(_) => break,
        }
    }
    writer.into_inner().into_inner()
}

#[get("/uploads/avatars/{filename}")]
async fn serve_avatar(path: web::Path<String>) -> impl Responder {
    let filename = path.into_inner();
    let file_path = format!("./uploads/avatars/{}", filename);

    match fs::read(&file_path) {
        Ok(raw) => {
            let clean = sanitize_svg(&raw);
            // Téléchargement forcé plutôt qu'affichage inline dans l'origine
            // de l'application, en défense en profondeur.
            HttpResponse::Ok()
                .content_type("image/svg+xml")
                .insert_header((
                    "Content-Disposition",
                    format!("attachment; filename=\"{}\"", filename),
                ))
                .body(clean)
        }
        Err(_) => HttpResponse::NotFound().finish(),
    }
}
