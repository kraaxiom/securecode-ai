# Remédiation — Cross-Origin-Opener-Policy (COOP) manquant

## Principe
Ajouter l'en-tête `Cross-Origin-Opener-Policy` sur toutes les réponses HTTP pour isoler le contexte de navigation de la page vis-à-vis des fenêtres tierces ouvertes via `window.open` ou `target="_blank"`, et prévenir les attaques de type tabnabbing.

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

    add_header Cross-Origin-Opener-Policy "same-origin" always;
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
    Header always set Cross-Origin-Opener-Policy "same-origin"
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header('Cross-Origin-Opener-Policy: same-origin');
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître `cross-origin-opener-policy: same-origin`.
- [ ] L'en-tête est présent sur toutes les routes de l'application, pas uniquement la page d'accueil.
- [ ] Les flux nécessitant des popups tiers légitimes (paiement, SSO) sont testés avec `same-origin-allow-popups` si besoin.
- [ ] Les liens `target="_blank"` vers des domaines externes portent `rel="noopener noreferrer"` en complément.
- [ ] Aucune régression fonctionnelle constatée sur les flux d'ouverture de fenêtres après activation (test en staging).
