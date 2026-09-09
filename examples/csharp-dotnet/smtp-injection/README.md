# SMTP Injection / Email Header Injection (CWE-93)

La version vulnérable insère le champ `name` fourni par l'utilisateur directement dans l'en-tête `Reply-To` et dans le corps du message, sans filtrer les caractères `\r`/`\n`, permettant à un attaquant d'ajouter des en-têtes arbitraires (`Bcc`, `Cc`) ou un second corps de message. La correction retire systématiquement les caractères de contrôle (`\r`, `\n`) des champs libres via `StripControlChars` avant toute insertion dans un en-tête ou le corps du message, en complément de l'échappement natif fourni par `MailMessage`.
