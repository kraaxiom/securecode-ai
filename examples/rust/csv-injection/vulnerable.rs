// CWE-1236 — Injection CSV (CSV Injection / Formula Injection)
// Les champs utilisateur sont écrits tels quels dans le fichier CSV exporté,
// sans vérifier si la cellule commence par un caractère déclencheur de
// formule (=, +, -, @), qui sera interprété par le tableur à l'ouverture.

use actix_web::{web, HttpResponse, Responder};
use csv::Writer;

pub async fn export_csv(rows: web::Json<Vec<UserRow>>) -> impl Responder {
    let mut wtr = Writer::from_writer(vec![]);

    // Vulnérable : aucune neutralisation du premier caractère de cellule
    for row in rows.iter() {
        wtr.write_record(&[row.name.as_str(), row.comment.as_str()])
            .ok();
    }

    let data = wtr.into_inner().unwrap_or_default();
    HttpResponse::Ok().content_type("text/csv").body(data)
}

#[derive(serde::Deserialize)]
pub struct UserRow {
    pub name: String,
    pub comment: String,
}
