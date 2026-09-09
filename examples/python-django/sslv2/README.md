# Protocole SSLv2 activé — Python/Django

`vulnerable.py` définit un `HTTPAdapter` `requests` personnalisé qui force la négociation TLS en `ssl.PROTOCOL_SSLv2` pour communiquer avec une API partenaire legacy (CWE-326, Inadequate Encryption Strength). SSLv2 est interdit par la RFC 6176 depuis 2011 : le handshake n'est pas authentifié, le MAC est faible, et le protocole est exploitable via l'attaque DROWN pour déchiffrer du trafic TLS partageant la même clé RSA.

`fixed.py` remplace cet adaptateur par un `ssl.SSLContext` imposant `minimum_version = ssl.TLSVersion.TLSv1_2` (avec TLS 1.3 en plafond), tout en conservant la vérification du certificat et du nom d'hôte du serveur.

## Pourquoi c'est dangereux
- SSLv2 permet à un attaquant en position de man-in-the-middle d'intercepter et de déchiffrer les échanges, y compris des données métier sensibles (factures, identifiants).
- L'attaque DROWN exploite des serveurs SSLv2 pour casser des sessions TLS modernes partageant la même clé RSA, même sur des connexions qui n'utilisent pas directement SSLv2.
- Aucun client ou serveur moderne ne devrait accepter de négocier ce protocole, considéré comme cassé depuis plus de 20 ans.

## Explication du correctif
- Remplacement de `ssl.PROTOCOL_SSLv2` par `ssl.PROTOCOL_TLS_CLIENT` avec `minimum_version` fixé explicitement à `TLSv1_2`.
- Ajout d'un plafond `maximum_version = TLSv1_3` pour privilégier la version la plus récente disponible.
- Conservation explicite de `check_hostname = True` et `verify_mode = ssl.CERT_REQUIRED` pour empêcher toute interception malgré le protocole renforcé.

## Notes résiduelles
- Si le partenaire legacy ne supporte réellement que SSLv2, la connexion doit être refusée et un plan de migration (ou un intermédiaire de chiffrement conforme) doit être mis en place — aucun compromis n'est acceptable sur ce protocole.
- Auditer périodiquement la configuration TLS sortante avec un outil comme `testssl.sh` pour s'assurer qu'aucune régression ne réintroduit un protocole obsolète.
