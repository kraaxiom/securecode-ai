# Commande de management Django vulnérable — CWE-349 (Data Poisoning /
# OWASP LLM04:2025). Le pipeline de fine-tuning ingère automatiquement des
# retours utilisateurs et des sources externes sans vérifier leur
# provenance, leur intégrité, ni filtrer les valeurs aberrantes. Un
# attaquant peut ainsi injecter des exemples empoisonnés (biais, backdoor
# comportementale) qui seront intégrés durablement au modèle.

import requests
from django.core.management.base import BaseCommand

from feedback.models import ConversationFeedback


class Command(BaseCommand):
    help = "Construit le jeu de fine-tuning et lance l'entraînement"

    def handle(self, *args, **options):
        dataset = []

        # Vulnérable : sources externes récupérées sans vérification de
        # provenance ni d'intégrité (pas de hash/signature attendue).
        external_sources = [
            "https://community-dataset-mirror.example/export.jsonl",
            "https://partner-feed.example/labels.jsonl",
        ]
        for url in external_sources:
            resp = requests.get(url, timeout=30)
            dataset.extend(_parse_jsonl(resp.text))

        # Vulnérable : le feedback utilisateur brut est réintégré tel quel
        # dans l'entraînement, sans échantillonnage humain ni filtrage.
        for feedback in ConversationFeedback.objects.filter(used_in_training=False):
            dataset.append({"prompt": feedback.prompt, "completion": feedback.corrected_response})
            feedback.used_in_training = True
            feedback.save()

        # Vulnérable : aucune détection d'anomalies statistiques, aucun
        # jeu de référence (golden set) pour détecter une dérive après
        # l'entraînement.
        finetune_model(dataset)


def _parse_jsonl(text):
    import json
    return [json.loads(line) for line in text.splitlines() if line.strip()]


def finetune_model(dataset):
    # Simule l'appel à l'API de fine-tuning
    ...
