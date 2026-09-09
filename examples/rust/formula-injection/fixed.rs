// Correction CWE-1236 — chaque cellule dont le premier caractère est un
// déclencheur de formule (=, +, -, @, tab, CR) est neutralisée par
// préfixage d'une apostrophe avant écriture dans l'export.

use actix_web::{web, HttpResponse, Responder};
use csv::Writer;

fn neutraliser_formule(valeur: &str) -> String {
    match valeur.chars().next() {
        Some('=') | Some('+') | Some('-') | Some('@') | Some('\t') | Some('\r') => {
            format!("'{}", valeur)
        }
        _ => valeur.to_string(),
    }
}

pub async fn export_report(items: web::Json<Vec<ReportItem>>) -> impl Responder {
    let mut wtr = Writer::from_writer(vec![]);

    // Sécurisé : neutralisation systématique avant écriture
    for item in items.iter() {
        let nom = neutraliser_formule(&item.nom);
        let commentaire = neutraliser_formule(&item.commentaire);
        wtr.write_record(&[nom, commentaire]).ok();
    }

    let data = wtr.into_inner().unwrap_or_default();
    HttpResponse::Ok().content_type("text/csv").body(data)
}

#[derive(serde::Deserialize)]
pub struct ReportItem {
    pub nom: String,
    pub commentaire: String,
}
