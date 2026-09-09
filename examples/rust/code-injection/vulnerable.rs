// CWE-94 — Injection de code (Code Injection)
// L'opérateur choisi par l'utilisateur pilote directement une évaluation
// dynamique via un moteur d'expression générique, sans liste blanche
// stricte des opérations autorisées.

use actix_web::{web, HttpResponse, Responder};

pub async fn compute(form: web::Json<FormulaRequest>) -> impl Responder {
    // Vulnérable : la formule brute fournie par l'utilisateur est passée
    // telle quelle à un évaluateur d'expression générique (ex: crate `evalexpr`).
    let expr = format!("{} {} {}", form.a, form.op, form.b);

    match evalexpr::eval(&expr) {
        Ok(value) => HttpResponse::Ok().json(serde_json::json!({ "result": value.to_string() })),
        Err(_) => HttpResponse::BadRequest().finish(),
    }
}

#[derive(serde::Deserialize)]
pub struct FormulaRequest {
    pub a: f64,
    pub op: String,
    pub b: f64,
}
