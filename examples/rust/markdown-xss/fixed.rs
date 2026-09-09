// Corrigé : CWE-79 — le HTML brut inline est désactivé côté parseur Markdown
// (pas d'option pour le réactiver ici) et le HTML généré est systématiquement
// sanitisé via une liste blanche de balises/attributs avant insertion,
// conformément à rules/remediation/markdown-xss.md.

use actix_web::{post, web, HttpResponse, Responder};
use ammonia::Builder;
use pulldown_cmark::{html, Options, Parser};
use serde::Deserialize;

#[derive(Deserialize)]
struct CommentInput {
    body_markdown: String,
}

#[post("/comments")]
async fn render_comment(input: web::Json<CommentInput>) -> impl Responder {
    // Le Markdown ne doit pas être configuré pour interpréter du HTML brut
    // non fiable (aucune extension HTML activée ici).
    let options = Options::empty();
    let parser = Parser::new_ext(&input.body_markdown, options);

    let mut html_output = String::new();
    html::push_html(&mut html_output, parser);

    // Sanitisation du HTML généré : liste blanche de balises, schémas d'URL
    // restreints à http/https/mailto.
    let clean_html = Builder::default()
        .link_rel(Some("noopener noreferrer"))
        .url_schemes(std::collections::HashSet::from(["http", "https", "mailto"]))
        .clean(&html_output)
        .to_string();

    HttpResponse::Ok()
        .content_type("text/html; charset=utf-8")
        .body(format!("<div class='comment'>{}</div>", clean_html))
}
