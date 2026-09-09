// Vulnérable : CWE-79 — XSS aveugle (Blind XSS)
// Le message d'un ticket de support soumis par un utilisateur externe non authentifié
// est réaffiché tel quel dans le back-office admin, sans encodage de sortie.
// L'admin qui consulte ce ticket exécute alors tout script injecté par l'attaquant,
// dans un contexte que celui-ci ne peut observer directement.

use actix_web::{get, web, HttpResponse, Responder};

struct Ticket {
    id: u32,
    message: String,
}

async fn fetch_ticket(id: u32) -> Ticket {
    // Simule une lecture en base de données
    Ticket { id, message: "contenu soumis par un utilisateur externe".to_string() }
}

#[get("/admin/tickets/{id}")]
async fn view_ticket(path: web::Path<u32>) -> impl Responder {
    let ticket = fetch_ticket(path.into_inner()).await;

    // Concaténation directe du champ externe dans le HTML admin, sans échappement.
    let body = format!(
        "<html><body><div class='ticket-message'>{}</div></body></html>",
        ticket.message
    );

    HttpResponse::Ok().content_type("text/html; charset=utf-8").body(body)
}
