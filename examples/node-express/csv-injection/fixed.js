// Correction : CSV Injection (CWE-1236)
// Chaque cellule dont le premier caractère est un déclencheur de formule est
// préfixée par une apostrophe, neutralisant son interprétation par le tableur.
const express = require('express');
const router = express.Router();

function sanitizeCsvCell(value) {
  if (typeof value === 'string' && /^[=+\-@\t\r]/.test(value)) {
    return `'${value}`;
  }
  return value || '';
}

router.get('/export/customers', (req, res) => {
  const customers = [
    { name: req.query.name, comment: req.query.comment },
  ];

  // Neutralisation systématique avant écriture dans le CSV.
  const csv = customers
    .map(c => [sanitizeCsvCell(c.name), sanitizeCsvCell(c.comment)].join(','))
    .join('\n');

  res.set('Content-Type', 'text/csv');
  res.set('Content-Disposition', 'attachment; filename="customers.csv"');
  res.send(csv);
});

module.exports = router;
