# Protocole TLS 1.0 (et 1.1) activé — Python/Django

`vulnerable.py` définit un `HTTPAdapter` `requests` personnalisé qui fige à la fois `minimum_version` et `maximum_version` sur `ssl.TLSVersion.TLSv1` pour appeler une passerelle de paiement partenaire (CWE-326, Inadequate Encryption Strength). TLS 1.0/1.1 sont officiellement dépréciés par la RFC 8996 (2021), vulnérables à l'attaque BEAST, dépourvus de suites de chiffrement authentifiées modernes, et explicitement exclus des référentiels PCI-DSS depuis 2018 pour tout flux traitant des données de carte.

`fixed.py` relève `minimum_version` à `ssl.TLSVersion.TLSv1_2` (avec TLS 1.3 en plafond), tout en conservant la vérification du certificat de la passerelle.

## Pourquoi c'est dangereux
- BEAST exploite une faiblesse du chaînage de blocs en mode CBC dans TLS 1.0, permettant à un attaquant de récupérer partiellement le texte en clair de la session.
- TLS 1.0/1.1 ne supportent pas les suites de chiffrement AEAD (ex: AES-GCM), qui sont la norme en TLS 1.2+.
- Un flux de paiement limité à TLS 1.0 est non conforme PCI-DSS, exposant l'entreprise à des sanctions et à la perte de son habilitation à traiter des paiements par carte.

## Explication du correctif
- Remplacement des bornes `minimum_version`/`maximum_version` figées sur `TLSv1` par `minimum_version = TLSv1_2` et `maximum_version = TLSv1_3`.
- Conservation explicite de `check_hostname = True` et `verify_mode = ssl.CERT_REQUIRED` pour garantir l'authenticité du serveur distant.

## Notes résiduelles
- Si la passerelle de paiement partenaire ne supporte que TLS 1.0, elle doit être considérée comme non conforme PCI-DSS et remplacée ou mise à jour avant toute mise en production.
- Documenter et planifier la migration de toute intégration tierce encore limitée à TLS 1.0/1.1 identifiée lors de l'audit.
