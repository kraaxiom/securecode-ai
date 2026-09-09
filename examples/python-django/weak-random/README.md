# Générateur aléatoire non cryptographique — Python/Django

`vulnerable.py` génère des tokens de réinitialisation de mot de passe, des clés d'API et des codes OTP avec le module standard `random` (`random.choices`, `random.random`, `random.randint`), un PRNG (Mersenne Twister) statistique et prévisible, inadapté à un usage sécuritaire (CWE-338, Use of Cryptographically Weak PRNG).

`fixed.py` remplace tous ces usages par le module `secrets`, un CSPRNG adossé à `os.urandom`, avec une entropie d'au moins 128 bits pour chaque token ou clé générée.

## Pourquoi c'est dangereux
- Le Mersenne Twister utilisé par `random` n'est pas conçu pour résister à un attaquant : observer un nombre suffisant de sorties permet de reconstruire son état interne et de prédire les valeurs passées et futures.
- Un token de réinitialisation de mot de passe prévisible permet à un attaquant de prendre le contrôle d'un compte sans connaître le mot de passe.
- Une clé d'API ou un code OTP prévisible compromet respectivement l'authentification programmatique et la vérification en deux étapes.

## Explication du correctif
- `secrets.token_urlsafe(32)` remplace la génération du token de réinitialisation, offrant une entropie cryptographiquement sûre.
- `secrets.token_hex(32)` remplace la génération de la clé d'API.
- `secrets.choice(...)` remplace `random.randint` pour tirer chaque chiffre du code OTP depuis un CSPRNG.

## Notes résiduelles
- Un code OTP à 6 chiffres reste intrinsèquement de faible entropie côté format (10^6 combinaisons) ; il doit toujours être combiné à une limitation du nombre de tentatives et une expiration courte côté serveur, quel que soit le générateur utilisé.
- Invalider tout token déjà émis avec l'ancien générateur `random` si un risque de compromission est suspecté.
