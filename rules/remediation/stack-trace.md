# Remédiation — Stack trace exposée au client

## Principe
Intercepter toutes les exceptions au niveau applicatif et ne jamais renvoyer la trace au client. Journaliser la trace complète côté serveur avec un identifiant de corrélation renvoyé au client.

## PHP (Laravel — handler d'exceptions)
```php
// Avant — vulnérable (app/Exceptions/Handler.php, APP_DEBUG=true en prod)
public function render($request, Throwable $exception)
{
    return parent::render($request, $exception); // affiche la trace complète si APP_DEBUG=true
}

// Après — sécurisé
public function render($request, Throwable $exception)
{
    $correlationId = Str::uuid()->toString();
    Log::error('Erreur applicative', [
        'correlation_id' => $correlationId,
        'exception' => $exception,
    ]);

    if ($request->expectsJson()) {
        return response()->json([
            'error' => 'Une erreur interne est survenue.',
            'correlation_id' => $correlationId,
        ], 500);
    }

    return response()->view('errors.500', ['correlationId' => $correlationId], 500);
}
```

## PHP (php.ini — production)
```ini
; Avant — vulnérable
display_errors = On

; Après — sécurisé
display_errors = Off
log_errors = On
error_log = /var/log/php/error.log
```

## Nginx (page d'erreur générique)
```nginx
# Avant — vulnérable
fastcgi_intercept_errors off;

# Après — sécurisé
fastcgi_intercept_errors on;
error_page 500 502 503 504 /50x.html;
location = /50x.html {
    internal;
}
```

## Checklist de vérification post-patch
- [ ] Aucune réponse HTTP (page ou JSON) ne contient de trace d'exécution, de chemin de fichier serveur ou de nom de classe interne.
- [ ] Toutes les exceptions non gérées sont interceptées par un handler global et journalisées côté serveur.
- [ ] Un identifiant de corrélation est renvoyé au client pour permettre le support sans exposer de détails techniques.
- [ ] `display_errors` est désactivé sur l'environnement de production.
