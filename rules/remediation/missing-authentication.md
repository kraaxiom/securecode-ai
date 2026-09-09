# Remédiation — Absence d'authentification sur un canal WebSocket

## Principe
Exiger une authentification explicite avant ou pendant le handshake WebSocket, avant d'accepter tout message, et réappliquer une vérification d'autorisation par type de message/action plutôt qu'une seule fois à la connexion.

## PHP (Ratchet)
```php
// Avant — vulnérable : aucune vérification d'identité
public function onOpen(ConnectionInterface $conn) {
    $this->clients->attach($conn);
}
public function onMessage(ConnectionInterface $from, $msg) {
    $data = json_decode($msg, true);
    if ($data['action'] === 'delete_account') {
        $this->accountService->delete($data['userId']);
    }
}

// Après — sécurisé
public function onOpen(ConnectionInterface $conn) {
    $token = $this->extractToken($conn->httpRequest);
    $user = TokenService::verify($token);
    if (!$user) {
        $conn->close(1008);
        return;
    }
    $conn->user = $user;
    $this->clients->attach($conn);
}
public function onMessage(ConnectionInterface $from, $msg) {
    if (!isset($from->user)) {
        $from->close(1008);
        return;
    }
    $data = json_decode($msg, true);
    if ($data['action'] === 'delete_account') {
        if ($from->user->id !== $data['userId'] && !$from->user->isAdmin) {
            $from->close(1008);
            return;
        }
        $this->accountService->delete($data['userId']);
    }
}
```

## Node.js (ws)
```js
// Avant — vulnérable
wss.on('connection', (ws) => {
  ws.on('message', (raw) => {
    const data = JSON.parse(raw);
    if (data.action === 'deleteAccount') {
      accountService.delete(data.userId); // aucune vérification d'identité
    }
  });
});

// Après — sécurisé
server.on('upgrade', (req, socket, head) => {
  verifyToken(extractToken(req))
    .then((user) => {
      wss.handleUpgrade(req, socket, head, (ws) => {
        ws.user = user;
        wss.emit('connection', ws, req);
      });
    })
    .catch(() => socket.destroy());
});

wss.on('connection', (ws) => {
  ws.on('message', (raw) => {
    if (!ws.user) return ws.close(1008);
    const data = JSON.parse(raw);
    if (data.action === 'deleteAccount') {
      if (ws.user.id !== data.userId && !ws.user.isAdmin) {
        return ws.close(1008);
      }
      accountService.delete(data.userId);
    }
  });
});
```

## Checklist de vérification post-patch
- [ ] Aucune connexion n'est acceptée sans vérification préalable d'un jeton ou d'une session valide.
- [ ] Chaque action sensible reçue via un message revérifie l'autorisation (pas seulement l'authentification à la connexion).
- [ ] Une connexion non authentifiée ou expirée est fermée avec le code 1008.
- [ ] Un test E2E confirme qu'une action sensible est refusée sans jeton valide.
- [ ] Les logs d'audit enregistrent l'identité de l'appelant pour chaque action sensible exécutée via WebSocket.
