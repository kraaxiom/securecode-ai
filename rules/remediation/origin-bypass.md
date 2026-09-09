# Remédiation — Contournement de la vérification d'origine WebSocket

## Principe
Comparer l'en-tête `Origin` par égalité stricte à une liste blanche codée en dur côté serveur. Ne jamais utiliser de sous-chaîne, de regex non ancrée, ou un en-tête contrôlable par le client comme source de vérité.

## PHP (Ratchet)
```php
// Avant — vulnérable : comparaison par sous-chaîne, contournable par attacker-example.com
public function onOpen(ConnectionInterface $conn) {
    $origin = $conn->httpRequest->getHeaderLine('Origin');
    if (strpos($origin, 'example.com') === false) {
        $conn->close(1008);
        return;
    }
    $this->clients->attach($conn);
}

// Après — sécurisé
private const ALLOWED_ORIGINS = ['https://app.example.com', 'https://admin.example.com'];

public function onOpen(ConnectionInterface $conn) {
    $origin = $conn->httpRequest->getHeaderLine('Origin');
    if (!in_array($origin, self::ALLOWED_ORIGINS, true)) {
        $conn->close(1008);
        return;
    }
    $this->clients->attach($conn);
}
```

## Node.js (ws)
```js
// Avant — vulnérable : regex non ancrée, matche attacker.com/example.com.evil.com
server.on('upgrade', (req, socket, head) => {
  const origin = req.headers.origin || '';
  if (!/example\.com/.test(origin)) {
    socket.destroy();
    return;
  }
  wss.handleUpgrade(req, socket, head, (ws) => wss.emit('connection', ws, req));
});

// Après — sécurisé
const ALLOWED_ORIGINS = new Set(['https://app.example.com', 'https://admin.example.com']);

server.on('upgrade', (req, socket, head) => {
  const origin = req.headers.origin;
  if (!origin || !ALLOWED_ORIGINS.has(origin)) {
    socket.destroy();
    return;
  }
  wss.handleUpgrade(req, socket, head, (ws) => wss.emit('connection', ws, req));
});
```

## Checklist de vérification post-patch
- [ ] La vérification d'origine utilise une égalité stricte (`===`/`in_array` strict) sur une liste fermée, jamais `indexOf`/`includes`/`strpos`.
- [ ] Toute regex de validation d'origine, si utilisée, est strictement ancrée (`^...$`) avec les points échappés.
- [ ] La liste des origines autorisées est une constante serveur, jamais dérivée d'un en-tête client (`Host`, `X-Forwarded-Host`).
- [ ] `Origin: null` ou absent est rejeté explicitement sauf besoin métier documenté.
- [ ] Un test confirme qu'un domaine du type `example.com.attacker.com` ou `attacker-example.com` est rejeté.
