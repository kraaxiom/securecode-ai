// CWE-1333 : correction — expression régulière non ambiguë (sans quantificateurs
// imbriqués) et limite de longueur imposée sur l'entrée avant application de la
// regex, conformément au guide de remédiation (défense en profondeur même avec
// un moteur à complexité linéaire garantie).

use actix_web::{web, HttpResponse};
use regex::Regex;

#[derive(serde::Deserialize)]
struct EmailForm {
    email: String,
}

const MAX_EMAIL_LEN: usize = 254;

async fn validate_email_handler(form: web::Form<EmailForm>) -> HttpResponse {
    let email = form.email.trim();

    // Limite de longueur imposée en amont de la regex
    if email.is_empty() || email.len() > MAX_EMAIL_LEN {
        return HttpResponse::BadRequest().body("Email invalide");
    }

    // Regex non ambiguë : aucun quantificateur imbriqué
    let re = Regex::new(r"^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$").unwrap();

    if re.is_match(email) {
        HttpResponse::Ok().body("Email valide")
    } else {
        HttpResponse::BadRequest().body("Email invalide")
    }
}
