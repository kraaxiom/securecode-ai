---
id: missing-rate-limit
category: api
cwe: CWE-770
owasp: API4:2023-Unrestricted Resource Consumption
severity_default: medium
languages: [php, js, python, java, csharp, go]
---

# Absence de limitation de débit (rate limiting)

## Description
L'absence de limitation de débit se produit lorsqu'une API n'impose aucune restriction sur le nombre de requêtes qu'un client peut émettre dans un intervalle de temps donné. Cela expose l'application à des abus tels que le bourrage d'identifiants (credential stuffing), l'énumération de comptes, l'épuisement de ressources serveur (CPU, mémoire, connexions base de données) ou des coûts d'infrastructure/API tierces incontrôlés. Ce risque concerne en priorité les endpoints d'authentification, de réinitialisation de mot de passe, et toute opération coûteuse en ressources.

## Où ça apparaît typiquement
- Endpoints de connexion (`/login`), d'inscription, ou de réinitialisation de mot de passe sans compteur de tentatives.
- API publiques ou partenaires sans quota par clé API ou par adresse IP.
- Opérations coûteuses (recherche, export, génération de rapport, appel à un service tiers facturé) accessibles sans limite de fréquence.
- Endpoints exposant l'infrastructure directement sans passerelle API (API gateway) centralisant le throttling.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Route sensible (authentification, réinitialisation) définie sans middleware ou décorateur de limitation de débit.
- Absence de dépendance ou de configuration de bibliothèque de rate limiting dans le projet (aucun composant type token bucket/sliding window présent).
- API gateway ou reverse proxy configuré sans règle de throttling pour les routes exposées publiquement.
- Absence de mécanisme de verrouillage progressif de compte après échecs d'authentification répétés.

## Remédiation
- Mettre en place une limitation de débit par IP, par utilisateur et/ou par clé API sur tous les endpoints sensibles ou coûteux, avec des seuils adaptés au contexte métier.
- Ajouter un verrouillage progressif ou un délai croissant après des échecs d'authentification répétés.
- Centraliser le throttling au niveau de la passerelle API lorsque c'est possible, en complément de la protection applicative.
- Retourner des codes de réponse standards (429 Too Many Requests) avec en-têtes indiquant les quotas restants.
- Voir `rules/remediation/missing-rate-limit.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/missing-rate-limit/` (et les répertoires équivalents pour js, python, java, csharp, go).

## Références
- OWASP API Security Top 10 2023: API4-Unrestricted Resource Consumption
- CWE-770: Allocation of Resources Without Limits or Throttling
