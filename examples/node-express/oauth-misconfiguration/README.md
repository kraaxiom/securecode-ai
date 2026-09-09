# Mauvaise configuration OAuth — Node/Express

`vulnerable.js` initie le flux d'autorisation sans générer de paramètre `state`, et traite le callback sans le vérifier ni valider strictement l'ID token reçu (CWE-287, Improper Authentication). Cela expose le callback à une CSRF permettant de lier une session victime à un compte contrôlé par l'attaquant.

`fixed.js` génère un `state` et un `nonce` uniques stockés en session avant la redirection, puis les transmet à `client.callback()` qui valide intégralement l'ID token (signature, émetteur, audience, expiration) avant d'en exploiter les claims.
