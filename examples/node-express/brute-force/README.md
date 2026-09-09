# Brute Force — Node/Express

`vulnerable.js` expose une route `/login` sans aucune limitation du nombre de tentatives d'authentification, ni par compte ni par IP (CWE-307, Improper Restriction of Excessive Authentication Attempts). Un attaquant peut ainsi tester un très grand nombre de mots de passe sans être ralenti.

`fixed.js` ajoute un middleware `express-rate-limit` combinant IP et identifiant de compte comme clé, journalise les échecs pour permettre la détection, et utilise un message d'erreur générique afin de ne pas faciliter l'énumération de comptes.
