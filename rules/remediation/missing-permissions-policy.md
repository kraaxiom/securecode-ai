# Remédiation — Permissions-Policy manquant

## Principe
Ajouter l'en-tête `Permissions-Policy` sur toutes les réponses HTTP pour désactiver explicitement les API sensibles du navigateur (caméra, micro, géolocalisation, etc.) non utilisées par l'application, réduisant la surface d'attaque en cas de XSS ou d'intégration tierce malveillante.

## Nginx
```nginx
# Avant — en-tête absent
server {
    listen 443 ssl;
    server_name example.com;
    root /var/www/html;
}

# Après — en-tête ajouté
server {
    listen 443 ssl;
    server_name example.com;
    root /var/www/html;

    add_header Permissions-Policy "camera=(), microphone=(), geolocation=(), payment=(), usb=()" always;
}
```

## Apache (.htaccess / vhost)
```apache
# Avant — en-tête absent
<VirtualHost *:443>
    DocumentRoot /var/www/html
</VirtualHost>

# Après — en-tête ajouté
<VirtualHost *:443>
    DocumentRoot /var/www/html
    Header always set Permissions-Policy "camera=(), microphone=(), geolocation=(), payment=(), usb=()"
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header('Permissions-Policy: camera=(), microphone=(), geolocation=(), payment=(), usb=()');
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître un en-tête `permissions-policy` listant les fonctionnalités désactivées.
- [ ] La liste des fonctionnalités bloquées correspond aux besoins réels de l'application (rien de nécessaire n'est cassé).
- [ ] Les iframes tierces intégrées portent également un attribut `allow` restrictif cohérent avec la politique serveur.
- [ ] La politique est revue lors de l'ajout de toute nouvelle fonctionnalité front-end utilisant une API navigateur sensible.
- [ ] Aucune régression fonctionnelle constatée après activation (test en staging avant production).
