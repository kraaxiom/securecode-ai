# Bucket S3 public

La version vulnérable téléverse chaque document avec l'ACL `public-read` codée en dur, rendant l'objet accessible à quiconque connaît son URL directe sans aucune authentification, ce qui expose des documents potentiellement sensibles à un accès public non contrôlé (CWE-284, Improper Access Control). La version corrigée stocke les objets avec l'ACL `private` et ne partage l'accès que via une URL pré-signée générée à la demande avec une expiration courte (15 minutes), garantissant que seul un utilisateur autorisé disposant du lien temporaire peut consulter le fichier.
