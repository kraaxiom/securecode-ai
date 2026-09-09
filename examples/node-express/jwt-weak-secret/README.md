# JWT secret faible — Node/Express

`vulnerable.js` signe les tokens avec `your-256-bit-secret`, la valeur d'exemple issue de la documentation de `jsonwebtoken`, jamais remplacée (CWE-326, Inadequate Encryption Strength). Un secret aussi prévisible peut être retrouvé hors ligne, permettant de forger des tokens avec n'importe quel rôle.

`fixed.js` charge le secret depuis `process.env.JWT_SECRET`, généré une fois via un générateur cryptographique et stocké dans un gestionnaire de secrets, et vérifie explicitement au démarrage qu'il fait au moins 256 bits (32 octets) d'entropie.
