# Remédiation — Messages d'erreur verbeux exposés

## Principe
Mettre en place un gestionnaire d'erreurs centralisé qui renvoie des messages génériques au client, et ne jamais propager directement une exception de bibliothèque/DB dans la réponse HTTP.

## PHP (PDO — capture d'erreur DB)
```php
// Avant — vulnérable
try {
    $stmt = $pdo->query("SELECT * FROM users WHERE id = $id");
} catch (PDOException $e) {
    echo "Erreur SQL : " . $e->getMessage(); // ex: fuite du schéma, du driver, de la requête
}

// Après — sécurisé
try {
    $stmt = $pdo->prepare("SELECT * FROM users WHERE id = :id");
    $stmt->execute(['id' => (int) $id]);
} catch (PDOException $e) {
    error_log('Erreur SQL: ' . $e->getMessage());
    http_response_code(500);
    echo json_encode(['error' => 'Une erreur interne est survenue.']);
}
```

## PHP (Laravel — config production)
```php
// Avant — vulnérable (config/app.php)
'debug' => true,

// Après — sécurisé
'debug' => env('APP_DEBUG', false),
```

## Nginx (masquer les pages d'erreur du serveur upstream)
```nginx
# Avant — vulnérable
proxy_intercept_errors off;

# Après — sécurisé
proxy_intercept_errors on;
error_page 500 502 503 504 /50x.html;
location = /50x.html {
    internal;
}
```

## Checklist de vérification post-patch
- [ ] Aucune réponse d'erreur (4xx/5xx) ne contient de message technique brut (driver DB, requête SQL, chemin de fichier).
- [ ] Un gestionnaire d'erreurs global/middleware centralise la mise en forme de toutes les réponses d'erreur.
- [ ] Les messages d'exception de bibliothèques tierces sont capturés et journalisés côté serveur, jamais renvoyés tels quels.
- [ ] Des codes d'erreur applicatifs stables et documentés remplacent les messages techniques bruts.
