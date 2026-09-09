# Remédiation — X-Frame-Options manquant

## Principe
Ajouter l'en-tête `X-Frame-Options` (et/ou la directive CSP `frame-ancestors`) sur toutes les réponses HTTP pour empêcher un site tiers d'intégrer la page dans une `<iframe>` et de mener une attaque de clickjacking.

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

    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header Content-Security-Policy "frame-ancestors 'self'" always;
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
    Header always set X-Frame-Options "SAMEORIGIN"
    Header always set Content-Security-Policy "frame-ancestors 'self'"
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header('X-Frame-Options: SAMEORIGIN');
header("Content-Security-Policy: frame-ancestors 'self'");
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître `x-frame-options: SAMEORIGIN` (ou `DENY` selon le besoin métier).
- [ ] L'en-tête est présent sur toutes les pages sensibles (changement de mot de passe, validation, transfert).
- [ ] La directive `frame-ancestors` en CSP est cohérente avec `X-Frame-Options` (défense en profondeur).
- [ ] Un test manuel confirme qu'une tentative d'intégration de la page dans une iframe d'un domaine externe échoue.
- [ ] Aucune régression constatée si l'application a besoin d'être elle-même embarquée dans un widget légitime (whitelist explicite au lieu de DENY total).
