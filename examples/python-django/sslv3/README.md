# Protocole SSLv3 activé — Python/Django

`vulnerable.py` définit un `HTTPAdapter` `requests` personnalisé qui force `ssl.PROTOCOL_SSLv3` pour notifier un webhook partenaire d'événements de paiement (CWE-326, Inadequate Encryption Strength). SSLv3 est interdit par la RFC 7568 depuis 2015 et vulnérable à l'attaque POODLE, qui exploite un padding oracle sur le chiffrement CBC pour déchiffrer des données interceptées.

`fixed.py` remplace cet adaptateur par un `ssl.SSLContext` imposant `minimum_version = ssl.TLSVersion.TLSv1_2` (TLS 1.3 en plafond), rendant POODLE inapplicable, tout en conservant la vérification du certificat du webhook.

## Pourquoi c'est dangereux
- POODLE permet à un attaquant en position de man-in-the-middle de déchiffrer progressivement des données sensibles (ici, des informations de paiement) en forçant un repli sur SSLv3.
- SSLv3 ne propose aucune suite de chiffrement authentifiée (AEAD) moderne, contrairement à TLS 1.2+.
- Le protocole est explicitement proscrit par la RFC 7568 : aucun cas d'usage légitime ne justifie de le conserver, même pour de la "compatibilité" avec un partenaire ancien.

## Explication du correctif
- Remplacement de `ssl.PROTOCOL_SSLv3` par `ssl.PROTOCOL_TLS_CLIENT` avec `minimum_version` fixé explicitement à `TLSv1_2`.
- Ajout d'un plafond `maximum_version = TLSv1_3` pour privilégier la version la plus récente disponible.
- Conservation explicite de `check_hostname = True` et `verify_mode = ssl.CERT_REQUIRED`.

## Notes résiduelles
- Si le partenaire ne supporte réellement que SSLv3, la connexion doit être coupée et un plan de migration convenu avec lui — POODLE n'a pas de mitigation applicative fiable côté client.
- Vérifier côté serveur (si l'application expose aussi un point d'entrée TLS) que `TLS_FALLBACK_SCSV` est actif pour empêcher tout downgrade forcé.
