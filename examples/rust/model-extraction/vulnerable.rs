// CWE-200 : Exposure of Sensitive Information to an Unauthorized Actor
// L'endpoint d'inférence est exposé sans limite de débit ni quota par
// client, et renvoie les logits complets du modèle — des informations
// riches non nécessaires au cas d'usage métier qui facilitent la
// reconstitution d'un modèle de substitution.

use actix_web::{web, HttpResponse};
use serde::Deserialize;

#[derive(Deserialize)]
struct RequeteInference {
    texte: String,
}

async fn endpoint_inference(
    payload: web::Json<RequeteInference>,
    modele: web::Data<ClientModele>,
) -> HttpResponse {
    // Vulnérable : aucune limite de débit ni quota appliqué par clé API —
    // un client peut interroger massivement et systématiquement le modèle.
    let resultat = modele.inferer(&payload.texte).await;

    // Vulnérable : les logits bruts et scores de confiance complets sont
    // renvoyés, alors que le cas d'usage n'a besoin que de la classe
    // prédite — cette richesse facilite la distillation du modèle.
    HttpResponse::Ok().json(serde_json::json!({
        "classe_predite": resultat.classe,
        "logits": resultat.logits_bruts,
        "probabilites_completes": resultat.probabilites,
    }))
}

struct ResultatInference {
    classe: String,
    logits_bruts: Vec<f32>,
    probabilites: Vec<f32>,
}

struct ClientModele;
impl ClientModele {
    async fn inferer(&self, _texte: &str) -> ResultatInference {
        ResultatInference {
            classe: "positif".to_string(),
            logits_bruts: vec![2.31, -1.05, 0.42],
            probabilites: vec![0.81, 0.05, 0.14],
        }
    }
}
