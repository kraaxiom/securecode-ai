// Fix : SSRF aveugle (Blind SSRF) (CWE-918)
// Mêmes contrôles qu'une SSRF classique (whitelist, résolution/validation
// d'IP, timeout), appliqués même si le résultat n'est jamais renvoyé au
// client. Les tentatives de destination interdite sont journalisées pour
// permettre la détection a posteriori d'un scan interne.
const express = require('express');
const dns = require('dns').promises;
const { isIP } = require('net');
const ipaddr = require('ipaddr.js');
const router = express.Router();

const ALLOWED_HOSTS = ['hooks.partenaire.example.com'];

function isPrivateOrReserved(ip) {
  if (!isIP(ip)) return true;
  return ipaddr.parse(ip).range() !== 'unicast';
}

async function safeValidateWebhook(rawUrl) {
  const parsed = new URL(rawUrl);
  if (parsed.protocol !== 'https:' || !ALLOWED_HOSTS.includes(parsed.hostname)) {
    throw new Error('URL non autorisée');
  }
  const { address } = await dns.lookup(parsed.hostname);
  if (isPrivateOrReserved(address)) {
    throw new Error('Destination interdite');
  }
  return fetch(parsed.toString(), {
    method: 'HEAD',
    redirect: 'manual',
    signal: AbortSignal.timeout(5000),
  });
}

router.post('/webhooks/register', (req, res) => {
  safeValidateWebhook(req.body.callback_url)
    .then(() => console.log('Webhook validé'))
    .catch((err) => console.warn('Webhook rejeté (audit)', err.message, req.body.callback_url));

  res.status(202).json({ status: 'en cours de validation' });
});

module.exports = router;
