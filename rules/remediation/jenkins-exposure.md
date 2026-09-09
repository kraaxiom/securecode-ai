# Remédiation — Exposition de Jenkins mal configuré

## Principe
Activer l'authentification et une matrice d'autorisation stricte, restreindre l'accès réseau à Jenkins, utiliser `withCredentials` avec des credentials scopés par dossier/job au lieu de valeurs codées en dur, désactiver la console Script pour les non-admins, et maintenir Jenkins/plugins à jour.

## Autorisation globale trop permissive
```groovy
// Avant — vulnérable (Configure Global Security)
// Authorization: "Anyone can do anything"
// Aucune authentification requise pour accéder à Jenkins

// Après — sécurisé
// Authorization: "Matrix-based security" ou "Role-Based Strategy"
// - Anonymous: aucun droit
// - Authenticated users: Overall/Read uniquement
// - Admins: droits complets, scopés à un groupe restreint
// Security Realm: authentification obligatoire (LDAP/SSO), pas de compte par défaut
```

## Credentials codés en dur dans un Jenkinsfile
```groovy
// Avant — vulnérable
pipeline {
    agent any
    stages {
        stage('Deploy') {
            steps {
                sh 'curl -u admin:P@ssw0rd123 https://api.exemple.com/deploy'
            }
        }
    }
}

// Après — sécurisé
pipeline {
    agent any
    stages {
        stage('Deploy') {
            steps {
                withCredentials([usernamePassword(credentialsId: 'deploy-prod', usernameVariable: 'USR', passwordVariable: 'PWD')]) {
                    sh 'curl -u $USR:$PWD https://api.exemple.com/deploy'
                }
            }
        }
    }
}
// credential 'deploy-prod' scopé au dossier/job, jamais global
```

## Console Script accessible et exposition réseau
```groovy
// Avant — vulnérable
// /script accessible à tout utilisateur authentifié (pas seulement les admins)
// Jenkins exposé directement sur Internet (0.0.0.0:8080), pas de VPN/allowlist IP

// Après — sécurisé
// Accès à /script restreint au rôle "Administer" uniquement (Role-Based Strategy)
// Jenkins accessible uniquement via VPN ou allowlist IP, reverse proxy avec TLS obligatoire
// Mise à jour régulière de Jenkins core et des plugins (Manage Jenkins > Plugin Manager)
```

## Checklist de vérification post-patch
- [ ] L'authentification est obligatoire et la matrice d'autorisation ne permet plus l'accès anonyme.
- [ ] Aucun credential n'est codé en dur dans un `Jenkinsfile` ou un script de job.
- [ ] Les credentials sont scopés au dossier/job via `withCredentials`, pas globaux.
- [ ] La console Script (`/script`) est inaccessible aux utilisateurs non-admin.
- [ ] Jenkins n'est plus exposé directement sur Internet sans VPN/allowlist IP.
- [ ] Jenkins core et les plugins sont à jour (aucune CVE connue non corrigée).
