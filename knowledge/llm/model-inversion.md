---
id: model-inversion
category: llm
cwe: CWE-200
owasp: LLM02:2025-Sensitive-Information-Disclosure
severity_default: high
languages: [python, js, java, csharp, go]
---

# Model Inversion

## Description
L'inversion de modèle est une technique par laquelle un attaquant, en interrogeant de façon répétée et méthodique un modèle, parvient à reconstruire des informations sensibles présentes dans ses données d'entraînement (données personnelles, secrets, documents propriétaires) même sans y avoir un accès direct. C'est un risque particulièrement critique pour les modèles fine-tunés sur des données internes ou personnelles, où le modèle peut mémoriser et restituer partiellement des exemples d'entraînement.

## Où ça apparaît typiquement
- Modèles fine-tunés sur des données internes sensibles (dossiers clients, données médicales, code propriétaire) exposés via une API.
- Absence de technique de confidentialité différentielle (differential privacy) lors de l'entraînement sur données sensibles.
- APIs de modèle ne limitant pas le nombre ou la diversité des requêtes permettant une reconstruction progressive.
- Modèles n'ayant pas fait l'objet d'un audit de mémorisation avant mise en production.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fine-tuning sur des données sensibles sans application de techniques de confidentialité différentielle ni de dédoublonnage préalable.
- Absence de test de mémorisation (memorization testing) avant le déploiement d'un modèle entraîné sur des données propriétaires.
- Aucune limitation sur le volume ou la structure des requêtes permettant des attaques par interrogation répétée.
- Réponses du modèle contenant occasionnellement des extraits très proches de documents d'entraînement identifiables (signal de sur-mémorisation).

## Remédiation
- Appliquer des techniques de confidentialité différentielle ou d'anonymisation sur les données sensibles avant tout entraînement/fine-tuning.
- Réaliser des tests de mémorisation avant mise en production pour détecter la restitution de données d'entraînement.
- Limiter et surveiller les volumes de requêtes permettant une reconstruction progressive de données mémorisées.
- Minimiser l'usage de données réelles sensibles dans l'entraînement au profit de données synthétiques quand c'est possible.
- Voir `rules/remediation/model-inversion.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/model-inversion/`.

## Références
- OWASP Top 10 for LLM Applications: LLM02:2025 – Sensitive Information Disclosure
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
