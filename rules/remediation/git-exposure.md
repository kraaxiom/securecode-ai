# Remédiation — Répertoire .git exposé

## Principe
Ne jamais déployer en clonant directement le dépôt dans le webroot ; utiliser un artefact de build séparé du dossier `.git`. Bloquer en défense en profondeur l'accès aux dotdirs au niveau du serveur web.

## Nginx
```nginx
# Avant — vulnérable
server {
    listen 80;
    root /var/www/html; # contient .git/
    index index.php;
}

# Après — sécurisé
server {
    listen 80;
    root /var/www/html;
    index index.php;

    location ~ /\.git {
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
<DirectoryMatch "^/var/www/html/\.git">
    Require all denied
</DirectoryMatch>

RedirectMatch 404 "(^|/)\.git"
```

## Checklist de vérification post-patch
- [ ] Une requête sur `/.git/HEAD` et `/.git/config` renvoie 403/404.
- [ ] Le processus de déploiement ne copie plus le dossier `.git` dans le webroot (déploiement par artefact/build, pas par `git pull` direct).
- [ ] Une règle serveur bloque explicitement tous les dossiers/fichiers commençant par un point, pas uniquement `.git`.
- [ ] Si une exposition a eu lieu, tous les secrets présents dans l'historique Git (y compris commits supprimés) ont été révoqués et régénérés.
