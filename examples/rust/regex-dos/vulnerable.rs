// CWE-1333 : Inefficient Regular Expression Complexity
// L'expression régulière utilise des quantificateurs imbriqués ambigus
// ("(a+)+") sur une entrée non bornée, provoquant un temps de correspondance
// exponentiel (backtracking catastrophique) pour certaines entrées malveillantes.

use actix_web::{web, HttpResponse};
use regex::Regex;

#[derive(serde::Deserialize)]
struct EmailForm {
    email: String,
}

async fn validate_email_handler(form: web::Form<EmailForm>) -> HttpResponse {
    // Vulnérable : quantificateurs imbriqués + aucune limite de longueur sur l'entrée
    let re = Regex::new(r"^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$").unwrap();

    if re.is_match(&form.email) {
        HttpResponse::Ok().body("Email valide")
    } else {
        HttpResponse::BadRequest().body("Email invalide")
    }
}
