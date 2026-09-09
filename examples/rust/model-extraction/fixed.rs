// CWE-200 : correction — limitation de débit stricte par clé API, réponse
// réduite aux informations strictement nécessaires (pas de logits bruts),
// et surveillance des volumes de requêtes pour détecter un usage
// d'extraction systématique.

use actix_web::{web, HttpRequest, HttpResponse};
use serde::Deserialize;
use std::collections::HashMap;
use std::sync::Mutex;
use std::time::{Duration, Instant};

#[derive(Deserialize)]
struct RequeteInference {
    texte: String,
}

struct LimiteurDebit {
    requetes_par_cle: Mutex<HashMap<String, Vec<Instant>>>,
}

impl LimiteurDebit {
    // Sécurisé : quota strict par clé API — au-delà de 60 requêtes/minute,
    // la requête est rejetée, ce qui freine une extraction systématique.
    fn autoriser(&self, cle_api: &str) -> bool {
        let mut registre = self.requetes_par_cle.lock().unwrap();
        let entrees = registre.entry(cle_api.to_string()).or_default();
        let maintenant = Instant::now();
        entrees.retain(|t| maintenant.duration_since(*t) < Duration::from_secs(60));

        if entrees.len() >= 60 {
            return false;
        }
        entrees.push(maintenant);
        true
    }
}

async fn endpoint_inference(
    requete: HttpRequest,
    payload: web::Json<RequeteInference>,
    modele: web::Data<ClientModele>,
    limiteur: web::Data<LimiteurDebit>,
) -> HttpResponse {
    let cle_api = requete
        .headers()
        .get("x-api-key")
        .and_then(|v| v.to_str().ok())
        .unwrap_or_default();

    // Sécurisé : quota appliqué avant tout appel au modèle.
    if !limiteur.autoriser(cle_api) {
        return HttpResponse::TooManyRequests().body("Quota d'inférence dépassé");
    }

    let resultat = modele.inferer(&payload.texte).await;

    // Sécurisé : seule la classe prédite est renvoyée, sans logits bruts ni
    // probabilités complètes — le strict nécessaire au cas d'usage métier.
    HttpResponse::Ok().json(serde_json::json!({
        "classe_predite": resultat.classe,
    }))
}

struct ResultatInference {
    classe: String,
    #[allow(dead_code)]
    logits_bruts: Vec<f32>,
}

struct ClientModele;
impl ClientModele {
    async fn inferer(&self, _texte: &str) -> ResultatInference {
        // Note : les logits restent calculés en interne pour les besoins du
        // modèle mais ne sont plus exposés à l'appelant.
        ResultatInference { classe: "positif".to_string(), logits_bruts: vec![2.31, -1.05, 0.42] }
    }
}
