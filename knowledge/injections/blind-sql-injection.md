---
id: blind-sql-injection
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Blind SQL Injection

## Description
L'injection SQL aveugle survient lorsqu'une entrée utilisateur non neutralisée atteint une requête SQL, mais que l'application ne renvoie ni le résultat de la requête ni un message d'erreur SQL exploitable. L'attaquant doit alors déduire des informations en observant des différences de comportement (contenu de la page, code HTTP) ou de temps de réponse. C'est une variante plus difficile à exploiter que l'injection classique, mais tout aussi dangereuse car elle permet une exfiltration progressive des données.

## Où ça apparaît typiquement
- Requêtes de connexion, de recherche ou de filtrage où le résultat n'est pas directement affiché.
- Champs utilisés dans des clauses `WHERE` conditionnelles construites par concaténation.
- API renvoyant uniquement un code de statut ou un message générique ("trouvé"/"non trouvé") selon le résultat de la requête.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de variables utilisateur dans une requête SQL sans paramètre lié, même quand le résultat n'est pas affiché.
- Logique conditionnelle basée sur le nombre de lignes retournées ou l'existence d'un enregistrement, alimentée par une entrée non validée.
- Absence de requêtes préparées dans les modules d'authentification ou de recherche.

## Remédiation
- Utiliser systématiquement des requêtes préparées avec paramètres liés, y compris pour les requêtes conditionnelles.
- Ne jamais construire dynamiquement une clause SQL à partir d'une entrée utilisateur, même en apparence booléenne.
- Uniformiser les temps de réponse et les messages d'erreur pour limiter les canaux d'inférence.
- Voir `rules/remediation/blind-sql-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/blind-sql-injection/`.

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89: Improper Neutralization of Special Elements used in an SQL Command
