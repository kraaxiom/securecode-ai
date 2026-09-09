// Correction CWE-1236 — chaque cellule dont le premier caractère est un
// déclencheur de formule (=, +, -, @, tab, CR) est préfixée d'une
// apostrophe avant écriture, neutralisant l'interprétation par le tableur.

use actix_web::{web, HttpResponse, Responder};
use csv::Writer;

fn sanitize_csv_cell(value: &str) -> String {
    match value.chars().next() {
        Some('=') | Some('+') | Some('-') | Some('@') | Some('\t') | Some('\r') => {
            format!("'{}", value)
        }
        _ => value.to_string(),
    }
}

pub async fn export_csv(rows: web::Json<Vec<UserRow>>) -> impl Responder {
    let mut wtr = Writer::from_writer(vec![]);

    // Sécurisé : neutralisation systématique avant écriture
    for row in rows.iter() {
        let name = sanitize_csv_cell(&row.name);
        let comment = sanitize_csv_cell(&row.comment);
        wtr.write_record(&[name, comment]).ok();
    }

    let data = wtr.into_inner().unwrap_or_default();
    HttpResponse::Ok().content_type("text/csv").body(data)
}

#[derive(serde::Deserialize)]
pub struct UserRow {
    pub name: String,
    pub comment: String,
}
