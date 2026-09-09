# Mauvaise configuration Firebase — Python/Django

`vulnerable.py` initialise le SDK Admin Firebase avec une clé de service codée en dur dans le code source, dans un contexte où les règles de sécurité Firestore associées sont restées en mode test (`allow read, write: if true;`). Combiné à l'absence de contrôle d'autorisation côté backend, n'importe quel utilisateur — authentifié ou non — peut lire ou modifier les données de n'importe quel autre utilisateur (CWE-284, Improper Access Control).

`fixed.py` charge la clé de service depuis un gestionnaire de secrets via variable d'environnement, ajoute un contrôle d'autorisation explicite côté backend (l'utilisateur ne peut accéder qu'à son propre profil), et documente les règles Firestore restrictives (auth + propriété du document) à déployer en parallèle.

## Pourquoi c'est dangereux
- Le mode test Firebase (`if true`) est pensé pour le prototypage rapide mais expose l'intégralité de la base de données si oublié en production.
- Une clé de service Admin SDK codée en dur donne un accès total (contournant les règles) à quiconque la récupère.
- L'absence de contrôle applicatif signifie que même un backend "propre" peut relayer des accès non autorisés si les règles côté client sont ouvertes.

## Explication du correctif
- Clé de service chargée depuis `os.environ` (issue d'un gestionnaire de secrets), jamais versionnée.
- Ajout d'une vérification explicite `request.user.id == user_id` avant toute lecture Firestore.
- Règles Firestore de référence limitant `read`/`write` à `request.auth != null && request.auth.uid == userId`.

## Notes résiduelles
- Les règles Firestore doivent être testées avec l'émulateur Firebase (cas autorisé + cas refusé) avant tout déploiement en production.
- Prévoir une validation de schéma sur les écritures (`hasOnly`, types, tailles) pour éviter l'injection de champs arbitraires.
