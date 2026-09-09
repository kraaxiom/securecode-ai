---
id: id-rsa
category: sensitive-files
cwe: CWE-522
owasp: A05:2021-Security Misconfiguration
severity_default: critical
languages: []
---

# Exposition d'une clé privée SSH (id_rsa)

## Description
Une clé privée SSH (`id_rsa`, `id_ed25519`, ou toute clé sans passphrase) exposée sur un service web accessible permet à un attaquant de s'authentifier directement sur tout système où la clé publique correspondante est autorisée, sans avoir besoin de mot de passe. C'est l'une des expositions les plus critiques car elle donne un accès shell direct, souvent avec des privilèges étendus si la clé appartient à un compte de déploiement.

## Où ça apparaît typiquement
- Répertoire `.ssh/` copié par erreur dans une image Docker ou un artefact de déploiement.
- Scripts de provisioning qui embarquent une clé privée dans le dépôt pour automatiser des connexions.
- Sauvegardes de configuration système incluant le répertoire home d'un utilisateur exposées via un service web.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichiers nommés `id_rsa`, `id_ed25519`, `*.pem` sans extension `.pub` présents dans un artefact de déploiement ou une image conteneur.
- Répertoire `.ssh` copié dans une image Docker (`COPY .ssh` ou équivalent) au lieu d'être monté en secret de build.
- Absence de `.gitignore`/`.dockerignore` excluant les clés privées.

## Remédiation
- Ne jamais inclure de clé privée dans un dépôt, une image conteneur ou un artefact déployé ; utiliser des secrets de build (BuildKit `--mount=type=secret`) ou un agent SSH.
- Bloquer l'accès à toute clé privée détectée dans un webroot et la révoquer immédiatement (retirer la clé publique correspondante des `authorized_keys`).
- Utiliser des clés avec passphrase et une rotation régulière pour les comptes de déploiement.
- Voir `rules/remediation/id-rsa.md` pour les diffs par configuration CI/Docker.

## Exemple avant/après
Voir `examples/config/id-rsa/` (Dockerfile avant/après utilisation de secrets de build).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-522: Insufficiently Protected Credentials
