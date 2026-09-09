# Credential Stuffing — Node/Express

`vulnerable.js` authentifie chaque requête sans aucune détection comportementale globale : rien ne corrèle un volume élevé de tentatives sur des comptes différents provenant d'une même origine (CWE-307, Improper Restriction of Excessive Authentication Attempts). Cela permet à un attaquant de rejouer massivement des couples identifiant/mot de passe issus de fuites tierces.

`fixed.js` ajoute une détection de vélocité par IP (tous comptes confondus) qui déclenche une alerte de sécurité au-delà d'un seuil, et propose la MFA lorsque l'authentification réussit, seule protection réellement efficace contre la réutilisation de mots de passe volés.
