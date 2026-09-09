# Commande de management Django corrigée — CWE-349 (Data Poisoning /
# OWASP LLM04:2025). Correction : chaque source est identifiée et
# vérifiée (liste blanche + hash signé), les données subissent une
# détection d'anomalies statistiques avant intégration, le feedback
# utilisateur passe par un échantillonnage humain, et le modèle est
# comparé à un jeu de référence (golden set) après l'entraînement.

import json

import requests
from django.core.management.base import BaseCommand

from feedback.models import ConversationFeedback
from core.audit import audit_log
from core.integrity import verify_integrity
from core.anomaly import filter_statistical_outliers
from core.human_review import sample_for_human_review
from core.evaluation import evaluate, ModelDriftDetected

# Sécurisé : liste blanche explicite des sources de confiance, chacune
# associée à un hash signé attendu — aucune source arbitraire n'est admise.
TRUSTED_SOURCES = {
    "internal-labeled-set": {
        "url": "https://internal.example/datasets/labeled-set.jsonl",
        "signed_hash": "sha256:7c3a...",
    },
    "verified-partner-feed": {
        "url": "https://partner-feed.example/labels.jsonl",
        "signed_hash": "sha256:9e1b...",
    },
}

ACCEPTABLE_DRIFT_THRESHOLD = 0.92


class Command(BaseCommand):
    help = "Construit le jeu de fine-tuning et lance l'entraînement (avec contrôles anti-poisoning)"

    def handle(self, *args, **options):
        dataset = []

        for source_id, source in TRUSTED_SOURCES.items():
            resp = requests.get(source["url"], timeout=30)
            # Sécurisé : intégrité vérifiée avant tout traitement.
            verify_integrity(resp.content, expected_hash=source["signed_hash"])
            batch = _parse_jsonl(resp.text)
            # Sécurisé : filtrage des exemples statistiquement aberrants
            # (motifs répétitifs suspects, labels incohérents).
            batch = filter_statistical_outliers(batch)
            dataset.extend(batch)
            audit_log.record("training_source_ingested", source_id=source_id, count=len(batch))

        # Sécurisé : le feedback utilisateur n'est pas réintégré
        # automatiquement — un échantillon est soumis à revue humaine
        # avant inclusion dans le jeu d'entraînement.
        pending_feedback = list(ConversationFeedback.objects.filter(used_in_training=False))
        approved = sample_for_human_review(pending_feedback)
        for feedback in approved:
            dataset.append({"prompt": feedback.prompt, "completion": feedback.corrected_response})
            feedback.used_in_training = True
            feedback.save()
        audit_log.record("feedback_reviewed", total=len(pending_feedback), approved=len(approved))

        model = finetune_model(dataset)

        # Sécurisé : comparaison à un jeu de référence fixe pour détecter
        # une dérive de comportement introduite par un cycle empoisonné.
        golden_set = load_golden_set()
        score = evaluate(model, golden_set)
        if score < ACCEPTABLE_DRIFT_THRESHOLD:
            audit_log.record("model_drift_detected", score=score)
            raise ModelDriftDetected(score)

        audit_log.record("finetune_completed", score=score)


def _parse_jsonl(text):
    return [json.loads(line) for line in text.splitlines() if line.strip()]


def finetune_model(dataset):
    ...


def load_golden_set():
    ...
