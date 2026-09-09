# Remédiation — Fuite de secrets dans GitHub Actions

## Principe
Toute valeur sensible doit provenir de GitHub Secrets (`${{ secrets.NOM }}`), jamais être codée en dur. Ne jamais afficher un secret dans les logs, restreindre `permissions:` du `GITHUB_TOKEN`, épingler les actions tierces sur un SHA, et éviter `pull_request_target` avec exécution de code non fiable.

## Identifiant en clair dans le workflow
```yaml
# Avant — vulnérable
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - run: curl -H "Authorization: Bearer ghp_1234567890abcdef" https://api.exemple.com/deploy

# Après — sécurisé
jobs:
  deploy:
    runs-on: ubuntu-latest
    permissions:
      contents: read
    steps:
      - run: curl -H "Authorization: Bearer ${{ secrets.DEPLOY_TOKEN }}" https://api.exemple.com/deploy
```

## Impression indirecte d'un secret
```yaml
# Avant — vulnérable
- name: Debug
  run: echo "Token=${{ secrets.API_TOKEN }}"

# Après — sécurisé
- name: Debug
  run: |
    echo "::add-mask::${{ secrets.API_TOKEN }}"
    echo "Déploiement démarré (secret masqué)"
```

## `pull_request_target` dangereux + action non épinglée + permissions larges
```yaml
# Avant — vulnérable
on: pull_request_target
permissions: write-all
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@main   # branche mutable
        with:
          ref: ${{ github.event.pull_request.head.sha }}  # checkout du code du fork avec accès aux secrets
      - uses: some-org/some-action@main

# Après — sécurisé
on: pull_request   # pas d'accès aux secrets du dépôt de base pour les forks
permissions:
  contents: read
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@8ade135a41bc03ea155e62e844d188df1ea18608  # SHA épinglé
      - uses: some-org/some-action@a1b2c3d4e5f6...  # SHA épinglé
```

## Checklist de vérification post-patch
- [ ] Aucun motif de token/clé en clair ne subsiste dans les fichiers `.github/workflows/*.yml`.
- [ ] Tous les secrets sont accédés via `${{ secrets.NOM }}`, jamais codés en dur.
- [ ] Aucune étape n'imprime la valeur brute d'un secret dans les logs.
- [ ] `permissions:` du `GITHUB_TOKEN` est restreint au strict nécessaire (pas de `write-all` par défaut).
- [ ] Les actions tierces sont épinglées sur un SHA de commit, pas sur une branche mutable.
- [ ] `pull_request_target` n'est plus utilisé avec un checkout de code de fork ayant accès aux secrets.
