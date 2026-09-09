# Remédiation — Utilisation de DES / 3DES

## Principe
Remplacer tout chiffrement symétrique basé sur DES ou 3DES (taille de bloc et de clé trop faibles, vulnérable aux attaques par force brute et par collision de bloc de type Sweet32) par AES-256 en mode authentifié (GCM). Ne jamais réutiliser le même nonce/IV pour deux chiffrements.

## PHP (openssl)
```php
// Avant — vulnérable
$cipher = openssl_encrypt($data, 'des-ede3-cbc', $key, 0, $iv);

// Après — sécurisé
$iv = random_bytes(12);
$tag = '';
$cipher = openssl_encrypt($data, 'aes-256-gcm', $key, OPENSSL_RAW_DATA, $iv, $tag);
// Stocker $iv et $tag avec le texte chiffré, nécessaires au déchiffrement
```

## Node.js (crypto)
```js
// Avant — vulnérable
const cipher = crypto.createCipheriv('des-ede3-cbc', key, iv);

// Après — sécurisé
const iv = crypto.randomBytes(12);
const cipher = crypto.createCipheriv('aes-256-gcm', key, iv);
let encrypted = cipher.update(data, 'utf8', 'hex');
encrypted += cipher.final('hex');
const authTag = cipher.getAuthTag();
```

## Python (cryptography)
```python
# Avant — vulnérable
from Crypto.Cipher import DES3
cipher = DES3.new(key, DES3.MODE_CBC, iv)

# Après — sécurisé
from cryptography.hazmat.primitives.ciphers.aead import AESGCM
aesgcm = AESGCM(key)  # clé de 32 octets (AES-256)
nonce = os.urandom(12)
ciphertext = aesgcm.encrypt(nonce, data, associated_data=None)
```

## Checklist de vérification post-patch
- [ ] Plus aucun appel à DES, 3DES, DES-EDE ou Blowfish dans le code corrigé.
- [ ] Le nouvel algorithme est AES-256 (idéalement en mode GCM authentifié).
- [ ] Un nonce/IV unique et aléatoire est généré à chaque chiffrement, jamais réutilisé.
- [ ] Les données déjà chiffrées avec l'ancien algorithme sont re-chiffrées ou migrées (plan de rotation des clés).
- [ ] Un test de non-régression confirme que le chiffrement/déchiffrement fonctionne toujours de bout en bout.
