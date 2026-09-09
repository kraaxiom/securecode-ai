// Vulnérable : CWE-1392 — utilisation d'identifiants par défaut
// Le compte administrateur est provisionné avec un mot de passe fixe et
// documenté, sans forcer son changement à la première connexion.

use actix_web::{post, HttpResponse, Responder};

struct AdminAccount {
    email: String,
    password_hash: String,
}

fn hash_password(password: &str) -> String {
    // Simule un hachage (non représentatif d'un vrai algorithme).
    format!("hash({})", password)
}

#[post("/setup/provision-admin")]
async fn provision_admin() -> impl Responder {
    // Mot de passe par défaut fixe, identique à chaque installation.
    let admin = AdminAccount {
        email: "admin@example.com".to_string(),
        password_hash: hash_password("admin123"),
    };

    HttpResponse::Ok().json(serde_json::json!({
        "email": admin.email,
        "provisioned": true
    }))
}
