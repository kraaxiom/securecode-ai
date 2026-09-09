# CWE-200 — Exposure of Sensitive Information to an Unauthorized Actor
# OWASP LLM02:2025 — Sensitive Information Disclosure (Model Inversion)
#
# Vue Django exposant une API d'inférence pour un modèle fine-tuné sur des
# dossiers clients internes (données sensibles). Aucune limitation de débit,
# aucun bruitage des scores de sortie, aucune anonymisation en amont :
# un attaquant peut interroger le modèle de façon répétée et méthodique
# pour reconstruire progressivement des données mémorisées (membership
# inference / model inversion).

import json

from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt

from .ml import internal_model  # modèle fine-tuné sur données clients brutes


def fine_tune(model, sensitive_customer_dataset):
    """Fine-tuning direct sur données sensibles, sans anonymisation ni DP."""
    # Vulnérable : aucune anonymisation, aucune confidentialité différentielle,
    # aucun test de mémorisation avant mise en production.
    return model.fit(sensitive_customer_dataset)


@csrf_exempt
def infer(request):
    """
    Endpoint d'inférence exposant les scores de confiance bruts (logits).

    Vulnérable :
    - Pas de rate limiting -> volume de requêtes illimité par client.
    - Retourne les logits/probabilités complètes -> permet une attaque
      par inférence d'appartenance (membership inference) affinée requête
      après requête.
    - Aucune surveillance de motif de requête suspect (balayage systématique
      de l'espace d'entrée).
    """
    payload = json.loads(request.body)
    user_input = payload["input"]

    # Aucune limite de volume ni de diversité de requêtes.
    result = internal_model.predict_proba(user_input)

    # Fuite : scores de confiance complets, non arrondis, non bruités.
    # Ces valeurs précises facilitent la reconstruction de données
    # d'entraînement mémorisées (ex. reconstruction d'un profil client).
    return JsonResponse({
        "logits": result.logits.tolist(),
        "confidence_scores": result.probabilities.tolist(),
        "prediction": result.top_label,
    })
