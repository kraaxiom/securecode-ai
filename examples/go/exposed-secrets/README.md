# Secrets exposes dans le code (CWE-798)

## Description de la vulnerabilite
Le fichier `vulnerable.go` illustre une application Go dans laquelle une cle API Stripe et des identifiants de base de donnees sont codes en dur en tant que constantes dans le code source. Pire encore, la cle API est journalisee en clair via `log.Printf`, ce qui l'expose aussi dans les systemes de collecte de logs. Si ce fichier est committe dans un depot Git (public ou prive partage), le secret doit etre considere compromis, meme apres suppression ulterieure du code.

## CWE reel utilise
**CWE-798 : Use of Hard-coded Credentials** (source : `knowledge/cloud/exposed-secrets.md`).

## Pourquoi c'est dangereux
Un secret code en dur peut fuiter de multiples facons : historique Git, logs applicatifs, captures d'ecran, partage de code, ou simplement lecture du depot par une personne non autorisee. Une fois expose, un attaquant peut l'utiliser pour usurper des appels API facturables (Stripe), acceder directement a la base de donnees de production, ou pivoter vers d'autres ressources cloud si le secret est plus largement privilegie.

## Comment le correctif fonctionne
Le fichier `fixed.go` :
- Supprime toute valeur secrete du code source ; seules des references (noms de secrets) comme `prod/stripe/secret_key` y figurent.
- Recupere les valeurs reelles a l'execution via AWS Secrets Manager (`secretsmanager.GetSecretValue`), qui pourrait etre remplace par Vault, Azure Key Vault ou GCP Secret Manager selon l'environnement.
- Ne journalise plus jamais la valeur du secret — seul un message de confirmation neutre est logue.
- Permet une rotation centralisee du secret sans modification du code applicatif.

En complement du correctif de code, toute fuite confirmee impose de revoquer/regenerer immediatement le secret expose et de purger l'historique Git si necessaire (voir `rules/remediation/exposed-secrets.md`).

## References
- OWASP Top 10 2021 — A05: Security Misconfiguration
- OWASP Cheat Sheet: Secrets Management
- CWE-798: Use of Hard-coded Credentials — https://cwe.mitre.org/data/definitions/798.html
- Voir egalement `rules/remediation/exposed-secrets.md`
