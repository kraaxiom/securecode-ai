---
id: dns-rebinding-modern
category: modern
cwe: CWE-350
owasp: A10:2021-Server-Side Request Forgery
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# DNS Rebinding

## Description
Le DNS rebinding exploite le fait que le navigateur applique la politique de même origine sur la base d'un nom de domaine, alors que la résolution DNS de ce domaine peut changer entre deux requêtes. Un attaquant héberge un domaine dont le TTL DNS est très court : la première résolution pointe vers un serveur légitime pour passer les vérifications, puis une résolution ultérieure pointe vers une adresse interne (localhost, réseau privé), permettant à du code exécuté dans le navigateur de la victime d'atteindre des services internes normalement inaccessibles depuis l'extérieur.

## Où ça apparaît typiquement
- Applications/services locaux (API de développement, interfaces d'administration IoT, agents desktop) qui font confiance à l'origine `Host` sans validation stricte.
- Services internes accessibles sans authentification, protégés uniquement par leur non-exposition réseau supposée.
- Absence de validation d'adresse IP côté serveur pour les requêtes provenant censément d'un domaine de confiance.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Vérification d'accès basée uniquement sur l'en-tête `Host`/`Origin` sans validation de l'adresse IP source réelle.
- Services d'API locaux sans authentification, accessibles via `http://localhost` ou une IP privée depuis n'importe quelle origine.
- Absence de pinning DNS ou de validation que l'IP résolue reste stable durant une session.

## Remédiation
- Exiger une authentification explicite pour tout service accessible localement, indépendamment de l'origine réseau supposée.
- Valider l'en-tête `Host` contre une liste stricte de valeurs attendues côté serveur, en plus de la couche réseau.
- Envisager un pare-feu applicatif qui bloque les résolutions DNS pointant vers des plages d'adresses privées pour les domaines externes.
- Voir `rules/remediation/dns-rebinding-modern.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/dns-rebinding-modern/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP: Server-Side Request Forgery Prevention Cheat Sheet
- CWE-350: Reliance on Reverse DNS Resolution for a Security-Critical Action
