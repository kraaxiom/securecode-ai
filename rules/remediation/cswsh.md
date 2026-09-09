# Remédiation — Cross-Site WebSocket Hijacking (CSWSH)

## Principe
Ne jamais authentifier un handshake WebSocket sur le seul cookie de session. Valider strictement l'en-tête `Origin` contre une liste blanche fermée et exiger un jeton d'authentification distinct transmis explicitement (pas automatiquement envoyé par le navigateur comme un cookie).

## PHP (Ratchet / ws)
```php
// Avant — vulnérable : aucune vérification d'Origin, auth uniquement par cookie de session
class ChatServer implements MessageComponentInterface {
    public function onOpen(ConnectionInterface $conn) {
        $this->clients->attach($conn);
    }
}

// Après — sécurisé
class ChatServer implements MessageComponentInterface {
    private array $allowedOrigins = ['https://app.example.com'];

    public function onOpen(ConnectionInterface $conn) {
        $origin = $conn->httpRequest->getHeaderLine('Origin');
        if (!in_array($origin, $this->allowedOrigins, true)) {
            $conn->close(1008);
            return;
        }
        $token = $conn->httpRequest->getUri()->getQuery();
        parse_str($token, $params);
        if (!isset($params['auth_token']) || !TokenService::verify($params['auth_token'])) {
            $conn->close(1008);
            return;
        }
        $this->clients->attach($conn);
    }
}
```

## Node.js (ws)
```js
// Avant — vulnérable
const wss = new WebSocket.Server({ port: 8080 });
wss.on('connection', (ws, req) => {
  // aucune vérification d'origine ni de jeton dédié
  wss.emit('ready', ws);
});

// Après — sécurisé
const ALLOWED_ORIGINS = new Set(['https://app.example.com']);

const server = http.createServer();
const wss = new WebSocket.Server({ noServer: true });

server.on('upgrade', (req, socket, head) => {
  const origin = req.headers.origin;
  if (!ALLOWED_ORIGINS.has(origin)) {
    socket.destroy();
    return;
  }
  const token = new URL(req.url, 'http://localhost').searchParams.get('auth_token');
  verifyToken(token)
    .then((user) => {
      wss.handleUpgrade(req, socket, head, (ws) => {
        ws.user = user;
        wss.emit('connection', ws, req);
      });
    })
    .catch(() => socket.destroy());
});
```

## Checklist de vérification post-patch
- [ ] L'en-tête `Origin` est comparé par égalité stricte à une liste blanche fixe côté serveur, pour chaque handshake.
- [ ] L'authentification du canal WebSocket ne repose pas uniquement sur le cookie de session envoyé automatiquement par le navigateur.
- [ ] Un jeton signé à courte durée de vie est requis et vérifié avant d'accepter la connexion.
- [ ] Une connexion provenant d'une origine non autorisée est fermée avec le code 1008 et journalisée.
- [ ] Un test confirme qu'une page hébergée sur un domaine tiers ne peut pas établir de connexion authentifiée.
