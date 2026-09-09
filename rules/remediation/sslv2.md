# Remédiation — Protocole SSLv2 activé

## Principe
SSLv2 est totalement cassé (attaques DROWN, faiblesses de handshake) et ne doit jamais être activé. Configurer explicitement le serveur et les clients pour n'accepter que TLS 1.2 minimum, idéalement TLS 1.3.

## PHP (contexte stream / cURL)
```php
// Avant — vulnérable
$context = stream_context_create([
    'ssl' => ['crypto_method' => STREAM_CRYPTO_METHOD_SSLv2_CLIENT],
]);

// Après — sécurisé
$context = stream_context_create([
    'ssl' => [
        'crypto_method' => STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT | STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT,
        'verify_peer' => true,
        'verify_peer_name' => true,
    ],
]);
```

## Node.js (https/tls)
```js
// Avant — vulnérable
const server = tls.createServer({ secureProtocol: 'SSLv2_method', ...opts });

// Après — sécurisé
const server = tls.createServer({
  minVersion: 'TLSv1.2',
  maxVersion: 'TLSv1.3',
  ...opts,
});
```

## Nginx (config serveur)
```nginx
# Avant — vulnérable
ssl_protocols SSLv2 SSLv3 TLSv1;

# Après — sécurisé
ssl_protocols TLSv1.2 TLSv1.3;
```

## Checklist de vérification post-patch
- [ ] SSLv2 n'apparaît plus dans aucune configuration serveur, client ou code applicatif.
- [ ] Seuls TLS 1.2 et TLS 1.3 sont acceptés (`ssl_protocols`, `minVersion`/`maxVersion`, `crypto_method`).
- [ ] Un scan externe (ex: `testssl.sh`, Qualys SSL Labs) confirme l'absence de SSLv2/SSLv3 sur l'hôte.
- [ ] Les clients legacy dépendant de SSLv2 ont été identifiés et migrés avant la coupure.
