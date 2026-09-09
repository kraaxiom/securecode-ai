## SMTP Injection - Email Header Injection (CWE-93)

Le code vulnérable insère les champs `name` et `subject` du formulaire de contact directement dans les en-têtes de l'email envoyé via Nodemailer. Un attaquant injectant `\r\n` peut ajouter des en-têtes arbitraires (`Cc`, `Bcc`) ou un second corps de message, transformant l'application en relais de spam. La correction supprime tout caractère de contrôle (`\r`, `\n`) des champs libres avant construction du message, en complément de l'échappement déjà assuré par Nodemailer.
