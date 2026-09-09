# JWT Algorithm Confusion

Le handler `admin_data` utilisait une `Validation::default()` n'imposant pas explicitement l'algorithme RS256, exposant l'application à une confusion d'algorithme où un token HS256 forgé avec la clé publique RSA comme secret pourrait être accepté (CWE-347). La version corrigée fixe explicitement `Validation::new(Algorithm::RS256)` et restreint `algorithms` à `[RS256]`, empêchant toute déduction de l'algorithme depuis le token lui-même.
