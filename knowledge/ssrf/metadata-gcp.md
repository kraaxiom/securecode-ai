---
id: metadata-gcp
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: critical
languages: [php, python, js, java, go, csharp, rust]
---

# SSRF vers le service de métadonnées GCP

## Description
Le service de métadonnées Google Cloud (accessible via `metadata.google.internal` ou `169.254.169.254`) fournit aux instances Compute Engine et à d'autres ressources GCP des informations de configuration, y compris des jetons d'accès pour le compte de service attaché. Une SSRF atteignant ce service permet à un attaquant de récupérer ces jetons et d'agir avec les permissions IAM du compte de service de la ressource compromise.

## Où ça apparaît typiquement
- Toute fonctionnalité vulnérable à une SSRF classique, hébergée sur GCE, Cloud Run ou GKE avec un compte de service attaché.
- Proxys ou fetchers d'URL qui n'excluent pas la plage link-local et le nom d'hôte de métadonnées.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de filtrage explicite de l'adresse `169.254.169.254` et du nom `metadata.google.internal` dans la logique de validation des destinations de requêtes sortantes.
- Compte de service attaché à la ressource avec des rôles IAM larges alors que le service expose une SSRF potentielle.
- Absence de vérification que les requêtes sortantes contrôlées par l'utilisateur ne peuvent pas injecter l'en-tête requis (`Metadata-Flavor: Google`) attendu par l'API de métadonnées.

## Remédiation
- Filtrer explicitement les plages d'adresses privées et link-local ainsi que le nom d'hôte de métadonnées dans toute logique applicative de requêtes sortantes.
- Appliquer le principe du moindre privilège sur les rôles IAM attribués au compte de service de chaque ressource.
- Ne jamais laisser une requête sortante contrôlée par l'utilisateur définir librement les en-têtes HTTP transmis.
- Voir `rules/remediation/metadata-gcp.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-django/metadata-gcp/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
