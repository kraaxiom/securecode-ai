# CWE-200 — Exposure of Sensitive Information to an Unauthorized Actor
# OWASP LLM02:2025 — Sensitive Information Disclosure (Model Inversion)
#
# Correction : anonymisation + confidentialité différentielle avant
# entraînement, test de mémorisation avant mise en production, et côté
# inférence : rate limiting strict, arrondi/bruitage des scores de sortie,
# et détection de motifs de reconstruction progressive.

import json

from django.core.cache import cache
from django.http import JsonResponse, HttpResponseForbidden
from django.views.decorators.csrf import csrf_exempt

from .ml import internal_model
from .privacy import anonymize_pii, dp_sgd_optimizer, run_memorization_test
from .audit import audit_log

MAX_ACCEPTABLE_MEMORIZATION = 0.15
RATE_LIMIT_PER_HOUR = 50
CONFIDENCE_ROUNDING = 2  # arrondi des scores pour limiter la précision exploitable


class ModelMemorizationTooHigh(Exception):
    pass


def fine_tune(model, sensitive_customer_dataset):
    """Fine-tuning avec anonymisation, confidentialité différentielle et
    test de mémorisation obligatoire avant mise en production."""
    # Correction : anonymisation des PII avant tout entraînement.
    anonymized = anonymize_pii(sensitive_customer_dataset)

    # Correction : entraînement avec bruit différentiel (DP-SGD) pour
    # empêcher la mémorisation exacte d'exemples individuels.
    trained = model.fit(anonymized, optimizer=dp_sgd_optimizer(noise_multiplier=1.1))

    # Correction : audit de mémorisation avant déploiement.
    memorization_score = run_memorization_test(trained, sensitive_customer_dataset)
    if memorization_score > MAX_ACCEPTABLE_MEMORIZATION:
        raise ModelMemorizationTooHigh(memorization_score)

    return trained


def _rate_limit_key(request):
    return f"model-inversion:infer:{request.headers.get('X-API-Key', 'anonymous')}"


def _detects_reconstruction_pattern(api_key, user_input):
    """Heuristique simple de détection de balayage systématique de
    l'espace d'entrée (motif typique d'une attaque par inversion)."""
    history_key = f"model-inversion:history:{api_key}"
    history = cache.get(history_key, [])
    history.append(user_input)
    cache.set(history_key, history[-200:], timeout=3600)
    # Signal simple : trop de variations proches dans une fenêtre courte.
    return len(set(history)) > 150


@csrf_exempt
def infer(request):
    """
    Endpoint d'inférence durci :
    - rate limiting par clé API,
    - détection de motifs de reconstruction,
    - scores de sortie arrondis (pas de logits bruts).
    """
    api_key = request.headers.get("X-API-Key")
    if not api_key:
        return HttpResponseForbidden("Clé API requise")

    # Correction : limitation stricte du volume de requêtes (50/heure).
    cache_key = _rate_limit_key(request)
    request_count = cache.get(cache_key, 0)
    if request_count >= RATE_LIMIT_PER_HOUR:
        audit_log.record("rate_limit_exceeded", api_key)
        return JsonResponse({"error": "Usage anormal détecté"}, status=429)
    cache.set(cache_key, request_count + 1, timeout=3600)

    payload = json.loads(request.body)
    user_input = payload["input"]

    # Correction : détection de motifs de reconstruction progressive.
    if _detects_reconstruction_pattern(api_key, user_input):
        audit_log.record("suspected_model_inversion", api_key)
        return JsonResponse({"error": "Usage anormal détecté"}, status=429)

    result = internal_model.predict_proba(user_input)

    # Correction : pas de logits bruts, scores arrondis/bruités pour limiter
    # la précision exploitable dans une attaque par inversion.
    rounded_scores = [round(p, CONFIDENCE_ROUNDING) for p in result.probabilities.tolist()]

    return JsonResponse({
        "confidence_scores": rounded_scores,
        "prediction": result.top_label,
    })
