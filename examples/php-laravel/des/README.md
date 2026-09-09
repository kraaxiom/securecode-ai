## Vulnérabilité

Utilisation de DES / 3DES (CWE-327 — Use of a Broken or Risky Cryptographic Algorithm) pour chiffrer un export de bulletins de paie dans un service Laravel.

## Impact

DES utilise une clé effective de 56 bits, cassable par force brute en quelques heures avec du matériel moderne. Le 3DES (DES-EDE3), malgré une clé nominale plus longue, opère sur des blocs de 64 bits et reste vulnérable à l'attaque Sweet32, qui permet de récupérer du texte en clair après observation d'un volume suffisant de trafic chiffré avec la même clé. Des données sensibles (paie, identifiants, dossiers RH) chiffrées avec ces algorithmes peuvent être compromises par un attaquant disposant d'un accès réseau ou du texte chiffré archivé.

## Cause racine

Choix d'un algorithme de chiffrement symétrique historique (`des-ede3-cbc`) encore présent dans OpenSSL pour compatibilité descendante, combiné à un IV dérivé de manière prévisible (`md5(uniqid())`) au lieu d'un générateur cryptographiquement sûr.

## Correction

- Remplacer DES/3DES par AES-256 en mode authentifié (AES-GCM).
- Générer un IV/nonce aléatoire unique de 12 octets via `random_bytes()` à chaque chiffrement, jamais réutilisé.
- Charger la clé de chiffrement (32 octets) depuis une variable d'environnement ou un gestionnaire de secrets, jamais codée en dur.
- Stocker le tag d'authentification GCM avec le texte chiffré et vérifier son intégrité au déchiffrement.
- Planifier la migration (re-chiffrement) des données déjà protégées par DES/3DES.

## Références

- CWE-327: Use of a Broken or Risky Cryptographic Algorithm
- OWASP A02:2021 - Cryptographic Failures
- NIST SP 800-131A Rev. 2 (retrait de TDEA/3DES)
