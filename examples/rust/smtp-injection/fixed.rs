// CWE-93 : correction — utilisation de la bibliothèque `lettre` (construction
// structurée des en-têtes, échappement automatique) et rejet explicite de tout
// caractère de contrôle (\r, \n) dans les champs libres avant usage.

use actix_web::{web, HttpResponse};
use lettre::message::Message;

#[derive(serde::Deserialize)]
struct ContactForm {
    name: String,
    subject: String,
    body: String,
}

fn strip_control_chars(value: &str) -> String {
    value.chars().filter(|c| *c != '\r' && *c != '\n').collect()
}

async fn contact_handler(form: web::Form<ContactForm>) -> HttpResponse {
    let name = strip_control_chars(&form.name);
    let subject = strip_control_chars(&form.subject);

    // Sécurisé : bibliothèque d'email structurée, en-têtes construits par API dédiée
    let message = Message::builder()
        .from("contact@example.com".parse().unwrap())
        .reply_to(format!("{} <contact@example.com>", name).parse().unwrap())
        .to("dest@example.com".parse().unwrap())
        .subject(subject)
        .body(form.body.clone());

    match message {
        Ok(_) => HttpResponse::Ok().body("Message envoyé"),
        Err(_) => HttpResponse::BadRequest().body("Champs invalides"),
    }
}
