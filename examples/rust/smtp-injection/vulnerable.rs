// CWE-93 : Improper Neutralization of CRLF Sequences ('CRLF Injection')
// Le champ `name` fourni par l'utilisateur est inséré tel quel dans les
// en-têtes de l'email construits manuellement. Un caractère de saut de ligne
// (CR/LF) permettrait d'injecter des en-têtes supplémentaires (Cc, Bcc, etc.).

use actix_web::{web, HttpResponse};

#[derive(serde::Deserialize)]
struct ContactForm {
    name: String,
    subject: String,
    body: String,
}

async fn contact_handler(form: web::Form<ContactForm>) -> HttpResponse {
    // Vulnérable : en-têtes construits par concaténation directe, sans filtrage CR/LF
    let raw_message = format!(
        "From: contact@example.com\r\nReply-To: {}\r\nSubject: {}\r\n\r\n{}",
        form.name, form.subject, form.body
    );

    // Envoi simplifié (représentation du transport SMTP bas niveau)
    match send_raw_smtp(&raw_message) {
        Ok(_) => HttpResponse::Ok().body("Message envoyé"),
        Err(_) => HttpResponse::InternalServerError().body("Échec de l'envoi"),
    }
}

fn send_raw_smtp(_message: &str) -> Result<(), ()> {
    Ok(())
}
