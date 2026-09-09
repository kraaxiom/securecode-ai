# Remédiation — Error-based SQL Injection

## Principe
Utiliser systématiquement des requêtes préparées avec paramètres liés, et désactiver l'affichage des erreurs SQL détaillées en production. Le détail de l'exception doit être journalisé côté serveur uniquement, jamais renvoyé au client.

## PHP (PDO)
```php
// Avant — vulnérable
try {
    $id = $_GET['id'];
    $result = $pdo->query("SELECT * FROM orders WHERE id = $id");
} catch (PDOException $e) {
    echo "Erreur SQL: " . $e->getMessage();
}

// Après — sécurisé
try {
    $stmt = $pdo->prepare("SELECT * FROM orders WHERE id = :id");
    $stmt->execute(['id' => (int) $_GET['id']]);
    $result = $stmt->fetchAll();
} catch (PDOException $e) {
    error_log($e->getMessage());
    http_response_code(500);
    echo "Une erreur est survenue.";
}
```

## Node.js (mysql2)
```js
// Avant — vulnérable
app.get('/orders/:id', (req, res) => {
  connection.query(`SELECT * FROM orders WHERE id = ${req.params.id}`, (err, rows) => {
    if (err) return res.status(500).json({ error: err.message });
    res.json(rows);
  });
});

// Après — sécurisé
app.get('/orders/:id', (req, res) => {
  connection.query('SELECT * FROM orders WHERE id = ?', [Number(req.params.id)], (err, rows) => {
    if (err) {
      logger.error(err);
      return res.status(500).json({ error: 'Une erreur est survenue.' });
    }
    res.json(rows);
  });
});
```

## Python (psycopg2)
```python
# Avant — vulnérable
try:
    order_id = request.args.get("id")
    cursor.execute(f"SELECT * FROM orders WHERE id = {order_id}")
except Exception as e:
    return jsonify({"error": str(e)}), 500

# Après — sécurisé
try:
    order_id = int(request.args.get("id"))
    cursor.execute("SELECT * FROM orders WHERE id = %s", [order_id])
except Exception as e:
    app.logger.error(e)
    return jsonify({"error": "Une erreur est survenue."}), 500
```

## Checklist de vérification post-patch
- [ ] Aucune concaténation de variable restante dans une requête SQL du fichier corrigé.
- [ ] Le mode debug/affichage détaillé des erreurs est désactivé en production (`display_errors=Off`, `DEBUG=False`, etc.).
- [ ] Aucun message d'exception natif du driver SQL n'est renvoyé au client — uniquement un message générique.
- [ ] Le détail de l'erreur est journalisé côté serveur pour investigation.
- [ ] Un test confirme qu'une entrée provoquant une erreur de syntaxe SQL ne fuit aucune information dans la réponse HTTP.
