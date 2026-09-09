---
id: metadata-aws
category: ssrf
cwe: CWE-918
owasp: A10:2021-Server-Side Request Forgery (SSRF)
severity_default: critical
languages: [php, python, js, java, go, csharp, rust]
---

# SSRF vers le service de métadonnées AWS (IMDS)

## Description
Le service de métadonnées d'instance AWS (IMDS, accessible à l'adresse link-local `169.254.169.254`) fournit aux instances EC2 des informations de configuration, y compris potentiellement des identifiants temporaires du rôle IAM attaché. Une SSRF permettant d'atteindre cette adresse depuis le serveur applicatif peut permettre à un attaquant de récupérer ces identifiants et d'usurper les permissions IAM de l'instance, souvent bien au-delà du périmètre applicatif initial.

## Où ça apparaît typiquement
- Toute fonctionnalité vulnérable à une SSRF classique, hébergée sur une infrastructure EC2/ECS/Lambda.
- Proxys ou fetchers d'URL qui n'excluent pas explicitement la plage d'adresses link-local (`169.254.0.0/16`).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de filtrage explicite de l'adresse `169.254.169.254` (et plus largement des plages link-local/privées) dans la logique de validation des destinations de requêtes sortantes.
- Utilisation d'IMDSv1 (sans jeton de session) au niveau de l'infrastructure, qui ne nécessite qu'une simple requête GET pour être interrogé.
- Rôle IAM attaché à l'instance/tâche avec des permissions larges alors que le service ne fait qu'exposer une SSRF potentielle.

## Remédiation
- Migrer vers IMDSv2 (jeton de session requis via une requête PUT), ce qui neutralise la plupart des SSRF classiques basées sur de simples requêtes GET.
- Filtrer explicitement les plages d'adresses privées et link-local (dont `169.254.169.254`) dans toute logique applicative de requêtes sortantes.
- Appliquer le principe du moindre privilège sur le rôle IAM attaché à chaque instance/service.
- Voir `rules/remediation/metadata-aws.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/go/metadata-aws/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
