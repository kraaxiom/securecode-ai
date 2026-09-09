# SSRF classique — Python/Django

CWE-918 — Server-Side Request Forgery (SSRF).

## Description

`vulnerable.py` implémente une fonctionnalité d'aperçu d'URL (`UrlPreviewView`) qui transmet directement l'URL fournie par l'utilisateur à `requests.get()`. Aucune restriction n'est appliquée sur le domaine, le schéma ou l'adresse IP de destination.

## Pourquoi le code vulnérable est dangereux

Un attaquant peut fournir une URL pointant vers des ressources internes normalement inaccessibles depuis l'extérieur : `http://127.0.0.1:8000/admin`, `http://10.0.0.5:6379/` (Redis interne), ou des services d'infrastructure. Le serveur applicatif agit comme relais et retourne le contenu récupéré, permettant un scan du réseau interne et l'exfiltration de données sensibles.

## Explication du correctif

`fixed.py` applique une défense en profondeur :
- **Whitelist de schémas** (`http`/`https`) et **whitelist de domaines** métier explicites — l'utilisateur ne peut jamais cibler un domaine arbitraire.
- **Résolution DNS explicite puis validation de l'IP** obtenue contre les plages privées, loopback, link-local et réservées (`resolve_and_validate`), avant toute requête sortante.
- **DNS pinning** : la requête HTTP est effectuée directement sur l'IP validée (et non sur le nom de domaine), ce qui empêche une seconde résolution DNS différente au moment de la connexion (voir `dns-rebinding`).
- **Redirections désactivées** (`allow_redirects=False`) pour éviter qu'une redirection serveur ne pointe vers une cible interne.
- **Timeout court** pour limiter l'impact d'un scan de ports.

## Notes résiduelles

La whitelist de domaines doit être maintenue et revue à chaque ajout d'intégration. Si le service tourne dans un environnement cloud, une isolation réseau complémentaire (segmentation, groupe de sécurité sans accès aux métadonnées) reste recommandée en défense en profondeur.
