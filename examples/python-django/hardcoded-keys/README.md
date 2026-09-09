# Clés cryptographiques codées en dur — Python/Django

`vulnerable.py` définit une clé de signature HMAC utilisée pour les tokens de réinitialisation de mot de passe, ainsi qu'une clé d'API de paiement, directement en littéral dans le code source (CWE-798, Use of Hard-coded Credentials). Quiconque a accès au dépôt Git, à son historique, ou au binaire déployé peut extraire ces secrets et forger des tokens valides ou débiter des cartes bancaires.

`fixed.py` charge ces mêmes secrets exclusivement depuis les variables d'environnement, avec un échec explicite (`RuntimeError`) si une variable requise est absente, et sans jamais écrire de valeur par défaut codée en dur.

## Pourquoi c'est dangereux
- Une clé de signature présente dans le code source n'est plus un secret : n'importe qui peut forger des tokens de réinitialisation de mot de passe valides pour n'importe quel compte.
- Une clé d'API de paiement exposée permet de débiter des cartes ou d'accéder au tableau de bord marchand du fournisseur.
- Le secret reste présent dans l'historique Git même après suppression du fichier, sauf purge explicite.

## Explication du correctif
- Suppression totale des littéraux de secrets du code source.
- Chargement via `os.environ`, avec une fonction utilitaire qui lève une exception explicite si la variable est absente (pas de valeur par défaut silencieuse).
- Séparation naturelle des secrets par environnement (dev/staging/prod), chaque déploiement injectant ses propres valeurs.
- En production, ces variables doivent elles-mêmes provenir d'un gestionnaire de secrets (Vault, AWS/Azure/GCP Secret Manager), pas d'un fichier `.env` en clair sur le serveur.

## Notes résiduelles
- Les clés ayant été exposées dans `vulnerable.py` (valeurs d'exemple) doivent être considérées comme compromises en conditions réelles et régénérées côté fournisseur.
- Purger l'historique Git de tout secret déjà committé (BFG Repo-Cleaner ou `git filter-repo`) si ce cas se présente réellement.
- Mettre en place une rotation périodique des clés critiques, indépendamment de toute fuite constatée.
