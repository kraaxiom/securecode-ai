# MD5 — Hachage faible

## CWE
CWE-328 — Use of Weak Hash

## Description
MD5 est vulnérable aux collisions et surtout beaucoup trop rapide pour être
utilisé comme fonction de hachage de mots de passe : les attaquants peuvent
tester des milliards de candidats par seconde sur du matériel grand public.

## Pourquoi c'est vulnérable
`vulnerable.go` hache directement le mot de passe avec `crypto/md5`, sans sel
ni facteur de coût. Un vidage de base de données permet un cassage massif via
rainbow tables ou force brute GPU.

## Correction
`fixed.go` utilise **Argon2id** (`golang.org/x/crypto/argon2`), avec sel
aléatoire par mot de passe et paramètres de coût mémoire/temps réglables,
rendant les attaques hors-ligne bien plus lentes et coûteuses.

## Références
- rules/remediation/md5.md
- knowledge/crypto/md5.md
- CWE-328: https://cwe.mitre.org/data/definitions/328.html
