---
id: browser-cache-poisoning
category: modern
cwe: CWE-525
owasp: A04:2021-Insecure Design
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Browser Cache Poisoning

## Description
Contrairement au cache partagé (CDN), le browser cache poisoning cible le cache local du navigateur de la victime : une application qui renvoie des en-têtes `Cache-Control` trop permissifs sur des réponses contenant des données sensibles ou spécifiques à l'utilisateur permet que ces réponses soient réutilisées de façon incorrecte, notamment sur des postes partagés, ou combinées à d'autres failles (XSS, redirection ouverte) pour persister une charge malveillante dans le cache local au-delà de la session.

## Où ça apparaît typiquement
- Réponses HTTP contenant des données utilisateur (profil, solde, panier) sans en-tête `Cache-Control: no-store` explicite.
- Pages d'authentification ou de déconnexion mises en cache par le navigateur, restant accessibles via le bouton "précédent" après déconnexion.
- Ressources générées dynamiquement (JS/CSS contenant des tokens) servies avec des en-têtes de cache longue durée.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de `Cache-Control: no-store` sur les routes affichant des données personnelles ou de session.
- En-têtes `Cache-Control: public, max-age=...` appliqués par défaut à toutes les réponses sans distinction de sensibilité.
- Pages de logout/déconnexion sans en-tête empêchant leur mise en cache.

## Remédiation
- Définir explicitement `Cache-Control: no-store` sur toute réponse contenant des données sensibles ou spécifiques à l'utilisateur.
- Différencier la politique de cache entre ressources statiques génériques et contenu dynamique personnalisé.
- Invalider/exclure du cache les pages liées à l'authentification (login, logout, changement de mot de passe).
- Voir `rules/remediation/browser-cache-poisoning.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/browser-cache-poisoning/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Testing Guide: Testing for Browser Cache Weaknesses
- CWE-525: Use of Web Browser Cache Containing Sensitive Information
