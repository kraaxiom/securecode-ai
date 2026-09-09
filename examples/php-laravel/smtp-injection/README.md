# SMTP Injection — Email Header Injection (CWE-93)

Le code vulnérable construit manuellement les en-têtes email (`From`, `Reply-To`) par concaténation du champ `name` avant de les passer à `mail()`, ce qui permet à un attaquant d'injecter des séquences CR/LF pour ajouter des en-têtes arbitraires. La correction utilise la façade `Mail` de Laravel (Symfony Mailer, qui échappe les en-têtes automatiquement) et supprime en complément tout caractère `\r`/`\n` des champs libres (CWE-93 : Improper Neutralization of CRLF Sequences).
