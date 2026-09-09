# Remédiation — Strict-Transport-Security (HSTS) manquant

## Principe
Ajouter l'en-tête `Strict-Transport-Security` sur toutes les réponses HTTPS pour forcer le navigateur à toujours contacter le site en HTTPS après une première visite réussie, empêchant les attaques de downgrade/SSL stripping. Le compléter par une redirection HTTP → HTTPS côté serveur.

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

    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
}

server {
    listen 80;
    server_name example.com;
    return 301 https://$host$request_uri;
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
    Header always set Strict-Transport-Security "max-age=31536000; includeSubDomains"
</VirtualHost>

<VirtualHost *:80>
    ServerName example.com
    Redirect permanent / https://example.com/
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header('Strict-Transport-Security: max-age=31536000; includeSubDomains');
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître `strict-transport-security: max-age=31536000; includeSubDomains`.
- [ ] La redirection HTTP → HTTPS est bien en place en complément de HSTS, pas en remplacement.
- [ ] Tous les sous-domaines couverts par `includeSubDomains` supportent réellement HTTPS avant activation.
- [ ] `preload` n'est activé que si le domaine est prêt à être inscrit durablement sur `hstspreload.org`.
- [ ] Aucune régression constatée sur les environnements internes/staging utilisant encore HTTP.
