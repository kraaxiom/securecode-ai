# Remédiation — Missing Audit (absence de journalisation d'audit)

## Principe
Activer la journalisation native pour les opérations sensibles, journaliser au niveau applicatif les accès aux données sensibles, et stocker les journaux d'audit dans un système séparé, en écriture seule pour les comptes applicatifs.

## PHP
```php
// Avant — vulnérable : accès à des données personnelles sans trace
public function getUserProfile($id) {
    return $this->pdo->query("SELECT * FROM users WHERE id = $id")->fetch();
}

// Après — sécurisé : journalisation applicative des accès sensibles
public function getUserProfile(PDO $pdo, PDO $auditPdo, int $requesterId, int $targetId) {
    $stmt = $pdo->prepare("SELECT * FROM users WHERE id = ?");
    $stmt->execute([$targetId]);
    $profile = $stmt->fetch();

    // Journal séparé, écriture seule pour le compte applicatif
    $auditPdo->prepare(
        "INSERT INTO audit_log (actor_id, action, target_id, occurred_at) VALUES (?, 'view_profile', ?, NOW())"
    )->execute([$requesterId, $targetId]);

    return $profile;
}
```

## JS / Node
```js
// Avant — vulnérable
app.get('/admin/users/:id', async (req, res) => {
  const user = await db.query('SELECT * FROM users WHERE id = $1', [req.params.id]);
  res.json(user.rows[0]); // aucune trace de l'accès admin
});

// Après — sécurisé
app.get('/admin/users/:id', requireAdmin, async (req, res) => {
  const user = await db.query('SELECT * FROM users WHERE id = $1', [req.params.id]);
  await auditDb.query(
    'INSERT INTO audit_log (actor_id, action, target_id, occurred_at) VALUES ($1, $2, $3, now())',
    [req.user.id, 'view_user', req.params.id]
  ); // base d'audit séparée, compte applicatif en écriture seule
  res.json(user.rows[0]);
});
```

## Python
```python
# Avant — vulnérable
@app.route('/admin/users/<int:user_id>')
def get_user(user_id):
    return db.fetch_one("SELECT * FROM users WHERE id = %s", [user_id])

# Après — sécurisé
@app.route('/admin/users/<int:user_id>')
@require_admin
def get_user(user_id):
    user = db.fetch_one("SELECT * FROM users WHERE id = %s", [user_id])
    audit_db.execute(
        "INSERT INTO audit_log (actor_id, action, target_id, occurred_at) VALUES (%s, %s, %s, now())",
        [current_user.id, "view_user", user_id],
    )  # journal séparé, non modifiable par le compte applicatif standard
    return user
```

## Checklist de vérification post-patch
- [ ] La journalisation native de la base de données est activée pour les opérations sensibles (connexions, DDL, changements de privilèges).
- [ ] Les accès en lecture aux données sensibles sont journalisés avec identité, horodatage et nature de l'opération.
- [ ] Les journaux d'audit sont stockés dans un système/compte séparé, en écriture seule pour les comptes applicatifs.
- [ ] Une politique de rétention et une revue périodique des journaux d'audit sont définies et documentées.
