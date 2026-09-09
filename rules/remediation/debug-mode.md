# Remédiation — Mode debug activé en production

## Principe
Désactiver strictement tout mode debug/flag de développement en production, et automatiser la vérification de cette configuration dans le pipeline de déploiement.

## PHP (Laravel — .env)
```env
# Avant — vulnérable
APP_ENV=production
APP_DEBUG=true

# Après — sécurisé
APP_ENV=production
APP_DEBUG=false
```

## PHP (php.ini)
```ini
; Avant — vulnérable
display_errors = On
error_reporting = E_ALL

; Après — sécurisé
display_errors = Off
log_errors = On
error_reporting = E_ALL & ~E_DEPRECATED & ~E_STRICT
```

## Nginx (bloquer un endpoint de debug/profiler resté actif)
```nginx
# Avant — vulnérable
location /_profiler {
    proxy_pass http://app_backend;
}

# Après — sécurisé
location /_profiler {
    deny all;
    return 404;
}
```

## Apache (bloquer un endpoint de debug/profiler resté actif)
```apache
# Avant — vulnérable
Alias /_profiler /var/www/app/public/_profiler

# Après — sécurisé
<Location "/_profiler">
    Require all denied
</Location>
```

## Checklist de vérification post-patch
- [ ] `APP_DEBUG`/`DEBUG`/équivalent framework est à `false`/`False`/`Production` sur l'environnement de production.
- [ ] `display_errors` est à `Off` dans le `php.ini` de production.
- [ ] Aucune console/profiler de debug (Werkzeug, Symfony profiler, etc.) n'est accessible sans authentification.
- [ ] Une vérification automatisée en CI/CD échoue le déploiement si un flag debug est actif en environnement de production.
