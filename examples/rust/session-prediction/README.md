# Session Prediction

Le handler `login` générait l'identifiant de session par un hash faible de l'ID utilisateur et de l'horodatage, deux valeurs prévisibles permettant à un attaquant de reconstruire un identifiant valide (CWE-330). La version corrigée génère l'identifiant exclusivement via `rand::thread_rng()` (CSPRNG, 256 bits d'entropie) et positionne le cookie de session en `HttpOnly`, `Secure` et `SameSite=Strict`.
