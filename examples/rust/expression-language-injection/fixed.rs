// Correction CWE-917 — le template est une source statique définie par le
// développeur ; la donnée utilisateur n'est transmise que comme variable
// de contexte de rendu, jamais insérée dans le code source du template.

use actix_web::{web, HttpResponse, Responder};
use tera::{Context, Tera};

pub async fn greet(query: web::Query<GreetQuery>) -> impl Responder {
    let nom = &query.nom;

    // Sécurisé : template statique, donnée injectée comme variable échappée
    let mut tera = Tera::default();
    tera.add_raw_template("greet", "Bonjour {{ nom }}").ok();

    let mut context = Context::new();
    context.insert("nom", nom);

    let rendered = tera.render("greet", &context).unwrap_or_default();

    HttpResponse::Ok().body(rendered)
}

#[derive(serde::Deserialize)]
pub struct GreetQuery {
    pub nom: String,
}
