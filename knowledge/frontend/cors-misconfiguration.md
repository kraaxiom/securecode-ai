---
id: cors-misconfiguration
category: frontend
cwe: CWE-942
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: [js, ts, php, python, java]
---

# Mauvaise configuration CORS

## Description
Le Cross-Origin Resource Sharing (CORS) permet à un serveur d'autoriser explicitement certaines origines à effectuer des requêtes cross-domain vers son API. Une mauvaise configuration survient quand le serveur reflète dynamiquement n'importe quelle origine (`Access-Control-Allow-Origin` calculé depuis l'en-tête `Origin` sans liste blanche) tout en autorisant les identifiants (`Access-Control-Allow-Credentials: true`). Cela permet à un site tiers malveillant de lire des réponses contenant des données authentifiées de la victime.

## Où ça apparaît typiquement
- Middleware CORS générique qui reflète l'en-tête `Origin` reçu directement dans la réponse (`Access-Control-Allow-Origin: <origin reçu>`).
- Utilisation du caractère générique `*` combiné à `Access-Control-Allow-Credentials: true` (combinaison invalide mais parfois contournée par du reflet dynamique).
- Listes blanches d'origines mal validées (comparaison par sous-chaîne ou expression régulière trop permissive, ex: `endswith('.example.com')` acceptant `evil-example.com`).
- API internes exposées avec CORS ouvert par facilité de développement, jamais restreintes en production.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Code reflétant `req.headers.origin` directement dans `Access-Control-Allow-Origin` sans vérification contre une liste explicite.
- Présence simultanée de `Access-Control-Allow-Origin: *` (ou reflet dynamique) et `Access-Control-Allow-Credentials: true`.
- Validation d'origine basée sur une regex ou un `startsWith`/`includes` trop permissif plutôt qu'une égalité stricte sur une liste blanche.
- Absence de restriction des méthodes/headers autorisés (`Access-Control-Allow-Methods: *`).

## Remédiation
- Définir une liste blanche stricte et statique des origines autorisées, comparée par égalité exacte.
- Ne jamais combiner reflet dynamique d'origine avec `Access-Control-Allow-Credentials: true`.
- Restreindre les méthodes et en-têtes autorisés au strict nécessaire.
- Séparer les API publiques (CORS ouvert, sans credentials) des API authentifiées (CORS restreint).
- Voir `rules/remediation/cors-misconfiguration.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/cors-misconfiguration/`.

## Références
- OWASP Cheat Sheet: Cross-Origin Resource Sharing (CORS) — via OWASP REST Security Cheat Sheet
- CWE-942: Permissive Cross-domain Policy with Untrusted Domains
