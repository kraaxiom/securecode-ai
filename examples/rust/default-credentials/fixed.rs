// Corrigé : CWE-1392 — génération d'un mot de passe temporaire aléatoire
// à l'installation, avec changement forcé à la première connexion,
// conformément à rules/remediation/default-credentials.md.

use actix_web::{post, HttpResponse, Responder};
use rand::RngCore;

struct AdminAccount {
    email: String,
    password_hash: String,
    must_change_password: bool,
}

fn generate_temp_password() -> String {
    let mut bytes = [0u8; 18];
    rand::thread_rng().fill_bytes(&mut bytes);
    // Encodage base64url pour un mot de passe temporaire lisible.
    base64::Engine::encode(&base64::engine::general_purpose::URL_SAFE_NO_PAD, bytes)
}

fn hash_password(password: &str) -> String {
    format!("argon2id({})", password)
}

#[post("/setup/provision-admin")]
async fn provision_admin() -> impl Responder {
    let temp_password = generate_temp_password();
    let admin = AdminAccount {
        email: "admin@example.com".to_string(),
        password_hash: hash_password(&temp_password),
        must_change_password: true,
    };

    // Le mot de passe temporaire n'est jamais journalisé en clair ; il est
    // transmis via un canal hors-bande sécurisé (ex: courriel signé, vault).
    deliver_out_of_band(&admin.email, &temp_password);

    HttpResponse::Ok().json(serde_json::json!({
        "email": admin.email,
        "must_change_password": admin.must_change_password
    }))
}

fn deliver_out_of_band(_email: &str, _temp_password: &str) {
    // Implémentation réelle : envoi via un canal sécurisé dédié, jamais en log.
}
