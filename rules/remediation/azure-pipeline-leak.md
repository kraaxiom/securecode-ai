# Remédiation — Fuite de secrets dans Azure Pipelines

## Principe
Ne jamais coder de secret en dur dans `azure-pipelines.yml`. Stocker les identifiants dans un Variable Group lié à Azure Key Vault, marquer chaque variable sensible comme secrète, restreindre les service connections aux pipelines explicites, et ne jamais imprimer une variable secrète dans les logs.

## Variable en clair dans le YAML
```yaml
# Avant — vulnérable
steps:
  - script: |
      curl -u admin:P@ssw0rd123 https://api.exemple.com/deploy
    displayName: 'Deploy'

# Après — sécurisé
variables:
  - group: prod-secrets   # Variable Group lié à Azure Key Vault

steps:
  - script: |
      curl -u admin:$(DEPLOY_PASSWORD) https://api.exemple.com/deploy
    displayName: 'Deploy'
    env:
      DEPLOY_PASSWORD: $(DEPLOY_PASSWORD)
```

## Impression d'un secret dans les logs
```yaml
# Avant — vulnérable
- script: echo "Token utilisé $(API_TOKEN)"

# Après — sécurisé
- script: echo "Déploiement en cours (token masqué automatiquement par Azure DevOps)"
  env:
    API_TOKEN: $(API_TOKEN)   # variable marquée "secret" -> masquage auto dans les logs
```

## Service connection ouverte à tous les pipelines
```yaml
# Avant — vulnérable
# Service connection "prod-azure" configurée avec
# "Grant access permission to all pipelines" = activé

# Après — sécurisé
# Service connection "prod-azure" configurée avec accès restreint,
# approuvée explicitement uniquement pour le pipeline "release-prod.yml"
resources:
  repositories:
    - repository: self
trigger:
  branches:
    include:
      - main
pr: none   # empêche le déclenchement automatique sur PR de fork ayant accès aux mêmes variables
```

## Checklist de vérification post-patch
- [ ] Aucun identifiant en clair (mot de passe, token, clé) ne subsiste dans le YAML du pipeline corrigé.
- [ ] Toutes les variables sensibles proviennent d'un Variable Group lié à Azure Key Vault et sont marquées "secret".
- [ ] Aucune étape de script n'affiche explicitement la valeur d'une variable secrète.
- [ ] Les service connections sensibles ne sont plus accessibles à "tous les pipelines" — accès restreint et approuvé explicitement.
- [ ] Les déclenchements automatiques sur pull request externe (fork) n'exposent plus de secrets de production.
