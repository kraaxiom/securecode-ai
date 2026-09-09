# Remédiation — Cookie Poisoning

## Principe
Ne jamais faire confiance à la valeur brute d'un cookie pour une décision de sécurité ou une logique métier. Utiliser le mécanisme de session natif du framework (identifiant opaque côté client, état stocké côté serveur) et, si un cookie applicatif doit porter une donnée, la signer (HMAC) ou la chiffrer avec vérification systématique côté serveur.

## PHP
```php
// Avant — vulnérable
$role = $_COOKIE['role'] ?? 'user';
if ($role === 'admin') {
    afficherPanneauAdmin();
}

// Après — sécurisé
session_start();
$role = $_SESSION['role'] ?? 'user'; // rôle stocké côté serveur, pas dans un cookie lisible
if ($role === 'admin') {
    afficherPanneauAdmin();
}
```

## Node.js (Express)
```js
// Avant — vulnérable
const role = req.cookies.role || 'user';
if (role === 'admin') {
  renderAdminPanel();
}

// Après — sécurisé
// express-session stocke l'état côté serveur, le cookie ne contient qu'un ID opaque signé
const role = req.session.role || 'user';
if (role === 'admin') {
  renderAdminPanel();
}
```

## Python (Flask/Django)
```python
# Avant — vulnérable
role = request.cookies.get('role', 'user')
if role == 'admin':
    show_admin_panel()

# Après — sécurisé
# Flask: utiliser la session signée (itsdangerous) au lieu d'un cookie applicatif brut
role = session.get('role', 'user')
if role == 'admin':
    show_admin_panel()
```

## Checklist de vérification post-patch
- [ ] Aucune donnée sensible (rôle, ID utilisateur, prix) n'est lue directement depuis un cookie sans signature ni revérification serveur.
- [ ] Le rôle/privilège est revalidé en base de données à chaque requête sensible, pas uniquement lu depuis la session.
- [ ] Les cookies applicatifs personnalisés restants sont signés (HMAC) ou chiffrés avec authentification.
- [ ] Un test confirme qu'une modification manuelle du cookie côté client n'accorde aucun privilège supplémentaire.
