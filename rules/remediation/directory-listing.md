# Remédiation — Listing de répertoire activé

## Principe
Désactiver explicitement le listing de répertoire au niveau du serveur web et placer les fichiers non destinés au public en dehors du webroot.

## Nginx
```nginx
# Avant — vulnérable
server {
    listen 80;
    root /var/www/html;

    location /uploads/ {
        autoindex on;
    }
}

# Après — sécurisé
server {
    listen 80;
    root /var/www/html;

    location /uploads/ {
        autoindex off;
    }
}
```

## Apache
```apache
# Avant — vulnérable
<Directory /var/www/html/uploads>
    Options +Indexes +FollowSymLinks
    AllowOverride None
    Require all granted
</Directory>

# Après — sécurisé
<Directory /var/www/html/uploads>
    Options -Indexes +FollowSymLinks
    AllowOverride None
    Require all granted
</Directory>
```

## Checklist de vérification post-patch
- [ ] Une requête sur un répertoire sans fichier index renvoie une erreur 403/404, pas une page de listing.
- [ ] `autoindex off;` (Nginx) ou `Options -Indexes` (Apache) est appliqué sur l'ensemble des blocs/répertoires du webroot, pas seulement le répertoire testé.
- [ ] Les répertoires sensibles (uploads, cache, logs) contiennent un fichier index vide ou une règle de refus explicite en défense en profondeur.
- [ ] Les fichiers non destinés au public sont déplacés en dehors du webroot lorsque possible.
