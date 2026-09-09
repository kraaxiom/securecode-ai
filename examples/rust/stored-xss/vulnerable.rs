// Vulnérable : CWE-79 — XSS stocké
// La biographie du profil utilisateur, persistée en base, est réaffichée à
// d'autres utilisateurs sans encodage de sortie, affectant potentiellement
// tous les visiteurs de la page de profil.

use actix_web::{get, web, HttpResponse, Responder};

struct UserProfile {
    id: u32,
    bio: String,
}

async fn fetch_user(id: u32) -> UserProfile {
    // Simule une lecture en base : `bio` provient d'une saisie utilisateur
    // stockée précédemment sans validation de format.
    UserProfile { id, bio: "présentation personnelle de l'utilisateur".to_string() }
}

#[get("/profile/{id}")]
async fn view_profile(path: web::Path<u32>) -> impl Responder {
    let user = fetch_user(path.into_inner()).await;

    let body = format!(
        "<html><body><div class='bio'>{}</div></body></html>",
        user.bio
    );

    HttpResponse::Ok().content_type("text/html; charset=utf-8").body(body)
}
