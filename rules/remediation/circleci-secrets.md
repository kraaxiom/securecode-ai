# Remédiation — Fuite de secrets dans CircleCI

## Principe
Stocker tout secret dans les variables d'environnement du projet ou un Context restreint, jamais en dur dans `.circleci/config.yml`. Interdire les commandes de dump d'environnement, exiger une approbation manuelle pour les builds provenant de forks, et faire tourner régulièrement les secrets.

## Identifiant en clair dans la configuration
```yaml
# Avant — vulnérable
jobs:
  deploy:
    docker:
      - image: cimg/base:stable
    steps:
      - run: curl -H "Authorization: Bearer sk_live_EXAMPLE_NOT_A_REAL_KEY" https://api.exemple.com/deploy

# Après — sécurisé
jobs:
  deploy:
    docker:
      - image: cimg/base:stable
    steps:
      - run: curl -H "Authorization: Bearer $DEPLOY_TOKEN" https://api.exemple.com/deploy
# DEPLOY_TOKEN défini dans Project Settings > Environment Variables (jamais dans le YAML)
```

## Commande de debug affichant un secret
```yaml
# Avant — vulnérable
- run: env   # dump toutes les variables, y compris les secrets, dans les logs

# Après — sécurisé
- run: echo "Build en cours pour le commit $CIRCLE_SHA1"
# aucune commande n'imprime le contenu des variables sensibles
```

## Context partagé sans restriction et builds de fork
```yaml
# Avant — vulnérable
workflows:
  build-and-deploy:
    jobs:
      - deploy:
          context: prod-secrets   # context accessible à toutes les branches, y compris PR de forks

# Après — sécurisé
workflows:
  build-and-deploy:
    jobs:
      - deploy:
          context: prod-secrets
          filters:
            branches:
              only: main   # limite l'accès au context aux builds de la branche protégée
          requires:
            - approve-deploy   # approbation manuelle obligatoire avant d'exposer le context
      - approve-deploy:
          type: approval
```

## Checklist de vérification post-patch
- [ ] Aucun identifiant en clair ne subsiste dans `.circleci/config.yml`.
- [ ] Toutes les valeurs sensibles proviennent des Environment Variables du projet ou d'un Context restreint.
- [ ] Aucune étape de job n'exécute `env`, `printenv` ou `echo $SECRET` sur une variable sensible.
- [ ] Les Contexts sensibles sont limités aux branches protégées et nécessitent une approbation manuelle.
- [ ] Les secrets exposés par le passé (si suspectés compromis) ont été régénérés/révoqués.
