# Remédiation — SMTP Injection (Email Header Injection)

## Principe
Ne jamais construire les en-têtes d'un email (`From`, `Subject`, `To`) par concaténation directe d'une entrée utilisateur. Utiliser une bibliothèque d'envoi d'email mature qui échappe automatiquement les en-têtes, et rejeter tout caractère de contrôle (`\r`, `\n`) dans les champs libres.

## PHP (PHPMailer plutôt que mail())
```php
// Avant — vulnérable
$name = $_POST['name'];
$subject = $_POST['subject'];
$headers = "From: contact@example.com\r\nReply-To: $name\r\n";
mail('dest@example.com', $subject, $body, $headers);

// Après — sécurisé
use PHPMailer\PHPMailer\PHPMailer;

$name = str_replace(["\r", "\n"], '', $_POST['name']);
$subject = str_replace(["\r", "\n"], '', $_POST['subject']);

$mail = new PHPMailer();
$mail->setFrom('contact@example.com');
$mail->addReplyTo($name . '@example.com'); // ou champ dédié validé
$mail->Subject = $subject;
$mail->Body = $body;
$mail->send();
```

## JavaScript / Node.js (Nodemailer)
```js
// Avant — vulnérable
const name = req.body.name;
const subject = req.body.subject;
transporter.sendMail({
  from: `"${name}" <contact@example.com>`,
  subject: subject,
  to: 'dest@example.com',
});

// Après — sécurisé : Nodemailer échappe les en-têtes, mais on filtre quand même les CR/LF
const stripControlChars = (s) => s.replace(/[\r\n]/g, '');
const name = stripControlChars(req.body.name);
const subject = stripControlChars(req.body.subject);
transporter.sendMail({
  from: `"${name}" <contact@example.com>`,
  subject,
  to: 'dest@example.com',
});
```

## Python (smtplib -> préférer une lib de haut niveau)
```python
# Avant — vulnérable
name = request.form['name']
subject = request.form['subject']
message = f"From: contact@example.com\r\nSubject: {subject}\r\n\r\n{body}"
server.sendmail('contact@example.com', 'dest@example.com', message)

# Après — sécurisé : email.message + rejet des caractères de contrôle
from email.message import EmailMessage

def strip_control_chars(value: str) -> str:
    return value.replace('\r', '').replace('\n', '')

name = strip_control_chars(request.form['name'])
subject = strip_control_chars(request.form['subject'])

msg = EmailMessage()
msg['From'] = 'contact@example.com'
msg['To'] = 'dest@example.com'
msg['Subject'] = subject  # email.message échappe/plie correctement les en-têtes
msg.set_content(body)
server.send_message(msg)
```

## Checklist de vérification post-patch
- [ ] Aucun en-tête email n'est construit par concaténation manuelle d'une entrée utilisateur.
- [ ] Tout caractère `\r` ou `\n` est rejeté ou supprimé des champs utilisés dans les en-têtes.
- [ ] Une bibliothèque d'email dédiée (PHPMailer, Nodemailer, `email.message`, Django `send_mail`) est utilisée plutôt que `mail()`/`smtplib` bas niveau.
- [ ] Le format de l'adresse email destinataire/expéditeur est validé strictement quand il provient de l'utilisateur.
- [ ] Un test confirme qu'une entrée contenant un saut de ligne ne permet pas d'ajouter un en-tête (`Cc`, `Bcc`) ou un second corps de message.
