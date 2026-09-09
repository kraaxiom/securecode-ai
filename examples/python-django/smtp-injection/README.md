# SMTP Injection — CWE-93

Le code vulnérable construit manuellement les en-têtes SMTP (`From`, `Reply-To`, `Subject`) par concaténation de champs utilisateur, permettant l'injection de CR/LF pour ajouter des en-têtes ou un second corps de message. La correction utilise `django.core.mail.send_mail`, qui échappe correctement les en-têtes, et retire explicitement les caractères de contrôle des champs libres avant usage. Voir `rules/remediation/smtp-injection.md` pour d'autres langages.
