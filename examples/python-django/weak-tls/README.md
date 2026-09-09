# Configuration TLS faible — Python/Django

`vulnerable.py` appelle une API de paiement tierce avec `requests.post(..., verify=False)`, désactivant totalement la vérification du certificat serveur, et construit un contexte SSL personnalisé autorisant des suites de chiffrement obsolètes (RC4, DES) sans forward secrecy (CWE-326, Inadequate Encryption Strength).

`fixed.py` conserve la vérification de certificat activée (`verify=True`, comportement par défaut de `requests`) et remplace le contexte SSL personnalisé par une configuration restreinte aux suites AEAD modernes (AES-GCM, ChaCha20-Poly1305) avec forward secrecy, TLS 1.2 minimum, et vérification stricte du hostname/certificat.

## Pourquoi c'est dangereux
- `verify=False` rend l'application aveugle à toute substitution de certificat : un attaquant en position de MITM (réseau Wi-Fi compromis, proxy malveillant, DNS spoofing) peut intercepter ou modifier les échanges, y compris des données de paiement.
- Des suites de chiffrement obsolètes (RC4, DES, CBC sans MAC-then-Encrypt correct) sont cassables ou affaiblies par des attaques connues, même sous TLS.
- L'absence de forward secrecy signifie qu'une compromission future de la clé privée du serveur permettrait de déchiffrer du trafic passé enregistré.

## Explication du correctif
- Suppression de `verify=False` : la vérification du certificat via le magasin de CA système reste active par défaut.
- Le contexte SSL personnalisé impose `check_hostname=True`, `verify_mode=ssl.CERT_REQUIRED` et `minimum_version=ssl.TLSVersion.TLSv1_2`.
- Les suites de chiffrement sont restreintes à `ECDHE+AESGCM:ECDHE+CHACHA20`, exclusivement des algorithmes AEAD avec échange de clé éphémère.

## Notes résiduelles
- Si un certificat interne (CA privée) est utilisé, préférer `verify="/path/to/ca-bundle.pem"` plutôt que de désactiver la vérification.
- Auditer périodiquement la configuration TLS exposée (SSL Labs, `testssl.sh`) pour détecter toute régression future.
