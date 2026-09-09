// Faille : HTTP Request Smuggling (CWE-444)
// Le serveur applicatif traite la requête sans vérifier la cohérence entre
// les en-têtes Content-Length et Transfer-Encoding, ce qui peut être exploité
// si un proxy en amont interprète différemment les limites de la requête.
const http = require('http');

const server = http.createServer((req, res) => {
  // Aucune vérification de cohérence des en-têtes de délimitation du corps.
  handleRequest(req, res);
});

function handleRequest(req, res) {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('OK');
}

module.exports = server;
