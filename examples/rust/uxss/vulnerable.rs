// Vulnérable : CWE-79 — surface d'exposition UXSS
// L'application sert une page intégrant un widget tiers via une iframe sans
// attribut sandbox et sans Content-Security-Policy ni Permissions-Policy,
// élargissant la surface d'exposition à une faille du composant tiers.

use actix_web::{get, HttpResponse, Responder};

#[get("/dashboard")]
async fn dashboard() -> impl Responder {
    // Aucun en-tête de sécurité émis, iframe tierce sans sandbox.
    let body = r#"
        <html>
        <body>
          <iframe src="https://widget-tiers.example.com/chat"></iframe>
        </body>
        </html>
    "#;

    HttpResponse::Ok().content_type("text/html; charset=utf-8").body(body)
}
