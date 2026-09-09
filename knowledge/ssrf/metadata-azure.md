---
id: metadata-azure
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: critical
languages: [php, python, js, java, go, csharp, rust]
---

# SSRF vers le service de métadonnées Azure (IMDS)

## Description
Le service Azure Instance Metadata Service (IMDS), accessible à l'adresse link-local `169.254.169.254`, expose des informations de configuration de la machine virtuelle ainsi que, via son endpoint dédié, des jetons d'accès pour les identités managées (Managed Identity) attachées. Une SSRF atteignant ce service permet potentiellement à un attaquant de récupérer un jeton Azure AD valide et d'agir avec les permissions de l'identité managée de la ressource.

## Où ça apparaît typiquement
- Toute fonctionnalité vulnérable à une SSRF classique, hébergée sur une VM Azure, un App Service ou une Function avec identité managée activée.
- Proxys ou fetchers d'URL qui n'excluent pas la plage link-local et ne vérifient pas les en-têtes requis par Azure.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de filtrage explicite de l'adresse `169.254.169.254` dans la logique de validation des destinations de requêtes sortantes.
- Identité managée attachée à la ressource avec des rôles Azure AD larges alors que le service expose une SSRF potentielle.
- Absence de vérification que les requêtes sortantes contrôlées par l'utilisateur ne peuvent pas injecter l'en-tête requis (`Metadata: true`) attendu par l'IMDS.

## Remédiation
- Filtrer explicitement les plages d'adresses privées et link-local (dont `169.254.169.254`) dans toute logique applicative de requêtes sortantes.
- Appliquer le principe du moindre privilège sur les rôles attribués à l'identité managée de chaque ressource.
- Ne jamais laisser une requête sortante contrôlée par l'utilisateur définir librement les en-têtes HTTP transmis.
- Voir `rules/remediation/metadata-azure.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/csharp-dotnet/metadata-azure/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
