---
id: crlf-injection
category: injections
cwe: CWE-93
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# CRLF Injection

## Description
L'injection CRLF consiste à insérer des séquences de retour chariot et saut de ligne (`\r\n`) dans une entrée utilisateur qui est ensuite écrite telle quelle dans un flux structuré par lignes (en-têtes HTTP, fichiers de log, commandes de protocole texte). Cela permet à l'attaquant d'ajouter de nouvelles lignes non prévues par le développeur, ouvrant la voie à la falsification d'en-têtes, l'empoisonnement de logs ou, dans le cas HTTP, au découpage de réponse (response splitting).

## Où ça apparaît typiquement
- Construction d'en-têtes HTTP (`Location`, `Set-Cookie`) à partir d'une valeur utilisateur non filtrée.
- Écriture de données utilisateur dans des fichiers de logs sans échappement des retours à la ligne.
- Construction de commandes de protocoles textuels (SMTP, IMAP) où chaque ligne a une signification.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Affectation directe d'une valeur utilisateur à un en-tête de réponse sans validation de l'absence de `\r`/`\n`.
- Écriture de journaux applicatifs par concaténation de chaîne incluant des champs utilisateur.
- Absence de fonction d'échappement dédiée avant écriture dans un flux structuré par ligne.

## Remédiation
- Utiliser les API de framework qui gèrent les en-têtes de façon sûre (rejettent ou encodent les caractères CR/LF) plutôt que de les concaténer manuellement.
- Filtrer ou rejeter toute entrée contenant `\r` ou `\n` avant utilisation dans un en-tête, un log ou une commande protocolaire.
- Utiliser des bibliothèques de logging structuré (JSON) qui échappent automatiquement les caractères de contrôle.
- Voir `rules/remediation/crlf-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/crlf-injection/`.

## Références
- OWASP Cheat Sheet: HTTP Headers Cheat Sheet
- CWE-93: Improper Neutralization of CRLF Sequences ('CRLF Injection')
