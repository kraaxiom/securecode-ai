# Remédiation — X-Content-Type-Options manquant

## Principe
Ajouter l'en-tête `X-Content-Type-Options: nosniff` sur toutes les réponses HTTP pour empêcher le navigateur de deviner (MIME sniffing) le type réel d'un contenu et de le réinterpréter comme exécutable, notamment sur les endpoints servant des fichiers uploadés par les utilisateurs.

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

    add_header X-Content-Type-Options "nosniff" always;
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
    Header always set X-Content-Type-Options "nosniff"
</VirtualHost>
```

## PHP (niveau applicatif)
```php
// Avant — vulnérable : aucun en-tête défini
<?php
echo "Bienvenue";

// Après — sécurisé
<?php
header('X-Content-Type-Options: nosniff');
echo "Bienvenue";
```

## Checklist de vérification post-patch
- [ ] `curl -I https://example.com/` fait apparaître `x-content-type-options: nosniff`.
- [ ] L'en-tête est présent en particulier sur les routes de téléchargement/servage de fichiers uploadés.
- [ ] Le `Content-Type` de chaque réponse (surtout fichiers uploadés) est déterminé par validation du contenu réel, pas par l'extension fournie par l'utilisateur.
- [ ] Les fichiers utilisateur potentiellement dangereux sont servis depuis un domaine séparé sans cookies de session.
- [ ] Aucune régression constatée sur le rendu des ressources statiques (CSS, JS, images) après activation.
