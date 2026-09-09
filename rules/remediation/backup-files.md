# Remédiation — Fichiers de sauvegarde exposés

## Principe
Interdire au niveau du serveur web l'accès à toute extension/suffixe typique de sauvegarde, ne jamais éditer directement en production, et stocker les sauvegardes en dehors du webroot.

## Nginx
```nginx
# Avant — vulnérable (aucune règle de blocage)
server {
    listen 80;
    root /var/www/html;
    index index.php;
}

# Après — sécurisé
server {
    listen 80;
    root /var/www/html;
    index index.php;

    location ~* \.(bak|old|orig|save|swp|swo|~|zip|tar|tar\.gz|tgz|sql|dist)$ {
        deny all;
        return 404;
    }
}
```

## Apache (.htaccess ou bloc VirtualHost)
```apache
# Avant — vulnérable (aucune règle de blocage)
DocumentRoot /var/www/html

# Après — sécurisé
<FilesMatch "\.(bak|old|orig|save|swp|swo|~|zip|tar|tar\.gz|tgz|sql|dist)$">
    Require all denied
</FilesMatch>
```

## Checklist de vérification post-patch
- [ ] Une requête sur un fichier `.bak`/`.old`/`~`/`.zip` connu dans le webroot renvoie 403/404.
- [ ] Aucun artefact de sauvegarde ou archive de déploiement ne subsiste dans le webroot après un déploiement.
- [ ] Le pipeline de déploiement ne copie que les fichiers applicatifs nécessaires (pas de fichiers temporaires d'éditeur).
- [ ] Les sauvegardes légitimes sont stockées en dehors du webroot, sur un stockage dédié et chiffré.
