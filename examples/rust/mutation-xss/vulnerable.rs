// Vulnérable : CWE-79 — XSS par mutation (mXSS)
// Le contenu d'un éditeur riche est sanitisé une seule fois côté serveur via
// un simple filtrage de balises (strip naïf), puis stocké et réaffiché sans
// nouvelle passe de sanitisation lors de sérialisations ultérieures — exposant
// l'application aux quirks de reparsing du navigateur.

use actix_web::{post, web, HttpResponse, Responder};
use serde::Deserialize;

#[derive(Deserialize)]
struct RichContentInput {
    html: String,
}

/// Sanitisation naïve : ne retire que quelques balises explicites,
/// ne tient pas compte des mutations de parsing du navigateur.
fn naive_strip(html: &str) -> String {
    html.replace("<script>", "").replace("</script>", "")
}

#[post("/rich-content")]
async fn save_rich_content(input: web::Json<RichContentInput>) -> impl Responder {
    let clean = naive_strip(&input.html);
    // Stocké tel quel, réaffiché plus tard via innerHTML côté client sans
    // re-sanitisation lors du prochain aller-retour DOM <-> chaîne.
    HttpResponse::Ok().json(serde_json::json!({ "stored_html": clean }))
}
