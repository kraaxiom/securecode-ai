# Remédiation — Default Credentials

## Principe
Ne jamais livrer un identifiant/mot de passe par défaut fixe en production. Générer un secret unique et aléatoire à l'installation et forcer son changement immédiat.

## PHP
```php
// Avant — vulnérable : compte admin avec mot de passe par défaut en seed
// database/seeders/AdminSeeder.php
User::create([
    'email' => 'admin@example.com',
    'password' => Hash::make('admin123'),
    'role' => 'admin',
]);

// Après — sécurisé : mot de passe aléatoire, changement forcé à la première connexion
$temporaryPassword = Str::random(24);
$admin = User::create([
    'email' => 'admin@example.com',
    'password' => Hash::make($temporaryPassword),
    'role' => 'admin',
    'must_change_password' => true,
]);
Log::channel('deploy')->info("Compte admin créé, mot de passe temporaire à récupérer via le canal sécurisé de déploiement.");
// Le mot de passe temporaire n'est jamais journalisé en clair ; il est transmis via un canal hors-bande sécurisé.
```

## JS / Node.js
```js
// Avant — vulnérable : identifiants par défaut en dur
const ADMIN_USER = 'admin';
const ADMIN_PASS = 'admin'; // valeur d'exemple laissée telle quelle

// Après — sécurisé : génération aléatoire + changement obligatoire
const crypto = require('crypto');

async function provisionAdminAccount() {
  const tempPassword = crypto.randomBytes(18).toString('base64url');
  const hash = await argon2.hash(tempPassword);
  await db.users.create({
    email: 'admin@example.com',
    passwordHash: hash,
    role: 'admin',
    mustChangePassword: true,
  });
  // Transmission du mot de passe temporaire via canal sécurisé hors-bande, jamais en log.
  return tempPassword;
}
```

## Python
```python
# Avant — vulnérable : identifiants par défaut en dur dans le script de provisioning
DEFAULT_ADMIN_PASSWORD = "changeme"

# Après — sécurisé : génération aléatoire cryptographique + flag de changement obligatoire
import secrets

def provision_admin_account():
    temp_password = secrets.token_urlsafe(24)
    User.objects.create(
        email="admin@example.com",
        password=make_password(temp_password),
        role="admin",
        must_change_password=True,
    )
    # Le mot de passe temporaire est transmis hors-bande, jamais journalisé en clair.
    return temp_password
```

## Checklist de vérification post-patch
- [ ] Aucun identifiant/mot de passe par défaut fixe ne subsiste dans le code, les seeds ou la configuration.
- [ ] Un flag `must_change_password` (ou équivalent) bloque l'accès normal tant que le mot de passe temporaire n'a pas été changé.
- [ ] Le mot de passe temporaire est généré aléatoirement à chaque provisioning, jamais réutilisé.
- [ ] Aucun mot de passe temporaire n'apparaît en clair dans les logs applicatifs.
- [ ] Les comptes de démonstration/test sont désactivés ou absents de la configuration de production.
