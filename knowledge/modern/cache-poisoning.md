---
id: cache-poisoning
category: modern
cwe: CWE-444
owasp: A04:2021-Insecure Design
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Web Cache Poisoning

## Description
Le Web Cache Poisoning exploite une divergence entre ce que le serveur d'origine utilise pour générer une réponse et ce que le cache (CDN, reverse proxy) utilise comme clé de cache. En envoyant une requête avec un en-tête non normalisé (`X-Forwarded-Host`, `X-Forwarded-Scheme`) qui influence la réponse du serveur sans faire partie de la clé de cache, un attaquant peut faire stocker une réponse malveillante (redirection, contenu injecté) qui sera ensuite servie à tous les visiteurs suivants.

## Où ça apparaît typiquement
- Serveurs d'origine qui reflètent un en-tête non standard (`X-Forwarded-Host`, `X-Original-URL`) dans la réponse (ex: pour construire des liens absolus).
- Configuration de cache/CDN qui met en cache des réponses sur la base de l'URL uniquement, sans inclure les en-têtes qui influencent le contenu.
- Applications générant du contenu dynamique (liens de réinitialisation de mot de passe, ressources) à partir d'en-têtes contrôlables par le client.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Code serveur qui construit des URLs absolues à partir d'en-têtes `X-Forwarded-*` sans validation contre une liste de hôtes autorisés.
- Configuration CDN/cache dont la clé de cache n'inclut pas les en-têtes qui affectent réellement le rendu de la page.
- Absence de directive `Vary` appropriée sur les réponses dont le contenu dépend d'en-têtes spécifiques.

## Remédiation
- Ne jamais faire confiance aux en-têtes `X-Forwarded-*`/`Host` pour générer du contenu sans validation contre une liste blanche.
- Configurer la clé de cache pour inclure tout en-tête qui influence la réponse, ou normaliser/supprimer ces en-têtes avant la génération de la réponse.
- Utiliser des en-têtes `Cache-Control` stricts (`private`, `no-store`) sur les réponses personnalisées ou sensibles.
- Voir `rules/remediation/cache-poisoning.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/cache-poisoning/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP: Web Cache Poisoning
- CWE-444: Inconsistent Interpretation of HTTP Requests ('HTTP Request/Response Smuggling')
