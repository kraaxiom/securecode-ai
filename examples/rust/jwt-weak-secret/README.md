# JWT Weak Secret

Le handler `issue_token` signait les tokens avec un secret HMAC court et codé en dur (`"secret"`), permettant une attaque hors ligne par dictionnaire pour forger des tokens arbitraires (CWE-326). La version corrigée charge le secret depuis une variable d'environnement (représentant un gestionnaire de secrets), vérifie explicitement une longueur minimale de 256 bits au démarrage, et ne conserve aucune valeur en dur dans le code source.
