# JWT Weak Secret

La version corrigée remplace le secret court et codé en dur (`secret123`) par un secret à haute entropie (généré via `random_bytes(32)`), stocké dans une variable d'environnement/gestionnaire de secrets, et ajoute une expiration courte au token. Cela empêche un attaquant de retrouver le secret par force brute hors ligne pour forger des tokens. La faille correspond à CWE-326 (Inadequate Encryption Strength), comme indiqué dans `knowledge/auth/jwt-weak-secret.md`.
