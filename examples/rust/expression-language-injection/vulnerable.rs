// CWE-917 — Injection de langage d'expression (Expression Language Injection)
// Le nom fourni par l'utilisateur est concaténé dans la source d'un template
// avant compilation, permettant l'injection de directives de template
// évaluées côté serveur.

use actix_web::{web, HttpResponse, Responder};
use tera::Tera;

pub async fn greet(query: web::Query<GreetQuery>) -> impl Responder {
    let nom = &query.nom;

    // Vulnérable : la source du template est construite par concaténation
    let tpl_source = format!("Bonjour {}", nom);
    let mut tera = Tera::default();
    tera.add_raw_template("greet", &tpl_source).ok();

    let rendered = tera
        .render("greet", &tera::Context::new())
        .unwrap_or_default();

    HttpResponse::Ok().body(rendered)
}

#[derive(serde::Deserialize)]
pub struct GreetQuery {
    pub nom: String,
}
