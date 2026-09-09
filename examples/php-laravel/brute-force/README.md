# Brute Force

La version corrigée introduit un `RateLimiter` clé par IP + email, bloquant temporairement les tentatives excessives au lieu de laisser `Auth::attempt` être appelé sans limite. Cela empêche un attaquant d'essayer un grand nombre de mots de passe de façon automatisée sur un même compte. La faille correspond à CWE-307 (Improper Restriction of Excessive Authentication Attempts), comme indiqué dans `knowledge/auth/brute-force.md`.
