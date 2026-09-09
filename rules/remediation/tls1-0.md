# Remédiation — Protocole TLS 1.0 activé

## Principe
TLS 1.0 est obsolète (vulnérable à BEAST, ne supporte pas les suites de chiffrement modernes) et son usage est proscrit par PCI-DSS depuis 2018. Configurer TLS 1.2 comme version minimale, TLS 1.3 recommandée.

## PHP (contexte stream / cURL)
```php
// Avant — vulnérable
$context = stream_context_create([
    'ssl' => ['crypto_method' => STREAM_CRYPTO_METHOD_TLSv1_0_CLIENT],
]);

// Après — sécurisé
$context = stream_context_create([
    'ssl' => ['crypto_method' => STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT | STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT],
]);
```

## Node.js (https/tls)
```js
// Avant — vulnérable
const agent = new https.Agent({ secureProtocol: 'TLSv1_method' });

// Après — sécurisé
const agent = new https.Agent({ minVersion: 'TLSv1.2', maxVersion: 'TLSv1.3' });
```

## Nginx (config serveur)
```nginx
# Avant — vulnérable
ssl_protocols TLSv1 TLSv1.1;

# Après — sécurisé
ssl_protocols TLSv1.2 TLSv1.3;
```

## Checklist de vérification post-patch
- [ ] TLS 1.0 et 1.1 n'apparaissent plus dans aucune configuration serveur, client ou code applicatif.
- [ ] Seuls TLS 1.2 et TLS 1.3 sont acceptés.
- [ ] Un scan externe confirme la conformité PCI-DSS (absence de TLS < 1.2).
- [ ] Les clients/API tierces encore limités à TLS 1.0 ont été identifiés avant la coupure.
