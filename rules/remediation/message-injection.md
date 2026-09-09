# Remédiation — Injection dans les messages WebSocket

## Principe
Traiter tout message WebSocket entrant comme une entrée non fiable : valider son schéma/type/taille, encoder toute donnée rediffusée à d'autres clients, et utiliser des requêtes paramétrées pour toute interaction avec une base de données.

## PHP (Ratchet)
```php
// Avant — vulnérable : rediffusion brute + requête concaténée
public function onMessage(ConnectionInterface $from, $msg) {
    $data = json_decode($msg, true);
    $pdo->query("INSERT INTO messages (content) VALUES ('{$data['text']}')");
    foreach ($this->clients as $client) {
        $client->send($data['text']); // pas d'échappement HTML côté client non plus
    }
}

// Après — sécurisé
public function onMessage(ConnectionInterface $from, $msg) {
    $data = json_decode($msg, true, 512, JSON_THROW_ON_ERROR);
    if (!isset($data['text']) || !is_string($data['text']) || mb_strlen($data['text']) > 500) {
        $from->close(1003);
        return;
    }
    $stmt = $pdo->prepare("INSERT INTO messages (content) VALUES (:text)");
    $stmt->execute(['text' => $data['text']]);

    $safe = htmlspecialchars($data['text'], ENT_QUOTES, 'UTF-8');
    foreach ($this->clients as $client) {
        $client->send(json_encode(['text' => $safe]));
    }
}
```

## Node.js (ws)
```js
// Avant — vulnérable
ws.on('message', (raw) => {
  const data = JSON.parse(raw);
  db.query(`INSERT INTO messages (content) VALUES ('${data.text}')`);
  broadcast(data.text); // rediffusé sans échappement
});

// Après — sécurisé
const schema = z.object({ text: z.string().max(500) });

ws.on('message', (raw) => {
  let data;
  try {
    data = schema.parse(JSON.parse(raw));
  } catch {
    ws.close(1003, 'Invalid payload');
    return;
  }
  db.query('INSERT INTO messages (content) VALUES (?)', [data.text]);
  const safe = escapeHtml(data.text);
  broadcast(JSON.stringify({ text: safe }));
});
```

## Checklist de vérification post-patch
- [ ] Chaque message entrant est validé par un schéma explicite (type, longueur, format) avant tout traitement.
- [ ] Toute écriture en base de données déclenchée par un message utilise une requête paramétrée.
- [ ] Tout contenu utilisateur rediffusé à d'autres clients est échappé selon le contexte de rendu (HTML/JS).
- [ ] Les messages malformés ferment la connexion (code 1003) sans planter le processus serveur.
- [ ] Un test confirme qu'un message contenant `<script>` ou un caractère `'` ne casse ni la base ni l'affichage côté client.
