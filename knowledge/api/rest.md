---
id: rest
category: api
cwe: CWE-284
owasp: A01:2021-Broken Access Control
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Mauvaise configuration de sécurité REST

## Description
Les API REST souffrent souvent de défauts de configuration transverses qui ne relèvent pas d'une vulnérabilité applicative unique mais d'une hygiène insuffisante : absence de vérification du verbe HTTP réellement utilisé (autorisation d'une méthode `DELETE` sur une route pensée pour `GET` uniquement), gestion d'erreurs trop verbeuse exposant des détails internes, ou absence de contrôle d'accès cohérent entre les différentes méthodes d'une même ressource. Ces failles élargissent la surface d'attaque même quand chaque contrôle pris isolément semble correct.

## Où ça apparaît typiquement
- Frameworks routant automatiquement toutes les méthodes HTTP vers un même contrôleur sans restriction explicite par verbe.
- Contrôleurs vérifiant l'autorisation pour `GET` mais oubliant de la répliquer pour `PUT`/`PATCH`/`DELETE` sur la même ressource.
- Gestionnaires d'erreur par défaut renvoyant la stack trace complète ou des messages SQL bruts au client en production.
- Absence d'en-têtes de sécurité HTTP (`Content-Type` strict, `X-Content-Type-Options`, CORS mal configuré en `*` avec credentials).
- Documentation d'API (Swagger/OpenAPI) exposée publiquement en production sans restriction d'accès.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Définition de route acceptant plusieurs verbes HTTP (`ANY`, `ALL`) sans distinction de contrôle d'accès par méthode.
- Contrôle d'autorisation présent sur une méthode d'un contrôleur REST mais absent sur une autre méthode de la même ressource.
- Configuration `debug`/`display_errors` activée en environnement de production, ou gestionnaire d'exception par défaut du framework non personnalisé.
- Configuration CORS avec origine `*` combinée à `Access-Control-Allow-Credentials: true`.
- Fichier de spécification OpenAPI/Swagger accessible sans authentification en production.

## Remédiation
- Restreindre explicitement chaque route à son verbe HTTP prévu et appliquer le contrôle d'autorisation de façon homogène sur toutes les méthodes d'une même ressource.
- Désactiver les messages d'erreur détaillés en production et centraliser la gestion d'erreurs pour ne renvoyer que des messages génériques au client.
- Configurer CORS de façon restrictive (liste blanche d'origines explicites, pas de wildcard avec credentials).
- Protéger ou désactiver l'accès public à la documentation d'API en production.
- Voir `rules/remediation/rest.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/rest/` (et les répertoires équivalents pour js, python, java, csharp, go).

## Références
- OWASP Top 10 2021: A01-Broken Access Control
- OWASP API Security Top 10 2023: API8-Security Misconfiguration
- CWE-284: Improper Access Control
