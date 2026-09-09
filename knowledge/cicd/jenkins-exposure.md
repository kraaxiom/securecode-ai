---
id: jenkins-exposure
category: cicd
cwe: CWE-16
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Exposition de Jenkins mal configuré

## Description
Une instance Jenkins exposée survient quand le serveur est accessible depuis Internet sans authentification stricte, avec des scripts en clair contenant des identifiants, ou avec la console Script (Groovy) accessible à des utilisateurs non administrateurs. Un Jenkins mal durci permet l'exécution de code arbitraire côté serveur et l'accès aux credentials stockés (SSH, cloud, registre Docker) utilisés par les pipelines.

## Où ça apparaît typiquement
- Interface Jenkins accessible publiquement sans authentification ("Anyone can do anything") ou avec des comptes par défaut.
- Console de script Groovy (`/script`) accessible à des utilisateurs non-admin.
- Credentials stockés dans Jenkins Credentials Manager avec une portée trop large (globale au lieu de scoping par job/dossier).
- Logs de build affichant des variables d'environnement sensibles.
- Plugins obsolètes avec des vulnérabilités connues non corrigées.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Endpoint Jenkins répondant sans redirection vers une page de connexion.
- Configuration de sécurité globale ("Configure Global Security") avec autorisation de type "Anyone can do anything" ou matrice trop permissive.
- Références à des credentials Jenkins codés en dur dans un `Jenkinsfile` plutôt qu'utilisés via `withCredentials`.
- Version de Jenkins ou de plugins clairement obsolète dans les métadonnées exposées.

## Remédiation
- Activer l'authentification et une matrice d'autorisation stricte (principe du moindre privilège par rôle/projet).
- Restreindre l'accès réseau à Jenkins (VPN, allowlist IP) et désactiver l'accès anonyme.
- Utiliser `withCredentials` avec des credentials scopés par dossier/job, jamais codés en dur.
- Maintenir Jenkins et ses plugins à jour, désactiver la console Script pour les non-admins.
- Voir `rules/remediation/jenkins-exposure.md`.

## Exemple avant/après
Voir `examples/groovy/jenkins-exposure/`.

## Références
- OWASP Cheat Sheet: CI/CD Security
- CIS Jenkins Benchmark
- CWE-16: Configuration
