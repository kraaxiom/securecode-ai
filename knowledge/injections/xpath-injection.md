---
id: xpath-injection
category: injections
cwe: CWE-643
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp]
---

# XPath Injection

## Description
L'injection XPath survient lorsqu'une entrée utilisateur non neutralisée est intégrée directement dans une expression XPath utilisée pour interroger un document XML, permettant à un attaquant de modifier la logique de sélection de nœuds. Selon le contexte (ex: authentification basée sur un document XML d'utilisateurs), cela peut permettre un contournement d'authentification ou l'extraction de données du document XML interrogé.

## Où ça apparaît typiquement
- Requêtes d'authentification ou de recherche construites dynamiquement contre un document XML (`//user[username='...' and password='...']`).
- API exposant une recherche dans un stockage basé sur XML avec le terme de recherche inséré tel quel dans l'expression.
- Systèmes legacy utilisant XML comme mini base de données interrogée via XPath.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Concaténation de chaînes pour construire une expression XPath à partir d'une entrée utilisateur.
- Absence de fonction d'échappement dédiée pour les valeurs insérées dans une expression XPath.
- Utilisation d'XPath pour des vérifications d'authentification sans passage par des requêtes paramétrées (`XPathExpression` avec variables liées).

## Remédiation
- Utiliser des expressions XPath paramétrées (variables liées) plutôt que la concaténation de chaînes, lorsque l'API du langage le permet.
- Échapper les caractères spéciaux XPath dans toute valeur utilisateur insérée dans une expression.
- Éviter d'utiliser XPath comme mécanisme d'authentification; préférer une base de données avec hachage de mot de passe.
- Voir `rules/remediation/xpath-injection.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/xpath-injection/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: XPath Injection Prevention (voir aussi Testing Guide)
- CWE-643: Improper Neutralization of Data within XPath Expressions ('XPath Injection')
