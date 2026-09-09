# Conteneur Azure Blob public — Python/Django

`vulnerable.py` crée un conteneur Azure Blob Storage avec `public_access=PublicAccess.CONTAINER`, autorisant la lecture et le listing anonymes de tous les blobs (CWE-284, Improper Access Control). Les documents utilisateur téléversés deviennent accessibles à quiconque connaît ou devine l'URL du conteneur.

`fixed.py` crée le conteneur sans accès public (privé par défaut) et remplace l'accès direct par un jeton SAS (`generate_blob_sas`) en lecture seule, expirant après une heure, généré uniquement après vérification que le demandeur est bien le propriétaire du document.

## Pourquoi c'est dangereux
- L'accès public au niveau "container" permet non seulement de lire un blob connu, mais aussi de lister tous les blobs du conteneur, exposant l'intégralité des documents stockés.
- Aucune trace d'authentification n'est nécessaire, rendant l'exfiltration difficile à détecter.

## Explication du correctif
- Création du conteneur sans paramètre `public_access` (privé par défaut).
- Génération d'un SAS scoped (`BlobSasPermissions(read=True)`) avec expiration courte au lieu d'un accès permanent.
- Contrôle d'autorisation explicite (préfixe `documents/{user.id}/`) avant toute émission de SAS.
- La clé de compte de stockage est chargée depuis les settings/variables d'environnement, jamais codée en dur.

## Notes résiduelles
- Restreindre également le firewall réseau du compte de stockage (pas de "All networks" ouvert) en complément du contrôle applicatif.
- Auditer périodiquement via Azure Defender for Storage / Azure Policy pour détecter toute réactivation accidentelle de l'accès public.
