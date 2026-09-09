# SMTP Injection / En-têtes email (CWE-93)

La version vulnérable insérait le champ `name` directement dans les en-têtes SMTP construits par concaténation, exposant l'application à une injection CRLF permettant l'ajout d'en-têtes arbitraires (Cc, Bcc, second corps de message). La correction utilise la bibliothèque `lettre` qui structure les en-têtes via une API dédiée, et filtre explicitement tout caractère `\r`/`\n` des champs libres avant usage, neutralisant CWE-93.
