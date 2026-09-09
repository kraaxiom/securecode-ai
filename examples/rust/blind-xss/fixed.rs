// Corrigé : CWE-79 — encodage de sortie contextuel systématique, y compris
// dans les interfaces internes/admin. Toute donnée externe reste non fiable
// quel que soit l'écran où elle est affichée (voir rules/remediation/blind-xss.md).

use actix_web::{get, web, HttpResponse, Responder};

struct Ticket {
    id: u32,
    message: String,
}

async fn fetch_ticket(id: u32) -> Ticket {
    Ticket { id, message: "contenu soumis par un utilisateur externe".to_string() }
}

/// Échappement HTML manuel minimal (équivalent htmlspecialchars ENT_QUOTES).
fn escape_html(input: &str) -> String {
    input
        .replace('&', "&amp;")
        .replace('<', "&lt;")
        .replace('>', "&gt;")
        .replace('"', "&quot;")
        .replace('\'', "&#39;")
}

#[get("/admin/tickets/{id}")]
async fn view_ticket(path: web::Path<u32>) -> impl Responder {
    let ticket = fetch_ticket(path.into_inner()).await;

    // La donnée externe passe par un encodage HTML avant insertion, même
    // affichée dans une interface interne/admin.
    let body = format!(
        "<html><body><div class='ticket-message'>{}</div></body></html>",
        escape_html(&ticket.message)
    );

    HttpResponse::Ok().content_type("text/html; charset=utf-8").body(body)
}
