# Utilisation par l'agent IA (progressive disclosure)

Ce document explique **comment** un agent (Claude Code, Cursor, Codex...) doit charger le contenu de `knowledge/`, `rules/` et `prompts/` pendant une tâche — pas seulement quels fichiers existent (voir `SKILL.md` pour la table de correspondance besoin → fichier).

## 1. Principe directeur

Le skill contient environ 230 fichiers de connaissance, plusieurs centaines de règles SAST réparties sur 7 langages, et 7 prompts. **Ne jamais charger l'ensemble en contexte.** Le coût en tokens serait prohibitif et la majorité du contenu serait sans rapport avec la tâche en cours. La règle est : charger uniquement ce qui est nécessaire à l'étape en cours, un niveau de granularité à la fois.

## 2. Ordre de chargement typique

1. **Identifier le langage/framework** avant toute chose (manifest du projet). Ne pas charger de règles tant que le langage n'est pas connu.
2. **Charger `rules/sast/<lang>/`** — uniquement le dossier du langage détecté, jamais les 7 langages. Si plusieurs langages coexistent dans le projet (ex: backend PHP + frontend JS), charger un langage à la fois, fichier par fichier scanné.
3. **Dans ce dossier, ne charger que les règles pertinentes pour le type de code examiné** — pas tout `rules/sast/php/` d'un coup si le fichier scanné ne touche que l'authentification (charger `auth/*.yaml` et `sessions/*.yaml`, pas `injections/*.yaml` si aucune requête n'est présente). Une heuristique simple : scanner d'abord les imports/fonctions utilisées dans le fichier pour cibler les catégories de règles à charger.
4. **`knowledge/<categorie>/<slug>.md` se charge finding par finding**, jamais en bloc. Un scan qui remonte 12 findings de 5 catégories différentes ne justifie pas de charger toute une catégorie `knowledge/` — seulement les slugs effectivement trouvés.
5. **`prompts/<tache>.md` se charge un seul à la fois**, celui correspondant à l'étape en cours du pipeline (`detect.md` pendant la détection, `explain.md` pendant l'explication, etc.) — jamais les 7 prompts en même temps.
6. **`rules/remediation/<slug>.md`** se charge seulement au moment de générer un patch pour ce slug précis, jamais en anticipation.

## 3. Cas par tâche

| Tâche | Ce qu'il faut charger | Ce qu'il ne faut PAS charger |
|---|---|---|
| Écrire une requête SQL (Mode Développement) | `knowledge/injections/sqli-union.md` (et slugs NoSQL/LDAP si pertinent) avant d'écrire le code | Toute la catégorie `injections/`, tous les langages |
| Scanner un fichier PHP | `rules/sast/php/*.yaml` filtré par catégories pertinentes au fichier | `rules/sast/js/`, `rules/sast/python/`, etc. |
| Expliquer un finding `xss-stored` | `knowledge/xss/xss-stored.md` | Les autres fichiers `knowledge/xss/*` |
| Générer un patch | `rules/remediation/<slug>.md` du finding concerné + `prompts/patch.md` | `rules/remediation/` en entier |
| Produire un rapport OWASP | `templates/report-owasp-top10.md` + `prompts/report.md` | `templates/report-asvs.md` (sauf si les deux formats sont demandés) |
| Revue d'architecture | `prompts/architecture-review.md` + `knowledge/` ciblé sur les zones sensibles identifiées | La base de connaissance entière |

## 4. Multi-fichiers / multi-catégories

Pour un scan de projet entier, itérer **catégorie par catégorie** (comme prescrit par `SKILL.md`) plutôt que de charger tout `knowledge/` en une passe : traiter les findings d'injection, produire leurs explications, libérer le contexte, passer à la catégorie suivante (XSS, auth, etc.). Cela garde le contexte de l'agent proportionnel à ce qui est réellement pertinent à chaque instant, et évite qu'un audit sur un petit fichier consomme autant de contexte qu'un audit de dépôt entier.

## 5. Cache de session

Rien n'empêche de garder en mémoire de session un fichier déjà chargé s'il est réutilisé dans la même tâche (ex: `sqli-union.md` revient sur 3 findings distincts du même scan) — inutile de le recharger à chaque occurrence. Mais ce cache ne doit pas s'étendre au-delà de la tâche/mission en cours : une nouvelle mission d'audit repart d'un chargement à la demande, pas d'un contexte hérité de la précédente.

## 6. Lien avec les schémas

Chaque étape produit une structure conforme à `schemas/*.json`. L'agent n'a pas besoin de charger ces schémas en entier à chaque appel : ils définissent le contrat de sortie attendu (voir les prompts, qui rappellent les champs requis), pas un contenu à consulter dynamiquement. Les consulter une fois en début de tâche suffit.
