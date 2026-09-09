# Remédiation — Cross-Origin-Embedder-Policy (COEP) manquant

## Principe
Ajouter l'en-tête `Cross-Origin-Embedder-Policy` sur toutes les réponses HTTP afin d'activer l'isolation cross-origin, en complément de `Cross-Origin-Opener-Policy`. Vérifier au préalable que toutes les ressources tierces embarquées définissent un `Cross-Origin-Resource-Policy` compatible, sous peine de bloquer leur chargement.

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

    add_header Cross-Origin-Embedder-Policy "require-corp" always;
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
    Header always set Cross-Origin-Embedder-Policy "require-corp"
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header('Cross-Origin-Embedder-Policy: require-corp');
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître `cross-origin-embedder-policy: require-corp`.
- [ ] L'en-tête est présent sur toutes les routes de l'application, pas uniquement la page d'accueil.
- [ ] Toutes les ressources tierces embarquées (images, scripts, iframes) renvoient un `Cross-Origin-Resource-Policy` compatible ou du CORS approprié.
- [ ] Aucune ressource légitime n'est bloquée après activation (test en staging avant production).
- [ ] `Cross-Origin-Opener-Policy: same-origin` est également présent pour une isolation cross-origin complète.
