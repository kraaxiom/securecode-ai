# Remédiation — Fuite de secrets dans GitLab CI

## Principe
Déclarer les secrets comme variables CI/CD "Protected" et "Masked" dans les paramètres du projet, jamais en dur dans `.gitlab-ci.yml`. Ne jamais journaliser une variable sensible, restreindre les artefacts/caches, et limiter les jobs sensibles à des runners dédiés.

## Identifiant en clair dans le pipeline
```yaml
# Avant — vulnérable
deploy:
  stage: deploy
  script:
    - curl -u admin:P@ssw0rd123 https://api.exemple.com/deploy

# Après — sécurisé
deploy:
  stage: deploy
  script:
    - curl -u admin:$DEPLOY_PASSWORD https://api.exemple.com/deploy
  rules:
    - if: '$CI_COMMIT_BRANCH == "main"'
# DEPLOY_PASSWORD défini dans Settings > CI/CD > Variables, coché "Protected" et "Masked"
```

## Script affichant une variable sensible
```yaml
# Avant — vulnérable
debug:
  script:
    - export -p   # imprime toutes les variables, y compris les secrets

# Après — sécurisé
debug:
  script:
    - echo "Pipeline $CI_PIPELINE_ID démarré pour $CI_COMMIT_REF_NAME"
# aucune commande n'expose le contenu des variables sensibles
```

## Variable non protégée accessible depuis n'importe quelle branche + artefacts à risque
```yaml
# Avant — vulnérable
# Variable PROD_API_KEY définie sans coche "Protected" -> lisible depuis une branche de contribution externe
build:
  script:
    - echo $PROD_API_KEY > config.json
  artifacts:
    paths:
      - config.json   # fichier contenant le secret publié en artefact téléchargeable

# Après — sécurisé
# Variable PROD_API_KEY marquée "Protected" (limitée aux branches/tags protégés) et "Masked"
build:
  script:
    - envsubst < config.template.json > config.json   # injection à l'exécution, jamais versionné/exposé
  artifacts:
    paths:
      - dist/   # exclut explicitement tout fichier contenant des secrets
  rules:
    - if: '$CI_COMMIT_REF_PROTECTED == "true"'
```

## Checklist de vérification post-patch
- [ ] Aucun identifiant en clair ne subsiste dans `.gitlab-ci.yml`.
- [ ] Toutes les variables sensibles sont marquées "Protected" et "Masked" dans les paramètres du projet.
- [ ] Aucun script de job n'imprime ou ne redirige une variable sensible vers un fichier d'artefact.
- [ ] Les jobs manipulant des secrets de production s'exécutent sur des runners dédiés et isolés.
- [ ] Les artefacts et caches ne contiennent plus de fichiers avec des secrets.
