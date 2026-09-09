# Remédiation — Répertoire .svn exposé

## Principe
Utiliser `svn export` (sans métadonnées) plutôt qu'un checkout direct pour le déploiement, et bloquer en défense en profondeur l'accès aux dotdirs au niveau du serveur web.

## Nginx
```nginx
# Avant — vulnérable
server {
    listen 80;
    root /var/www/html; # contient .svn/
    index index.php;
}

# Après — sécurisé
server {
    listen 80;
    root /var/www/html;
    index index.php;

    location ~ /\.svn {
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
<DirectoryMatch "^/var/www/html/\.svn">
    Require all denied
</DirectoryMatch>

RedirectMatch 404 "(^|/)\.svn"
```

## Checklist de vérification post-patch
- [ ] Une requête sur `/.svn/entries` et `/.svn/wc.db` renvoie 403/404.
- [ ] Le déploiement utilise `svn export` (sans métadonnées) plutôt qu'un `svn checkout`/`svn update` direct dans le webroot.
- [ ] Une règle serveur bloque explicitement tous les dossiers/fichiers commençant par un point, pas uniquement `.svn`.
- [ ] Si une exposition a eu lieu, tous les secrets présents dans les révisions du dépôt ont été révoqués et régénérés.
