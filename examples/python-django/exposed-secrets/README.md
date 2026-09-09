# Secrets exposés — Python/Django

`vulnerable.py` montre un `settings.py` où la `SECRET_KEY` Django, les identifiants de base de données, les clés d'accès AWS et une clé API Stripe de production sont codés en dur et donc versionnés dans le dépôt Git (CWE-798, Use of Hard-coded Credentials). Quiconque a accès au dépôt — y compris via l'historique Git après suppression — obtient un accès complet aux services concernés.

`fixed.py` charge l'ensemble de ces valeurs depuis des variables d'environnement, elles-mêmes injectées à l'exécution par un gestionnaire de secrets dédié (AWS Secrets Manager, Azure Key Vault, GCP Secret Manager, Vault). Aucun secret ne subsiste dans le code source.

## Pourquoi c'est dangereux
- Un dépôt Git rendu public par erreur, un fork accidentel ou une simple capture d'écran suffit à divulguer des identifiants de production.
- Même après suppression du code, le secret reste présent dans l'historique Git tant qu'il n'a pas été purgé.
- Une `SECRET_KEY` Django compromise permet de forger des sessions, des tokens de réinitialisation de mot de passe, et des signatures cryptographiques de l'application.

## Explication du correctif
- Remplacement de toutes les valeurs en dur par `os.environ[...]` (échoue explicitement si la variable manque, évitant un déploiement silencieusement mal configuré).
- Les secrets réels sont injectés par la plateforme d'hébergement depuis un gestionnaire de secrets, jamais stockés dans un fichier versionné.
- `.env` et fichiers de configuration sensibles doivent être ajoutés au `.gitignore`.

## Notes résiduelles
- Le correctif de code seul ne suffit pas en cas de fuite déjà survenue : il faut révoquer/régénérer immédiatement les secrets exposés auprès du fournisseur, purger l'historique Git (`git filter-repo`/BFG) et auditer les logs d'accès pendant la période d'exposition.
- Mettre en place un scan automatique de secrets (gitleaks, trufflehog) en pre-commit et en CI pour empêcher toute réintroduction future.
