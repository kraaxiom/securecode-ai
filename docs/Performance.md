# Performance sur gros dépôts

## 1. Problème

Un scan complet (charger toutes les règles pertinentes, analyser chaque fichier, charger la connaissance pour chaque finding) devient coûteux en temps et en tokens sur un dépôt de grande taille. La stratégie retenue combine scan incrémental et mise en cache, sans jamais sacrifier la couverture sur le code effectivement modifié.

## 2. Scan incrémental (`--diff-only`)

- Le scan par défaut en Mode Développement assisté porte déjà sur les "fichiers modifiés/créés au minimum" (`SKILL.md`), sauf zone sensible ou premier scan du projet — c'est la base de la stratégie incrémentale.
- En Mode Audit ou en CLI (`securecode scan --diff-only`, `docs/CLI.md`), l'incrémentalité s'appuie sur `git diff` entre la branche courante et une base de comparaison (branche par défaut du dépôt, ou dernier commit scanné connu via le findings store) : seuls les fichiers apparaissant dans le diff sont analysés.
- **Limite volontaire** : un scan incrémental ne détecte pas une vulnérabilité déjà présente dans du code non touché par le diff. C'est un compromis assumé de vitesse pour l'usage courant (dev, CI par MR) — un scan complet reste nécessaire périodiquement (première fois sur le projet, zone sensible touchée, ou à intervalle régulier planifié par l'utilisateur) conformément à `SKILL.md`.
- Un fichier renommé ou déplacé sans modification de contenu ne doit pas déclencher une ré-analyse complète — seule sa clé de corrélation dans le store (`docs/Database.md`) doit être mise à jour vers le nouveau chemin.

## 3. Cache de règles

- Les règles `rules/sast/<lang>/*.yaml` ne changent pas pendant une session de scan : elles peuvent être chargées une fois et réutilisées pour tous les fichiers du même langage, plutôt que rechargées par fichier.
- Le filtrage par pertinence (catégories de règles applicables à un fichier donné, voir `docs/AI.md`) doit être fait sur la base d'une analyse rapide du fichier (imports, mots-clés), pas en évaluant toutes les règles du langage sur chaque fichier — cela réduit le nombre de patterns effectivement testés.

## 4. Cache de résultats

- Un fichier dont le contenu n'a pas changé depuis le dernier scan (comparaison de hash de contenu, pas seulement de date de modification) ne doit pas être ré-analysé : ses findings précédents restent valides et sont repris tels quels depuis le store (`docs/Database.md`), avec `last_seen_at` mis à jour.
- Ce cache de résultats est invalidé dès que : le contenu du fichier change, la règle correspondante change de version, ou l'utilisateur force un scan complet (`securecode scan` sans `--diff-only`, ou option explicite de purge du cache).
- Le cache est local au projet (répertoire `.securecode/` par exemple), au même titre que le findings store — pas de cache partagé implicite entre projets ou entre utilisateurs, pour éviter toute fuite d'information entre dépôts distincts.

## 5. Chargement de connaissance à la demande

Comme détaillé dans `docs/AI.md`, ne charger `knowledge/<categorie>/<slug>.md` que pour les slugs effectivement présents dans les findings d'un scan, jamais par anticipation. Sur un gros scan avec des centaines de findings répartis sur peu de slugs distincts (cas fréquent : beaucoup d'occurrences du même pattern), charger chaque fichier de connaissance une seule fois et le réutiliser pour toutes les occurrences de ce slug plutôt que de le recharger à chaque finding.

## 6. Parallélisation

L'analyse de fichiers indépendants (pas de dépendance de type inter-fichier nécessaire au pattern matching) peut être menée en parallèle. Les patterns qui nécessitent une analyse inter-fichiers (ex: flux de données remontant d'un contrôleur vers un service dans un autre fichier) sortent du cadre du matching `ast-pattern`/`regex` par fichier décrit dans `docs/Scanner.md` et doivent être traités comme un cas explicite documenté dans la règle concernée plutôt que supposés fonctionner en parallèle naïf.

## 7. Priorité au signal, pas au volume

Sur un dépôt très large où un scan complet resterait trop coûteux même incrémental, prioriser les zones sensibles identifiées par `SKILL.md` (auth, paiement, upload, accès aux données) plutôt que de tenter un balayage exhaustif à chaque exécution — un scan complet périodique planifié (CI nocturne, par exemple) reste la façon appropriée d'obtenir une couverture totale sans bloquer le flux de travail quotidien.
