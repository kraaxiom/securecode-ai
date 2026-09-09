# Remédiation — Referrer-Policy manquant

## Principe
Ajouter l'en-tête `Referrer-Policy` sur toutes les réponses HTTP pour contrôler quelles informations d'URL sont transmises dans l'en-tête `Referer` lors de la navigation vers un site tiers ou du chargement de ressources externes, évitant la fuite de tokens ou identifiants présents en query string.

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

    add_header Referrer-Policy "strict-origin-when-cross-origin" always;
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
    Header always set Referrer-Policy "strict-origin-when-cross-origin"
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header('Referrer-Policy: strict-origin-when-cross-origin');
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître `referrer-policy: strict-origin-when-cross-origin`.
- [ ] Aucune donnée sensible (token, identifiant de session) n'est présente dans les paramètres d'URL de pages contenant des liens externes.
- [ ] Les pages avec des flux critiques (réinitialisation de mot de passe, confirmation d'email) utilisent une politique plus stricte (`no-referrer`) si nécessaire.
- [ ] Les intégrations analytics/tierces continuent de fonctionner correctement après application de la politique.
- [ ] Test en staging confirmant l'absence de régression sur les liens sortants et ressources externes.
