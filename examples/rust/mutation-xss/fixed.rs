// Corrigé : CWE-79 — utilisation d'une bibliothèque de sanitisation HTML
// activement maintenue (ammonia, basée sur html5ever, un parseur conforme aux
// règles de parsing du navigateur) plutôt qu'un filtrage naïf par remplacement
// de chaînes, conformément à rules/remediation/mutation-xss.md.

use actix_web::{post, web, HttpResponse, Responder};
use ammonia::Builder;
use serde::Deserialize;

#[derive(Deserialize)]
struct RichContentInput {
    html: String,
}

/// Sanitisation via une bibliothèque maintenue, tenant compte des règles
/// de parsing réelles du navigateur (pas de simple remplacement de chaînes).
fn sanitize(html: &str) -> String {
    Builder::default().clean(html).to_string()
}

#[post("/rich-content")]
async fn save_rich_content(input: web::Json<RichContentInput>) -> impl Responder {
    let clean = sanitize(&input.html);
    // Chaque nouvelle sérialisation (avant stockage ou réaffichage) doit
    // repasser par `sanitize` si le contenu est de nouveau manipulé.
    HttpResponse::Ok().json(serde_json::json!({ "stored_html": clean }))
}
