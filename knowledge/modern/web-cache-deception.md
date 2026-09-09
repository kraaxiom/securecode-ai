---
id: web-cache-deception
category: modern
cwe: CWE-524
owasp: A04:2021-Insecure Design
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# Web Cache Deception

## Description
Le Web Cache Deception survient quand un cache partagé décide de mettre en cache une réponse en se basant sur l'extension apparente de l'URL (ex: `.css`, `.jpg`) plutôt que sur le type de contenu réel ou l'authentification requise. Un attaquant construit une URL comme `/compte/profil/inexistant.css` : le serveur d'application ignore le suffixe et retourne la page de profil personnalisée de la victime, tandis que le cache, trompé par l'extension statique, la stocke et la sert ensuite à n'importe quel visiteur non authentifié.

## Où ça apparaît typiquement
- Routeurs backend qui ignorent les segments d'URL après le chemin logique (path matching permissif) et traitent tout comme la même route.
- Caches/CDN configurés pour mettre en cache par extension de fichier sans vérifier le `Content-Type` réel de la réponse.
- Pages authentifiées contenant des données personnelles sans en-tête `Cache-Control: private, no-store` explicite.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Règles de cache basées sur un motif d'extension d'URL (`*.css`, `*.js`, `*.jpg`) sans vérification du `Content-Type` de la réponse.
- Routes authentifiées sans en-tête `Cache-Control: private, no-store` explicite sur les réponses contenant des données utilisateur.
- Routeur applicatif qui accepte des segments de chemin arbitraires après une route valide sans retourner une 404.

## Remédiation
- Mettre en cache uniquement sur la base du `Content-Type` réel de la réponse, jamais sur l'apparence de l'URL.
- Ajouter `Cache-Control: private, no-store` explicitement sur toute réponse contenant des données utilisateur/authentifiées.
- Configurer le routeur pour retourner une 404 stricte sur les chemins non reconnus plutôt que de retomber sur une route générique.
- Voir `rules/remediation/web-cache-deception.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/web-cache-deception/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP: Web Cache Deception
- CWE-524: Use of Cache Containing Sensitive Information
