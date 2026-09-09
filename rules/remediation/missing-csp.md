# Remédiation — Content-Security-Policy manquant

## Principe
Ajouter un en-tête `Content-Security-Policy` restrictif sur toutes les réponses HTTP pour limiter les sources de scripts, styles, images et connexions autorisées, réduisant fortement l'impact d'une éventuelle faille XSS. Commencer par un mode `Report-Only` en staging avant application stricte en production.

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

    add_header Content-Security-Policy "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data:; connect-src 'self'; frame-ancestors 'self'" always;
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
    Header always set Content-Security-Policy "default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data:; connect-src 'self'; frame-ancestors 'self'"
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header("Content-Security-Policy: default-src 'self'; script-src 'self'; style-src 'self'; img-src 'self' data:; connect-src 'self'; frame-ancestors 'self'");
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître un en-tête `content-security-policy` non vide.
- [ ] La politique évite `'unsafe-inline'` et `'unsafe-eval'` ; les scripts inline nécessaires utilisent nonce ou hash.
- [ ] La CSP a été validée en mode `Content-Security-Policy-Report-Only` en staging avant bascule stricte.
- [ ] Aucune ressource légitime (scripts, styles, images, appels API) n'est bloquée en production après activation.
- [ ] La directive `frame-ancestors` est cohérente avec la politique anti-clickjacking (`X-Frame-Options`).
