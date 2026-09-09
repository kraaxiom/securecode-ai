# Model Inversion — CWE-200 / OWASP LLM02:2025

L'endpoint d'inférence vulnérable expose un modèle fine-tuné sur des dossiers clients internes sans aucune limite de débit ni bruitage des scores de sortie, et retourne les logits/probabilités bruts. Un attaquant peut interroger l'API de façon répétée et méthodique pour reconstruire progressivement des informations mémorisées lors de l'entraînement (attaque par inversion de modèle / inférence d'appartenance), en plus d'un fine-tuning réalisé directement sur des données sensibles brutes sans anonymisation ni confidentialité différentielle.

## Pourquoi c'est dangereux

- Un modèle fine-tuné sur des données sensibles peut mémoriser des exemples d'entraînement et les restituer partiellement lors d'interrogations ciblées.
- L'absence de limite de requêtes permet un balayage systématique de l'espace d'entrée pour affiner une reconstruction.
- Les scores de confiance bruts (logits complets, non arrondis) donnent à l'attaquant la granularité nécessaire pour distinguer des exemples mémorisés d'exemples génériques.
- Aucun test de mémorisation n'est réalisé avant mise en production, donc le risque de sur-mémorisation n'est jamais mesuré.

## Correction appliquée

- **Anonymisation + confidentialité différentielle (DP-SGD)** appliquées avant tout fine-tuning sur données sensibles.
- **Test de mémorisation** obligatoire avant mise en production ; le déploiement est bloqué si le score dépasse un seuil acceptable.
- **Rate limiting** strict par clé API (50 requêtes/heure) sur l'endpoint d'inférence.
- **Détection de motifs de reconstruction** : suivi de l'historique des requêtes par clé API pour repérer un balayage systématique.
- **Arrondi des scores de sortie** : plus de logits bruts, seules des probabilités arrondies sont retournées, réduisant la précision exploitable.

## Notes résiduelles

- La détection de motifs de reconstruction est une heuristique simple ; en production, privilégier une solution de détection d'anomalies dédiée (analyse de similarité sémantique des requêtes, pas seulement de leur unicité textuelle).
- Le rate limiting doit être complété par une limite au niveau utilisateur/IP en plus de la clé API pour éviter la rotation de clés.
- Privilégier des données synthétiques à la place de données clients réelles chaque fois que c'est possible pour l'entraînement.
