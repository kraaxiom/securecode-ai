// Correction : HTTP Request Smuggling (CWE-444)
// Toute requête présentant simultanément Content-Length et Transfer-Encoding
// est rejetée en défense en profondeur, même si le proxy en amont devrait
// déjà filtrer ce cas ambigu.
const http = require('http');

const server = http.createServer((req, res) => {
  const hasContentLength = 'content-length' in req.headers;
  const hasTransferEncoding = 'transfer-encoding' in req.headers;

  // Rejet explicite des requêtes ambiguës avant tout traitement.
  if (hasContentLength && hasTransferEncoding) {
    res.writeHead(400, { 'Content-Type': 'text/plain' });
    return res.end('Requête ambiguë refusée');
  }

  handleRequest(req, res);
});

function handleRequest(req, res) {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('OK');
}

module.exports = server;
