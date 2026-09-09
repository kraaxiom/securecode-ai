# Clés/secrets codés en dur

## CWE
CWE-798 — Use of Hard-coded Credentials

## Description
Des clés API, mots de passe ou secrets cryptographiques écrits directement
dans le code source sont exposés à quiconque accède au dépôt (y compris via
l'historique git, un build public ou une fuite de dépôt).

## Pourquoi c'est vulnérable
Dans `vulnerable.go`, `stripeSecretKey`, `awsAccessKey`, `awsSecretKey` et
`dbPassword` sont des constantes littérales. Ces secrets :
- sont visibles par toute personne ayant accès au code ou à son historique ;
- ne peuvent pas être tournés (rotation) sans redéployer une nouvelle version ;
- finissent souvent dans des logs de build ou des images de conteneur.

## Correction
`fixed.go` charge chaque secret via `os.Getenv`, avec échec explicite si la
variable est absente. En production, ces variables proviennent idéalement
d'un gestionnaire de secrets (Vault, AWS Secrets Manager, GCP Secret Manager).
Tout secret déjà exposé doit être révoqué et régénéré côté fournisseur.

## Références
- rules/remediation/hardcoded-keys.md
- knowledge/crypto/hardcoded-keys.md
- CWE-798: https://cwe.mitre.org/data/definitions/798.html
