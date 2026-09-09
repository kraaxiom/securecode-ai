---
id: ssi
category: injections
cwe: CWE-97
owasp: A03:2021-Injection
severity_default: high
languages: [php]
---

# Server-Side Includes (SSI) Injection

## Description
L'injection SSI survient lorsqu'une entrée utilisateur non neutralisée est intégrée dans une page HTML traitée par un serveur web supportant les Server-Side Includes, permettant à un attaquant d'injecter des directives SSI (`<!--#exec cmd="..." -->`) interprétées côté serveur avant l'envoi de la réponse. Cela peut conduire à une exécution de commande sur le serveur, dans les configurations où l'exécution de commande via SSI est activée.

## Où ça apparaît typiquement
- Serveurs web (Apache avec `mod_include`, IIS) configurés pour interpréter les fichiers `.shtml` ou équivalents.
- Fonctionnalités permettant à l'utilisateur d'influencer le contenu d'une page interprétée en SSI (commentaires, profils, contenu généré affiché dans une page `.shtml`).
- Applications historiques ou legacy encore configurées avec SSI activé par défaut.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Écriture de contenu utilisateur non filtré dans un fichier servi avec l'extension ou la configuration SSI activée.
- Configuration serveur autorisant `Includes` ou `IncludesNOEXEC` sans restriction sur les répertoires utilisateur.
- Absence d'échappement des séquences `<!--#` dans le contenu généré dynamiquement.

## Remédiation
- Désactiver SSI (et en particulier l'exécution de commande via SSI) sur les répertoires où du contenu utilisateur est stocké.
- Échapper ou supprimer les séquences `<!--#` dans tout contenu utilisateur inséré dans une page interprétée.
- Préférer un moteur de templates côté application plutôt que SSI pour du contenu dynamique.
- Voir `rules/remediation/ssi.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/ssi/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Testing Guide: Testing for SSI Injection
- CWE-97: Improper Neutralization of Server-Side Includes (SSI) Within a Web Page
