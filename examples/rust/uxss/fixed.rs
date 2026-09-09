// Corrigé : CWE-79 — CSP stricte, Permissions-Policy, et iframe tierce isolée
// via un attribut sandbox minimal, conformément à rules/remediation/uxss.md.
// La cause première d'un UXSS est souvent hors du code applicatif, mais
// l'application reste responsable de limiter sa surface d'exposition.

use actix_web::{get, HttpResponse, Responder};

#[get("/dashboard")]
async fn dashboard() -> impl Responder {
    let body = r#"
        <html>
        <body>
          <iframe
            src="https://widget-tiers.example.com/chat"
            sandbox="allow-scripts allow-same-origin"
            referrerpolicy="no-referrer">
          </iframe>
        </body>
        </html>
    "#;

    HttpResponse::Ok()
        .content_type("text/html; charset=utf-8")
        .insert_header((
            "Content-Security-Policy",
            "default-src 'self'; script-src 'self'; frame-src https://widget-tiers.example.com",
        ))
        .insert_header((
            "Permissions-Policy",
            "camera=(), microphone=(), geolocation=()",
        ))
        .body(body)
}
