# Bucket S3 public — Python/Django

`vulnerable.py` téléverse les factures utilisateur dans un bucket S3 avec une ACL `public-read` et désactive le Block Public Access lors de la création du bucket (CWE-284, Improper Access Control). Toute personne disposant de l'URL de l'objet — ou capable de lister le bucket — peut accéder à des documents sensibles sans authentification.

`fixed.py` conserve le bucket entièrement privé (Block Public Access activé sur les 4 paramètres, chiffrement SSE-KMS) et remplace l'accès public permanent par une URL pré-signée à durée de vie courte (15 minutes), générée uniquement après vérification explicite que l'utilisateur demandeur est bien le propriétaire de la ressource.

## Pourquoi c'est dangereux
- Un bucket public expose potentiellement des données personnelles, financières ou d'identification à quiconque devine ou obtient une URL.
- Sans contrôle d'autorisation applicatif, l'accès n'est plus lié à l'identité de l'utilisateur mais uniquement à la connaissance d'une URL.
- Les scanners automatisés (Shodan, GrayhatWarfare) recensent activement les buckets publics mal configurés.

## Explication du correctif
- Suppression de toute ACL publique lors de l'upload (`upload_fileobj` sans `ACL: public-read`).
- Activation intégrale du Block Public Access au niveau du bucket dès sa création.
- Remplacement de l'URL publique permanente par une URL pré-signée (`generate_presigned_url`) avec expiration courte.
- Ajout d'un contrôle d'autorisation explicite (vérification du préfixe `invoices/{user.id}/`) avant toute génération d'URL.

## Notes résiduelles
- Les URLs pré-signées restent valides jusqu'à expiration même si partagées par erreur ; garder la durée de vie la plus courte possible selon l'usage.
- Prévoir une revue périodique (AWS Config / Access Analyzer) pour détecter toute dérive de configuration future.
