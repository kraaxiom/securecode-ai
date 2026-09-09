---
id: dns-rebinding
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# DNS Rebinding (contournement de filtre SSRF)

## Description
Le DNS rebinding est une technique permettant de contourner les protections anti-SSRF basées uniquement sur une validation de l'adresse IP effectuée avant la requête. L'attaquant contrôle un domaine dont la réponse DNS change entre le moment de la validation (résolution vers une IP publique légitime) et le moment de la requête HTTP réelle (résolution vers une IP interne), exploitant le décalage temporel entre les deux résolutions DNS (TOCTOU).

## Où ça apparaît typiquement
- Toute fonctionnalité de type SSRF qui valide un hôte en résolvant son DNS puis effectue la requête HTTP séparément, en laissant le client HTTP résoudre le DNS une seconde fois.
- Services de vérification d'URL ("webhook validator") qui testent une URL une fois puis la réutilisent plus tard sans revalidation.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Validation de l'adresse IP effectuée via une résolution DNS distincte de celle utilisée par le client HTTP pour la requête effective.
- Absence de verrouillage de l'IP résolue (pinning) pour la durée de la connexion.
- Délai important ou requêtes répétées vers la même URL sans revalidation entre chaque appel.

## Remédiation
- Résoudre le DNS une seule fois, valider l'IP obtenue, puis effectuer la connexion directement sur cette IP validée (DNS pinning) plutôt que de laisser le client HTTP résoudre à nouveau le nom.
- Utiliser un client HTTP ou un proxy sortant qui applique une politique de filtrage réseau au niveau de la connexion TCP elle-même, pas seulement au niveau applicatif.
- Combiner avec l'isolation réseau du service effectuant les requêtes sortantes.
- Voir `rules/remediation/dns-rebinding.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/go/dns-rebinding/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
