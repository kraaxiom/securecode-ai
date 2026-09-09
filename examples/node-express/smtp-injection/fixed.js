// Correction : SMTP Injection - Email Header Injection (CWE-93)
// Nodemailer échappe déjà les en-têtes, mais on filtre en plus tout
// caractère de contrôle (\r, \n) dans les champs libres avant usage,
// en défense en profondeur.
const express = require('express');
const nodemailer = require('nodemailer');
const router = express.Router();

const transporter = nodemailer.createTransport({ host: 'smtp.example.com', port: 587 });

// Supprime tout caractère de contrôle pouvant injecter un en-tête additionnel.
function stripControlChars(value) {
  return String(value).replace(/[\r\n]/g, '');
}

router.post('/contact', (req, res) => {
  const name = stripControlChars(req.body.name || '');
  const subject = stripControlChars(req.body.subject || '');
  const body = req.body.body;

  if (name.length > 100 || subject.length > 200) {
    return res.status(400).json({ error: 'Champ trop long' });
  }

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
