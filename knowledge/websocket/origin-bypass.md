---
id: origin-bypass
category: websocket
cwe: CWE-346
owasp: A01:2021-Broken Access Control
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Contournement de la vérification d'origine WebSocket

## Description
Une vérification d'origine mal implémentée donne un faux sentiment de sécurité : le serveur croit filtrer les connexions par domaine mais utilise une logique trop permissive (comparaison partielle, expression régulière trop large, whitelist dynamique mal construite). Un attaquant peut alors héberger une page sur un domaine qui satisfait accidentellement la condition de validation et effectuer un Cross-Site WebSocket Hijacking malgré la présence d'un contrôle.

## Où ça apparaît typiquement
- Vérification d'origine par `indexOf`/`includes`/`startsWith` sur une sous-chaîne au lieu d'une comparaison exacte de domaine.
- Regex de validation d'origine mal ancrée (ex: `.*\.example\.com` sans ancrage strict, permettant `example.com.attacker.com` ou `notexample.com`).
- Whitelist construite dynamiquement à partir d'un en-tête contrôlable par le client (ex: `Host` ou `X-Forwarded-Host`).
- Acceptation de `null` comme origine valide (fréquent dans certains environnements sandboxés/iframes).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Comparaison d'origine faite avec une méthode de sous-chaîne plutôt qu'une égalité stricte sur une liste fermée.
- Regex de validation d'origine sans ancres `^`/`$` ou avec un `.` non échappé.
- Origine de référence dérivée d'un en-tête modifiable par le client plutôt que d'une constante serveur.
- Cas particulier `Origin: null` traité comme autorisé.

## Remédiation
- Comparer l'en-tête `Origin` par égalité stricte à une liste blanche codée en dur côté serveur (pas de préfixe/suffixe/regex approximative).
- Ne jamais dériver la liste d'origines autorisées d'un en-tête fourni par le client.
- Rejeter explicitement `Origin: null` sauf besoin métier justifié et documenté.
- Voir `rules/remediation/origin-bypass.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/origin-bypass/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Testing Guide: Testing WebSockets
- CWE-346: Origin Validation Error
