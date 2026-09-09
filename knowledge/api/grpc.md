---
id: grpc
category: api
cwe: CWE-306
owasp: API2:2023-Broken Authentication
severity_default: high
languages: [go, java, python, csharp, js]
---

# Sécurité des interfaces gRPC

## Description
gRPC expose des méthodes de procédure distante (RPC) définies dans un fichier `.proto`, souvent utilisées pour la communication interservices. Un risque fréquent est l'absence de vérification d'authentification ou d'autorisation sur certaines méthodes RPC, en particulier lorsque le service est supposé "interne" mais devient accessible depuis un réseau plus large (mesh, exposition accidentelle). La réflexion gRPC (server reflection), lorsqu'elle est active en production, permet également à un client de découvrir dynamiquement l'ensemble des services et méthodes disponibles, élargissant la surface d'attaque.

## Où ça apparaît typiquement
- Services gRPC internes exposés sans intercepteur d'authentification/autorisation global, en s'appuyant uniquement sur la segmentation réseau.
- Méthodes RPC individuelles omises de la liste des méthodes protégées lors de l'ajout d'un intercepteur (protection partielle du service).
- Activation du service de réflexion gRPC (`grpc.reflection.v1alpha`) en environnement de production.
- Canaux gRPC configurés sans TLS (`insecure` credentials) pour des communications traversant un réseau non totalement isolé.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Définition de service `.proto` sans commentaire ni annotation indiquant le niveau d'autorisation requis par méthode.
- Serveur gRPC instancié sans intercepteur (`UnaryInterceptor`/`StreamInterceptor`) global de vérification de token/session.
- Enregistrement explicite du service de réflexion (`reflection.Register(...)`) sans garde conditionnelle sur l'environnement.
- Utilisation de `grpc.WithInsecure()` ou équivalent dans du code destiné à un déploiement traversant un réseau non fiable.

## Remédiation
- Appliquer un intercepteur d'authentification/autorisation global à toutes les méthodes RPC, avec une liste blanche explicite des méthodes publiques (health check) plutôt qu'une liste noire.
- Désactiver le service de réflexion gRPC en production, ou le restreindre à un réseau d'administration isolé.
- Chiffrer systématiquement les canaux gRPC avec TLS, y compris pour les communications interservices.
- Documenter dans le fichier `.proto` le niveau d'autorisation attendu pour chaque méthode afin de faciliter la revue.
- Voir `rules/remediation/grpc.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/go/grpc/` (et les répertoires équivalents pour java, python, csharp, js).

## Références
- OWASP API Security Top 10 2023: API2-Broken Authentication
- CWE-306: Missing Authentication for Critical Function
