# Politique de mot de passe faible — Node/Express

`vulnerable.js` n'impose qu'une longueur minimale de 6 caractères et n'effectue aucune vérification contre les mots de passe déjà compromis (CWE-521, Weak Password Requirements). Cela facilite le brute force et le credential stuffing sur les comptes créés.

`fixed.js` porte la longueur minimale à 12 caractères et vérifie chaque mot de passe via le service `hibp` (HaveIBeenPwned, k-anonymity) avant de l'accepter, sans imposer de règles de composition artificielles, conformément aux recommandations NIST 800-63B.
