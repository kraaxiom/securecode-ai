// Corrigé : CWE-79 — encodage de sortie contextuel systématique de toute
// donnée persistée d'origine utilisateur, conformément à
// rules/remediation/stored-xss.md.

use actix_web::{get, web, HttpResponse, Responder};

struct UserProfile {
    id: u32,
    bio: String,
}

async fn fetch_user(id: u32) -> UserProfile {
    UserProfile { id, bio: "présentation personnelle de l'utilisateur".to_string() }
}

fn escape_html(input: &str) -> String {
    input
        .replace('&', "&amp;")
        .replace('<', "&lt;")
        .replace('>', "&gt;")
        .replace('"', "&quot;")
        .replace('\'', "&#39;")
}

#[get("/profile/{id}")]
async fn view_profile(path: web::Path<u32>) -> impl Responder {
    let user = fetch_user(path.into_inner()).await;

    // Donnée stockée échappée avant affichage, même provenant de la base.
    let body = format!(
        "<html><body><div class='bio'>{}</div></body></html>",
        escape_html(&user.bio)
    );

    HttpResponse::Ok().content_type("text/html; charset=utf-8").body(body)
}
