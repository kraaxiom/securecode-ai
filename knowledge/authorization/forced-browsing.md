---
id: forced-browsing
category: authorization
cwe: CWE-425
owasp: A01:2021-Broken Access Control
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Forced Browsing

## Description
Le forced browsing consiste, pour un attaquant, à accéder directement à des URL ou endpoints qui ne sont pas liés depuis l'interface utilisateur normale mais qui restent accessibles côté serveur sans contrôle d'autorisation propre. L'application repose alors à tort sur le fait que ces pages/fonctionnalités ne sont pas découvrables (sécurité par obscurité) plutôt que sur une vérification d'accès effective.

## Où ça apparaît typiquement
- Pages d'administration ou de configuration accessibles via une URL non référencée mais non protégée.
- Fichiers de sauvegarde, journaux ou exports laissés accessibles dans un répertoire public sans lien direct.
- Étapes intermédiaires d'un flux multi-étapes (paiement, inscription) accessibles directement en sautant les étapes précédentes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Route ou fichier statique sensible accessible sans middleware d'authentification/autorisation, uniquement protégé par l'absence de lien visible.
- Flux multi-étapes où chaque étape ne vérifie pas que les étapes précédentes ont été correctement complétées côté serveur.
- Fichiers de configuration, sauvegardes ou logs déployés dans un répertoire servi statiquement.

## Remédiation
- Appliquer un contrôle d'authentification/autorisation explicite sur toute route ou ressource sensible, indépendamment de sa visibilité dans l'interface.
- Ne jamais stocker de fichiers sensibles dans un répertoire servi publiquement.
- Vérifier côté serveur l'état d'avancement réel d'un flux multi-étapes avant d'autoriser l'accès à une étape donnée.
- Voir `rules/remediation/forced-browsing.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/forced-browsing/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Authorization
- CWE-425: Direct Request ('Forced Browsing')
