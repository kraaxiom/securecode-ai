# Remédiation — Utilisation du chiffrement RC4

## Principe
RC4 est un algorithme de flux cassé (biais statistiques exploitables, attaques pratiques contre TLS/WEP). Le remplacer par AES-256-GCM pour le chiffrement applicatif, et désactiver RC4 dans la configuration des suites de chiffrement TLS du serveur.

## PHP
```php
// Avant — vulnérable
$cipher = openssl_encrypt($data, 'rc4', $key, OPENSSL_RAW_DATA);

// Après — sécurisé
$iv = random_bytes(12);
$tag = '';
$cipher = openssl_encrypt($data, 'aes-256-gcm', $key, OPENSSL_RAW_DATA, $iv, $tag);
```

## Node.js
```js
// Avant — vulnérable
const cipher = crypto.createCipheriv('rc4', key, '');

// Après — sécurisé
const iv = crypto.randomBytes(12);
const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
```

## Python
```python
# Avant — vulnérable
from Crypto.Cipher import ARC4
cipher = ARC4.new(key)

# Après — sécurisé
from cryptography.hazmat.primitives.ciphers.aead import AESGCM
aesgcm = AESGCM(key)
nonce = os.urandom(12)
ciphertext = aesgcm.encrypt(nonce, data, None)
```

## Checklist de vérification post-patch
- [ ] Plus aucun usage de RC4 (ou ARC4) dans le code applicatif corrigé.
- [ ] La configuration TLS du serveur (nginx/Apache/HAProxy) exclut explicitement les suites RC4.
- [ ] Le nouvel algorithme (AES-256-GCM) est utilisé avec un nonce unique par chiffrement.
- [ ] Un test de non-régression confirme le chiffrement/déchiffrement bout en bout.
