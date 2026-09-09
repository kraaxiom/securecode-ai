// Correction CWE-444 — rejet explicite en défense en profondeur de toute
// requête présentant simultanément Content-Length et Transfer-Encoding.
// La normalisation stricte au niveau du proxy reste indispensable, mais
// l'application applicative doit aussi refuser ces requêtes ambiguës.

use actix_web::{web, HttpRequest, HttpResponse, Responder};

pub async fn handle_api(req: HttpRequest, body: web::Bytes) -> impl Responder {
    let has_length = req.headers().contains_key("content-length");
    let has_encoding = req.headers().contains_key("transfer-encoding");

    // Sécurisé : rejet explicite de la requête ambiguë
    if has_length && has_encoding {
        return HttpResponse::BadRequest()
            .body("Requête ambiguë refusée (Content-Length et Transfer-Encoding conjoints)");
    }

    HttpResponse::Ok().body(body)
}
