// Faille : SMTP Injection - Email Header Injection (CWE-93)
// Les champs "name" et "subject" du formulaire de contact sont insérés
// directement dans les en-têtes de l'email. Un attaquant peut injecter
// des séquences CR/LF pour ajouter des en-têtes arbitraires (Cc, Bcc)
// ou un second corps de message, transformant l'application en relais
// de spam/phishing.
const express = require('express');
const nodemailer = require('nodemailer');
const router = express.Router();

const transporter = nodemailer.createTransport({ host: 'smtp.example.com', port: 587 });

router.post('/contact', (req, res) => {
  const { name, subject, body } = req.body;

  // Champs utilisateur insérés tels quels dans les en-têtes : dangereux.
  transporter.sendMail(
    {
      from: `"${name}" <contact@example.com>`,
      subject,
      to: 'dest@example.com',
      text: body,
    },
    (err) => {
      if (err) return res.status(500).json({ error: "Erreur d'envoi" });
      res.json({ status: 'envoyé' });
    }
  );
});

module.exports = router;
