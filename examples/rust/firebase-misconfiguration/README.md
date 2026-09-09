# Mauvaise configuration Firebase (CWE-284)

Le service backend déploie, via l'API Firebase Rules, des règles Firestore en mode test (`allow read, write: if true;`), rendant toute la base de données lisible et modifiable par quiconque sans authentification. La correction déploie des règles exigeant `request.auth != null` combiné à une vérification de propriété (`request.auth.uid == userId`), et valide le schéma des champs autorisés en écriture (`hasOnly`). Élimine la classe de vulnérabilité CWE-284 (Improper Access Control).
