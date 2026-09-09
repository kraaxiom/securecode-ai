---
id: github-secrets
category: cicd
cwe: CWE-798
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Fuite de secrets dans GitHub Actions

## Description
Une fuite de secrets GitHub Actions survient quand des identifiants (tokens, clés API, mots de passe) sont codés en dur dans les workflows YAML, affichés dans les logs de build, ou exposés via des `pull_request` provenant de forks non fiables. GitHub masque automatiquement les secrets déclarés via `secrets.*` dans les logs, mais un secret manipulé (encodé, concaténé, imprimé indirectement) échappe à ce masquage. Les workflows déclenchés par `pull_request_target` combinés à un checkout du code du fork sont particulièrement dangereux.

## Où ça apparaît typiquement
- Valeurs sensibles écrites en clair dans un fichier `.github/workflows/*.yml` au lieu d'utiliser `secrets.NOM_SECRET`.
- `echo`, `print` ou `curl -v` d'une variable contenant un secret, affichée dans les logs de la CI.
- Utilisation de `pull_request_target` avec exécution de code provenant d'un fork externe ayant accès aux secrets du dépôt.
- Secrets exposés à des actions tierces non vérifiées (`uses: <org>/<action>@<branch mutable>`).
- Artefacts de build (logs, caches, artifacts téléchargeables) contenant des tokens.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Chaîne ressemblant à un token/clé (motif `ghp_`, `AKIA`, `-----BEGIN PRIVATE KEY-----`, etc.) présente en clair dans un fichier `.yml` sous `.github/workflows/`.
- Utilisation de `pull_request_target` sans restriction claire sur le code exécuté depuis le fork.
- Référence à une action tierce épinglée sur une branche (`@main`, `@master`) plutôt que sur un SHA de commit.
- Absence de scission entre secrets nécessaires au build et secrets de déploiement (permissions `GITHUB_TOKEN` trop larges, `permissions: write-all`).

## Remédiation
- Stocker tout secret dans GitHub Secrets (ou un gestionnaire externe type Vault) et y accéder via `${{ secrets.NOM }}`, jamais en dur.
- Ne jamais afficher un secret dans les logs ; masquer explicitement toute valeur dérivée avec `::add-mask::`.
- Restreindre `permissions:` du `GITHUB_TOKEN` au strict nécessaire (principe du moindre privilège).
- Épingler les actions tierces sur un SHA de commit précis, pas sur une branche mutable.
- Éviter `pull_request_target` avec checkout de code non fiable, ou isoler l'exécution dans un environnement sans accès aux secrets.
- Voir `rules/remediation/github-secrets.md`.

## Exemple avant/après
Voir `examples/yaml/github-secrets/`.

## Références
- OWASP Cheat Sheet: CI/CD Security
- GitHub Docs: Security hardening for GitHub Actions
- CWE-798: Use of Hard-coded Credentials
