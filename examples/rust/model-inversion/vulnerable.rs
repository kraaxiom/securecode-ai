// CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
// OWASP LLM02:2025 - Sensitive Information Disclosure
//
// Scénario vulnérable : une API d'inférence expose les logprobs bruts avec
// une haute confiance, et ne limite ni ne surveille le volume de requêtes.
// Un modèle fine-tuné sur des données internes sensibles peut ainsi être
// interrogé de façon répétée et méthodique pour reconstruire progressivement
// des fragments de ses données d'entraînement (attaque par inversion de modèle).

use async_openai::{types::CreateCompletionRequestArgs, Client};
use axum::{extract::Json, routing::post, Router};
use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
struct InferRequest {
    input: String,
}

#[derive(Serialize)]
struct InferResponse {
    text: String,
    // VULNÉRABLE : les logprobs complets (top-k, haute précision) sont renvoyés
    // tels quels. Ces valeurs facilitent grandement la reconstruction de
    // données mémorisées par le modèle (attaque par inversion / extraction).
    logprobs: Vec<f32>,
    top_logprobs: Vec<Vec<(String, f32)>>,
}

// VULNÉRABLE : le modèle a été fine-tuné directement sur des données internes
// sensibles (dossiers clients, documents propriétaires) sans anonymisation
// préalable ni confidentialité différentielle (differential privacy).
async fn fine_tune_on_raw_customer_data(client: &Client<async_openai::config::OpenAIConfig>) {
    // Pas d'anonymisation, pas de dp-sgd, pas de test de mémorisation avant
    // mise en production : le modèle peut mémoriser des exemples exacts.
    let _ = client; // (placeholder d'appel de fine-tuning réel)
}

// VULNÉRABLE : endpoint d'inférence sans rate limiting, sans quota par clé
// d'API, et sans détection de motifs de requêtes évoquant une reconstruction
// progressive (variations systématiques d'un même préfixe, par exemple).
async fn infer(Json(req): Json<InferRequest>) -> Json<InferResponse> {
    let client = Client::new();

    let request = CreateCompletionRequestArgs::default()
        .model("internal-finetuned-model")
        .prompt(req.input)
        .logprobs(20u8) // top-20 logprobs renvoyés en clair à l'appelant
        .build()
        .expect("build request");

    let response = client
        .completions()
        .create(request)
        .await
        .expect("completion request failed");

    let choice = &response.choices[0];

    // Aucune journalisation d'audit, aucune limite de volume, aucun scoring
    // d'anomalie sur les patterns de requêtes de cet appelant.
    Json(InferResponse {
        text: choice.text.clone(),
        logprobs: vec![], // (illustratif : les vraies valeurs viendraient de choice.logprobs)
        top_logprobs: vec![],
    })
}

fn app() -> Router {
    Router::new().route("/v1/infer", post(infer))
}
