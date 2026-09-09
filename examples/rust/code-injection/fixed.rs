// Correction CWE-94 — suppression totale de l'évaluation dynamique de code.
// L'opération est restreinte à une liste blanche de fonctions autorisées,
// aucune chaîne fournie par l'utilisateur n'est interprétée comme du code.

use actix_web::{web, HttpResponse, Responder};

pub async fn compute(form: web::Json<FormulaRequest>) -> impl Responder {
    // Sécurisé : mapping déclaratif d'opérations autorisées, pas d'eval
    let result = match form.op.as_str() {
        "add" => form.a + form.b,
        "sub" => form.a - form.b,
        "mul" => form.a * form.b,
        "div" if form.b != 0.0 => form.a / form.b,
        _ => return HttpResponse::BadRequest().json(serde_json::json!({
            "error": "Opération non autorisée"
        })),
    };

    HttpResponse::Ok().json(serde_json::json!({ "result": result }))
}

#[derive(serde::Deserialize)]
pub struct FormulaRequest {
    pub a: f64,
    pub op: String,
    pub b: f64,
}
