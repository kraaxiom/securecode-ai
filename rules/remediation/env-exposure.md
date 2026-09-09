# Remédiation — Fichier .env exposé

## Principe
Placer le fichier `.env` en dehors du webroot lorsque c'est possible, et bloquer explicitement l'accès aux dotfiles au niveau du serveur web en défense en profondeur. Toujours faire tourner les secrets d'un `.env` ayant été exposé.

## Nginx
```nginx
# Avant — vulnérable
server {
    listen 80;
    root /var/www/html; # contient .env
    index index.php;
}

# Après — sécurisé
server {
    listen 80;
    root /var/www/html;
    index index.php;

    location ~ /\.(?!well-known) {
        deny all;
        return 404;
    }
}
```

## Apache
```apache
# Avant — vulnérable
DocumentRoot /var/www/html

# Après — sécurisé
<FilesMatch "^\.">
    Require all denied
</FilesMatch>
```

## Checklist de vérification post-patch
- [ ] Le fichier `.env` est déplacé en dehors du webroot (ou, à défaut, une règle serveur bloque explicitement son accès).
- [ ] Une requête sur `/.env` renvoie 403/404.
- [ ] Toutes les règles serveur bloquent l'ensemble des dotfiles (`.env`, `.git`, `.htaccess`), pas uniquement `.env`.
- [ ] Tous les secrets présents dans le `.env` exposé (identifiants DB, clés API, secrets de session) ont été révoqués et régénérés.
- [ ] Un contrôle automatisé post-déploiement vérifie que `.env` n'est pas accessible publiquement.
