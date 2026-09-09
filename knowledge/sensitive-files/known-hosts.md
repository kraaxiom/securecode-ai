---
id: known-hosts
category: sensitive-files
cwe: CWE-200
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: []
---

# Exposition du fichier known_hosts

## Description
Le fichier `known_hosts` liste les empreintes des serveurs SSH auxquels un système s'est déjà connecté. Bien qu'il ne contienne pas de secret exploitable directement, sa divulgation révèle l'infrastructure interne d'une organisation (noms d'hôtes, adresses IP internes, topologie réseau), fournissant à un attaquant une cartographie utile pour préparer des attaques ciblées ou du pivoting.

## Où ça apparaît typiquement
- Répertoire `.ssh/` copié entièrement dans une image conteneur ou un artefact de déploiement.
- Sauvegardes de répertoire home exposées via un service de fichiers statiques mal configuré.
- Outils d'automatisation qui embarquent le fichier dans un artefact de build pour préserver l'historique de connexion.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Fichier `known_hosts` présent dans un artefact de déploiement ou une image conteneur accessible.
- Répertoire `.ssh` entier copié sans filtrage dans un contexte de build ou de packaging.
- Absence d'exclusion explicite de `.ssh/known_hosts` dans `.dockerignore`/`.gitignore`.

## Remédiation
- Exclure le répertoire `.ssh` complet des artefacts de build et de déploiement.
- Ne conserver que les données strictement nécessaires à l'exécution dans les images/artefacts publiés.
- Considérer toute information de topologie interne exposée comme un signal à investiguer pour une reconnaissance active en cours.
- Voir `rules/remediation/known-hosts.md` pour les diffs par configuration CI/Docker.

## Exemple avant/après
Voir `examples/config/known-hosts/` (Dockerfile avant/après exclusion du répertoire .ssh).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-200: Exposure of Sensitive Information to an Unauthorized Actor
