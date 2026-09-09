# Base des findings (Database)

## 1. Objectif

Contrairement à un scan ponctuel isolé, un projet audité à plusieurs reprises dans le temps bénéficie d'un **suivi persistant** des findings : savoir qu'une vulnérabilité détectée la semaine dernière a été corrigée, écartée comme faux positif, ou acceptée comme risque connu, plutôt que de la re-signaler à l'identique à chaque scan. Ce document décrit le schéma logique de ce store, construit directement sur `schemas/finding.schema.json` — il ne prescrit pas de moteur de stockage particulier (fichier JSON local, SQLite, base applicative dans un futur SaaS) : voir `docs/API.md` pour l'usage distant éventuel.

## 2. Modèle de données

Le store est une collection de findings, chacun conforme à `finding.schema.json`, avec `status` comme champ pivot du cycle de vie :

| status | Signification |
|---|---|
| `open` | Détecté, non traité |
| `patched` | Un patch a été appliqué et confirmé (lié à un `patch.schema.json` dont `applied: true`) |
| `false_positive` | Revu et jugé non exploitable dans ce contexte (voir `docs/Engine.md` section 2) |
| `accepted_risk` | Reconnu réel mais délibérément non corrigé (décision utilisateur, ex: dette technique assumée, compensation par un contrôle externe) |

## 3. Clé d'identification dans le temps

- **Clé de corrélation** (identité fonctionnelle d'un finding entre deux scans) : `rule_id + file + line_start`, avec tolérance de léger décalage de ligne (voir `docs/Engine.md` section 4).
- **`id`** : identifiant unique du finding pour un scan donné — ne sert pas de clé de suivi.

Au moment d'un nouveau scan, le moteur doit :
1. Chercher, pour chaque nouveau finding détecté, un enregistrement existant avec la même clé de corrélation.
2. S'il existe et que son `status` est `false_positive` ou `accepted_risk`, ne pas le remonter comme `open` — le signaler séparément comme "connu, statut conservé" plutôt que comme nouvelle alerte, sauf si le code à cette ligne a matériellement changé (auquel cas il redevient `open` avec un nouvel examen).
3. S'il existe avec `status: patched` mais que le pattern vulnérable est de nouveau détecté au même endroit (régression), le repasser à `open` et le signaler explicitement comme régression.
4. Sinon, créer un nouvel enregistrement `open`.

## 4. Champs additionnels utiles au suivi (extension raisonnable du schéma)

`finding.schema.json` définit déjà `detected_at`. Pour un usage de store persistant, il est recommandé (sans que cela contredise le schéma, qui n'interdit pas de champs additionnels côté implémentation) de conserver aussi, en dehors de l'objet finding strict :

- `first_seen_at` — première apparition de cette clé de corrélation.
- `last_seen_at` — dernier scan où le finding a été retrouvé `open`.
- `status_changed_at` / `status_changed_by` — traçabilité d'un changement manuel de statut (notamment `false_positive`/`accepted_risk`, qui doivent rester des décisions explicites et attribuables).
- `patch_id` — référence au `patch.schema.json` correspondant quand `status: patched`.

Ces champs sont des métadonnées de suivi, pas des champs du finding lui-même envoyés dans le rapport — le rapport (`report.schema.json`) continue d'embarquer des findings conformes au schéma strict.

## 5. Portée et confidentialité

Le store de findings peut contenir des extraits de code sensibles (`snippet`) et des chemins de fichiers internes au projet. Il reste local au projet/à l'agent par défaut, au même titre que le reste du skill (voir `docs/Architecture.md` section 1 : pas de service central obligatoire). Toute synchronisation vers un store distant (futur SaaS, `docs/API.md`) est un choix explicite de l'utilisateur, jamais un comportement par défaut du skill.

## 6. Utilisation par le pipeline

- **Mode Audit** (`docs/Workflow.md` Pipeline B) : le rapport généré à l'étape 5 doit distinguer les findings réellement nouveaux des findings déjà connus avec un statut résolu, pour éviter de fatiguer l'utilisateur avec des alertes déjà traitées.
- **Mode Développement assisté** (Pipeline A) : le scan de fin de tâche s'appuie sur le store pour savoir si un finding sur du code préexistant a déjà été explicitement `accepted_risk` par le passé — auquel cas il ne bloque pas la fin de la tâche en cours.
- **CI/CD** (`docs/CLI.md`) : le seuil d'échec de pipeline (score, finding `critical` dans le diff) s'évalue sur les findings `open` uniquement — un finding `false_positive` ou `accepted_risk` ne doit jamais faire échouer un build.
