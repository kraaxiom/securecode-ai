# Remédiation — Configuration TLS faible (suites de chiffrement / vérification désactivée)

## Principe
Une configuration TLS permissive (suites de chiffrement faibles, vérification de certificat désactivée) expose à l'interception et à la falsification de trafic (MITM). Restreindre les suites aux algorithmes modernes (AEAD, forward secrecy) et ne jamais désactiver la vérification de certificat, même en développement.

## PHP (cURL)
```php
// Avant — vulnérable
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, false);
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, false);

// Après — sécurisé
curl_setopt($ch, CURLOPT_SSL_VERIFYPEER, true);
curl_setopt($ch, CURLOPT_SSL_VERIFYHOST, 2);
curl_setopt($ch, CURLOPT_SSLVERSION, CURL_SSLVERSION_TLSv1_2);
```

## Node.js (https/axios)
```js
// Avant — vulnérable
const agent = new https.Agent({ rejectUnauthorized: false });

// Après — sécurisé
const agent = new https.Agent({ rejectUnauthorized: true, minVersion: 'TLSv1.2' });
```

## Python (requests)
```python
# Avant — vulnérable
requests.get(url, verify=False)

# Après — sécurisé
requests.get(url, verify=True)  # utiliser le magasin de CA système ou un bundle explicite
```

## Checklist de vérification post-patch
- [ ] Aucune désactivation de la vérification de certificat (`verify=False`, `rejectUnauthorized: false`, `CURLOPT_SSL_VERIFYPEER=false`) ne subsiste, y compris en environnement de test.
- [ ] Les suites de chiffrement faibles (RC4, DES, suites sans forward secrecy) sont exclues de la configuration serveur.
- [ ] La version minimale de TLS acceptée est 1.2.
- [ ] Un scan externe (SSL Labs, `testssl.sh`) confirme une note A/A+ ou l'absence de faiblesses connues.
