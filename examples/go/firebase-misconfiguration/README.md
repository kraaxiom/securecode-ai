# Mauvaise configuration Firebase (CWE-284)

## Description de la vulnerabilite
Le fichier `vulnerable.go` illustre un script de deploiement Go, utilisant l'API Firebase Rules (`google.golang.org/api/firebaserules/v1`), qui publie des regles de securite Firestore en "mode test" (`allow read, write: if true;`). Ces regles, generees automatiquement a la creation d'un projet Firebase pour faciliter le prototypage, autorisent n'importe qui — authentifie ou non — a lire et ecrire l'integralite de la base de donnees si elles restent actives en production.

## CWE reel utilise
**CWE-284 : Improper Access Control** (source : `knowledge/cloud/firebase-misconfiguration.md`).

## Pourquoi c'est dangereux
Comme la logique d'acces Firebase s'execute souvent cote client, des regles laxistes exposent directement toute la base de donnees a Internet, sans passer par un serveur applicatif intermediaire qui pourrait filtrer les requetes. Un attaquant peut alors lire, modifier ou supprimer n'importe quelle donnee de n'importe quel utilisateur simplement en interrogeant l'API Firestore publique du projet.

## Comment le correctif fonctionne
Le fichier `fixed.go` deploie des regles qui :
- Exigent `request.auth != null` (l'utilisateur doit etre authentifie) ET verifient la propriete de la ressource (`request.auth.uid == userId`), empechant un utilisateur connecte d'acceder aux donnees d'un autre.
- Valident le schema des donnees a la creation (`request.resource.data.keys().hasOnly([...])`), empechant l'injection de champs arbitraires non prevus.
- Sont deployees via la meme API `firebaserules.Ruleset`/`Release`, garantissant que le mode test ne reste jamais actif en production.

Ces regles doivent etre testees avec l'emulateur Firebase (cas autorise + cas refuse) avant tout deploiement en production.

## References
- OWASP Cloud Security Cheat Sheet
- Firebase Docs: Security Rules
- CWE-284: Improper Access Control — https://cwe.mitre.org/data/definitions/284.html
- Voir egalement `rules/remediation/firebase-misconfiguration.md`
