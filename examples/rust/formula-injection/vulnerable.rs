// CWE-1236 — Injection de formule (Formula Injection)
// Le commentaire utilisateur est écrit tel quel dans le fichier CSV exporté
// vers un tableur, sans vérifier le premier caractère de la cellule.

use actix_web::{web, HttpResponse, Responder};
use csv::Writer;

pub async fn export_report(items: web::Json<Vec<ReportItem>>) -> impl Responder {
    let mut wtr = Writer::from_writer(vec![]);

    // Vulnérable : écriture directe sans neutralisation de formule
    for item in items.iter() {
        wtr.write_record(&[item.nom.as_str(), item.commentaire.as_str()])
            .ok();
    }

    let data = wtr.into_inner().unwrap_or_default();
    HttpResponse::Ok().content_type("text/csv").body(data)
}

#[derive(serde::Deserialize)]
pub struct ReportItem {
    pub nom: String,
    pub commentaire: String,
}
