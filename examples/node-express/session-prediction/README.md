# Session Prediction — Node/Express

`vulnerable.js` génère les identifiants de session par concaténation d'un ID utilisateur, d'un timestamp et d'un compteur (CWE-330, Use of Insufficiently Random Values), des valeurs qu'un attaquant peut deviner ou reconstruire pour usurper une session valide.

`fixed.js` délègue la génération à `express-session` avec `genid` basé sur `crypto.randomBytes(32)` (générateur cryptographique, 256 bits d'entropie), régénère l'ID à chaque connexion, et marque le cookie `HttpOnly`, `Secure` et `SameSite`.
