# SSRF classique — exemple Go

## Description de la vulnérabilité
Le fichier `vulnerable.go` implémente un service d'aperçu d'URL (« URL
preview ») qui effectue une requête HTTP sortante vers une adresse
entièrement contrôlée par l'utilisateur (`GET /preview?url=...`), sans
aucune validation de la destination. Le contenu récupéré est ensuite
renvoyé tel quel au client, exposant potentiellement des ressources
internes (interfaces d'administration, API internes, fichiers accessibles
en HTTP) qui ne devraient pas être atteignables depuis l'extérieur.

## Référence CWE
- **CWE-918** : Server-Side Request Forgery (SSRF)
- OWASP Top 10 2021 : A10 — Server-Side Request Forgery

## Scénario d'exploitation (résumé conceptuel)
Un attaquant fournit, comme valeur du paramètre `url`, l'adresse d'une
ressource interne au lieu d'une URL publique légitime (par exemple un
service d'administration écoutant sur le réseau interne, ou une base de
données exposant une API HTTP). Le serveur applicatif, agissant comme
proxy involontaire, effectue la requête depuis sa propre position réseau
— qui a généralement accès à des segments internes inaccessibles
directement depuis Internet — puis retourne le contenu obtenu dans sa
réponse. L'attaquant peut ainsi lire des données normalement protégées par
la segmentation réseau, sans jamais s'y connecter directement. Le
comportement de suivi automatique des redirections HTTP peut en outre
permettre de contourner un filtrage partiel appliqué uniquement sur l'URL
initiale.

## Explication du correctif
Le fichier `fixed.go` applique plusieurs contrôles complémentaires :

1. **Whitelist stricte de domaines** (`domainAllowlist`) : seuls des hôtes
   explicitement autorisés peuvent être contactés — l'utilisateur ne peut
   jamais choisir librement la destination.
2. **Résolution DNS explicite et unique**, suivie d'une validation de
   l'adresse IP obtenue contre les plages privées, loopback et link-local
   (`isBlockedIP`), avant toute connexion.
3. **DNS pinning** : la connexion TCP est établie directement sur l'IP déjà
   validée (`safeDialContext`), pas sur un nouveau nom résolu séparément
   par le client HTTP — ce qui neutralise également les attaques de type
   DNS rebinding (voir `examples/go/dns-rebinding/`).
4. **Désactivation des redirections automatiques** (`CheckRedirect`) pour
   empêcher qu'une URL publique validée redirige ensuite vers une
   ressource interne.
5. **Messages d'erreur génériques** côté client pour éviter de fournir un
   oracle de scan réseau interne à un attaquant.

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
- `knowledge/ssrf/ssrf-classique.md`
