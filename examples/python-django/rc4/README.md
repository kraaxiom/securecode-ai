# Chiffrement RC4 cassé — Python/Django

`vulnerable.py` chiffre le contenu de pièces jointes sensibles avec RC4 (`Crypto.Cipher.ARC4`) en utilisant une clé statique codée en dur, réutilisée pour tous les fichiers (CWE-327, Use of a Broken or Risky Cryptographic Algorithm). RC4 présente des biais statistiques connus dans son flux de sortie et est interdit dans TLS depuis la RFC 7465 (2015).

`fixed.py` remplace RC4 par AES-256-GCM, chiffrement authentifié sans biais statistique connu. La clé est chargée depuis la configuration Django (gestionnaire de secrets), et un nonce aléatoire unique est généré à chaque chiffrement de fichier.

## Pourquoi c'est dangereux
- Les biais statistiques du flux de sortie RC4 permettent de récupérer partiellement le texte en clair, en particulier sur les premiers octets du flux (attaques démontrées sur WEP et sur les cookies chiffrés en TLS-RC4).
- L'absence de nonce/IV et la réutilisation d'une clé statique pour tous les fichiers font que deux fichiers chiffrés avec la même clé produisent des flux corrélés, exploitables par analyse croisée.
- Aucun mécanisme d'intégrité : un attaquant peut altérer le fichier chiffré sans que l'application le détecte au déchiffrement.

## Explication du correctif
- Remplacement de `ARC4.new(...)` par `AESGCM`, chiffrement authentifié avec clé de 256 bits.
- Génération d'un nonce aléatoire de 12 octets (`os.urandom(12)`) à chaque chiffrement de fichier, stocké avec le texte chiffré.
- Vérification automatique de l'intégrité et de l'authenticité par AES-GCM lors du déchiffrement, avec gestion explicite de l'échec (fichier corrompu ou falsifié).
- Clé chargée depuis `settings.ENCRYPTION_KEY`, elle-même issue d'une variable d'environnement ou d'un coffre-fort de secrets.

## Notes résiduelles
- Prévoir un plan de migration pour re-chiffrer les pièces jointes déjà stockées en RC4.
- S'assurer également que les suites de chiffrement TLS du serveur web excluent explicitement RC4, en complément de ce correctif applicatif.
