# Chiffrement DES/3DES obsolète — Python/Django

`vulnerable.py` chiffre le numéro de sécurité sociale d'un utilisateur avec Triple DES (3DES) en mode CBC, en utilisant une clé et un IV codés en dur et statiques (CWE-327, Use of a Broken or Risky Cryptographic Algorithm). Le bloc de 64 bits de 3DES le rend vulnérable à l'attaque Sweet32, et l'algorithme est retiré des standards NIST depuis 2023.

`fixed.py` remplace 3DES par AES-256-GCM, un chiffrement authentifié moderne à bloc de 128 bits. La clé est chargée depuis la configuration Django (issue d'un gestionnaire de secrets), et un nonce aléatoire unique est généré à chaque chiffrement.

## Pourquoi c'est dangereux
- La taille de bloc de 64 bits de 3DES permet l'attaque Sweet32, qui devient praticable dès que plusieurs gigaoctets de données sont chiffrés avec la même clé.
- Une clé et un IV statiques réutilisés pour tous les utilisateurs suppriment toute variabilité cryptographique, facilitant l'analyse de motifs.
- Le mode CBC sans authentification (pas de HMAC/GCM) permet à un attaquant de modifier le texte chiffré sans que l'application le détecte.

## Explication du correctif
- Remplacement de `DES3.new(...)` par `AESGCM`, chiffrement authentifié avec clé de 256 bits.
- Génération d'un nonce aléatoire de 12 octets (`os.urandom(12)`) à chaque chiffrement, stocké avec le texte chiffré (non secret par nature).
- Vérification automatique de l'intégrité et de l'authenticité par AES-GCM lors du déchiffrement, avec gestion explicite de l'échec.
- Clé chargée depuis `settings.ENCRYPTION_KEY`, elle-même issue d'une variable d'environnement ou d'un coffre-fort de secrets.

## Notes résiduelles
- Prévoir un plan de migration pour re-chiffrer les données existantes encore protégées par 3DES.
- La clé AES doit faire l'objet d'une rotation périodique et d'un stockage dans un gestionnaire de secrets dédié (Vault, AWS/Azure/GCP Secret Manager), pas seulement dans une variable d'environnement en clair.
