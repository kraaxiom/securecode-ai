// Vulnérable : CWE-79 — XSS via rendu Markdown
// Le contenu Markdown d'un commentaire utilisateur est converti en HTML puis
// inséré tel quel dans la réponse, sans sanitisation du HTML brut potentiellement
// contenu dans le Markdown (balises actives, liens javascript:).

use actix_web::{post, web, HttpResponse, Responder};
use pulldown_cmark::{html, Options, Parser};
use serde::Deserialize;

#[derive(Deserialize)]
struct CommentInput {
    body_markdown: String,
}

#[post("/comments")]
async fn render_comment(input: web::Json<CommentInput>) -> impl Responder {
    // Options par défaut : le parser laisse passer le HTML brut inline
    // présent dans le Markdown utilisateur.
    let options = Options::empty();
    let parser = Parser::new_ext(&input.body_markdown, options);

    let mut html_output = String::new();
    html::push_html(&mut html_output, parser);

    // HTML généré inséré directement, sans sanitisation.
    HttpResponse::Ok()
        .content_type("text/html; charset=utf-8")
        .body(format!("<div class='comment'>{}</div>", html_output))
}
