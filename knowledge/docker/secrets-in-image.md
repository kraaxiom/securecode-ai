---
id: secrets-in-image
category: docker
cwe: CWE-798
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Secrets intégrés dans une image Docker

## Description
Copier un secret (clé API, mot de passe, certificat privé) dans une couche d'image Docker, même s'il est ensuite supprimé dans une couche ultérieure, le laisse récupérable dans l'historique des couches de l'image. Toute personne ayant accès à l'image (registre, export `docker save`) peut extraire ce secret.

## Où ça apparaît typiquement
- Instruction `COPY`/`ADD` d'un fichier `.env` ou de credentials suivie d'une suppression dans une étape ultérieure du même `Dockerfile`.
- Argument `ARG` ou `ENV` contenant directement un secret, visible via `docker history`.
- Fichiers de configuration contenant des identifiants copiés pour les besoins du build puis oubliés dans l'image finale.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- `ARG`/`ENV` avec un nom évoquant un secret (`API_KEY`, `PASSWORD`, `TOKEN`) assigné à une valeur littérale.
- `COPY`/`ADD` d'un fichier de credentials sans utilisation du mécanisme `--mount=type=secret` du BuildKit.
- Suppression d'un fichier secret dans une couche postérieure à sa copie, au lieu d'un build multi-stage propre.

## Remédiation
- Utiliser les secrets de build BuildKit (`RUN --mount=type=secret`) qui ne persistent pas dans les couches de l'image.
- Injecter les secrets à l'exécution (variables d'environnement du conteneur, gestionnaire de secrets) plutôt qu'au build.
- Utiliser des builds multi-stage pour que les étapes contenant des secrets n'apparaissent jamais dans l'image finale.
- Voir `rules/remediation/secrets-in-image.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/secrets-in-image/` (à créer selon le même schéma).

## Références
- Docker documentation: Build secrets
- OWASP Secrets Management Cheat Sheet
- CWE-798: Use of Hard-coded Credentials
