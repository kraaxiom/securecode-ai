# Remédiation — Page phpinfo() exposée

## Principe
Supprimer tout fichier de test appelant `phpinfo()` avant le déploiement en production. Si un diagnostic serveur est nécessaire, le protéger par authentification stricte et restriction IP.

## PHP
```php
// Avant — vulnérable (info.php accessible publiquement)
<?php
phpinfo();

// Après — sécurisé (fichier supprimé, ou si diagnostic requis, protégé)
<?php
if (!isset($_SERVER['PHP_AUTH_USER']) || $_SERVER['PHP_AUTH_USER'] !== $_ENV['DIAG_USER']) {
    header('WWW-Authenticate: Basic realm="Diagnostic"');
    header('HTTP/1.0 401 Unauthorized');
    exit('Accès refusé');
}
phpinfo();
```

## Nginx (bloquer un fichier phpinfo oublié, en complément de sa suppression)
```nginx
# Avant — vulnérable
location ~ \.php$ {
    fastcgi_pass unix:/run/php/php-fpm.sock;
}

# Après — sécurisé
location ~ /(phpinfo|info|test)\.php$ {
    deny all;
    return 404;
}

location ~ \.php$ {
    fastcgi_pass unix:/run/php/php-fpm.sock;
}
```

## Apache
```apache
# Avant — vulnérable
DocumentRoot /var/www/html

# Après — sécurisé
<FilesMatch "^(phpinfo|info|test)\.php$">
    Require all denied
</FilesMatch>
```

## Checklist de vérification post-patch
- [ ] Aucun fichier appelant `phpinfo()` sans authentification n'est présent dans le webroot de production.
- [ ] Une requête sur les noms de fichiers courants (`phpinfo.php`, `info.php`, `test.php`) renvoie 403/404.
- [ ] Si un endpoint de diagnostic reste nécessaire, il est protégé par authentification et restriction IP.
- [ ] Un contrôle en CI/CD interdit la présence de fichiers de debug connus dans l'artefact de déploiement.
