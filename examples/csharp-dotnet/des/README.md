# Utilisation de DES / 3DES en C#/.NET

**CWE-327 — Use of a Broken or Risky Cryptographic Algorithm**

## Vulnérabilité

`DESCryptoServiceProvider`/`DES.Create()` et `TripleDESCryptoServiceProvider`/`TripleDES.Create()` reposent sur des algorithmes obsolètes :

- **DES** utilise une clé effective de 56 bits, cassable par force brute en quelques heures avec du matériel moderne (FPGA, GPU en cluster).
- **3DES (Triple DES)** utilise un bloc de 64 bits, vulnérable à l'attaque **Sweet32** : après environ 2^32 blocs chiffrés avec la même clé, des collisions de bloc permettent de récupérer des fragments de texte en clair. Le NIST a retiré 3DES de ses recommandations en 2023.
- L'exemple `Vulnerable.cs` utilise en plus le mode CBC sans authentification (pas de MAC/AEAD), ce qui expose potentiellement à des attaques de type padding oracle si le texte chiffré est accessible et modifiable par un attaquant.

Ce fichier ne contient aucun code d'exploitation : il illustre uniquement la configuration à risque (choix d'algorithme et de mode) à des fins de détection et de formation.

## Correction

`Fixed.cs` remplace DES/3DES par **AES-256 en mode GCM** (`System.Security.Cryptography.AesGcm`), une construction AEAD (Authenticated Encryption with Associated Data) :

- Clé de 256 bits, générée via `RandomNumberGenerator.GetBytes` (CSPRNG), à stocker dans un gestionnaire de secrets (Azure Key Vault, AWS KMS, HashiCorp Vault) plutôt que codée en dur.
- Le mode GCM apporte confidentialité **et** intégrité : toute altération du texte chiffré fait échouer le déchiffrement au lieu de produire silencieusement des données corrompues.
- Un nonce (IV) de 12 octets est généré aléatoirement à chaque chiffrement et n'est jamais réutilisé avec la même clé, condition indispensable à la sécurité de GCM.
- Les données déjà chiffrées en DES/3DES doivent faire l'objet d'un plan de re-chiffrement/migration lors du déploiement du correctif.

## Références

- CWE-327: Use of a Broken or Risky Cryptographic Algorithm
- OWASP Cryptographic Storage Cheat Sheet
- NIST SP 800-131A Rev. 2 (retrait de TDEA/3DES)
- `rules/remediation/des.md`, `knowledge/crypto/des.md`
