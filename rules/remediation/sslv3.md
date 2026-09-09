# Remédiation — Protocole SSLv3 activé

## Principe
SSLv3 est vulnérable à l'attaque POODLE permettant le déchiffrement de données via un downgrade de protocole. Il doit être désactivé partout, au profit de TLS 1.2 minimum.

## PHP (contexte stream / cURL)
```php
// Avant — vulnérable
$context = stream_context_create([
    'ssl' => ['crypto_method' => STREAM_CRYPTO_METHOD_SSLv3_CLIENT],
]);

// Après — sécurisé
$context = stream_context_create([
    'ssl' => [
        'crypto_method' => STREAM_CRYPTO_METHOD_TLSv1_2_CLIENT | STREAM_CRYPTO_METHOD_TLSv1_3_CLIENT,
        'verify_peer' => true,
    ],
]);
```

## Node.js (https/tls)
```js
// Avant — vulnérable
https.request({ secureProtocol: 'SSLv3_method', ...opts });

// Après — sécurisé
https.request({ minVersion: 'TLSv1.2', maxVersion: 'TLSv1.3', ...opts });
```

## Nginx (config serveur)
```nginx
# Avant — vulnérable
ssl_protocols SSLv3 TLSv1;

# Après — sécurisé
ssl_protocols TLSv1.2 TLSv1.3;
```

## Checklist de vérification post-patch
- [ ] SSLv3 n'apparaît plus dans aucune configuration serveur, client ou code applicatif.
- [ ] Seuls TLS 1.2 et TLS 1.3 sont acceptés.
- [ ] Un scan externe (`testssl.sh`, SSL Labs) confirme que POODLE n'est plus applicable sur l'hôte.
- [ ] Les intégrations tierces encore dépendantes de SSLv3 ont été recensées et migrées.
