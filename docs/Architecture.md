# Architecture

## 1. Vue d'ensemble

SecureCode AI est un **skill** (pas un service tournant en continu) : un ensemble de fichiers de connaissance, de règles machine-readable et de prompts, chargés à la demande par un agent de code IA (Claude Code, Cursor, Codex...) qui dispose d'un accès au système de fichiers du projet audité. Il n'y a pas de serveur central obligatoire — le pipeline complet peut s'exécuter localement, dans la session de l'agent.

Voir `SPEC.md` section 3 pour l'arborescence complète du dépôt. Ce document décrit comment ces dossiers coopèrent au moment de l'exécution.

## 2. Composants

| Dossier | Rôle | Nature |
|---|---|---|
| `knowledge/` | Explication humaine d'une vulnérabilité (CWE, OWASP, remédiation) | Markdown, chargé à la demande |
| `rules/sast/<lang>/` | Détection de patterns dans le code source | YAML, machine-readable |
| `rules/remediation/` | Recette de correction avant/après par langage | Markdown |
| `prompts/` | Contrat d'entrée/sortie pour chaque tâche IA (detect, explain, patch, generate-tests, score, report, architecture-review) | Markdown |
| `templates/` | Gabarits de rapport final | Markdown |
| `schemas/` | Contrat de données strict (JSON Schema) entre chaque étape | JSON |
| `examples/` | Corpus vulnérable → corrigé, par langage | Code source |
| `tests/` | Fixtures et jeu d'évaluation du skill lui-même | JSON + code |

## 3. Flux de données

Le flux ci-dessous est celui du pipeline générique décrit dans `SPEC.md` section 5. Il est instancié différemment selon le mode opératoire (Mode Développement assisté vs Mode Audit) — voir `docs/Workflow.md` pour le détail des deux pipelines et de la position exacte de la porte de confirmation dans chacun.

```
┌─────────────┐
│  Projet /    │  fichiers modifiés ou arborescence complète
│  fichier     │
└──────┬───────┘
       │
       ▼
┌───────────────────────┐
│ 1. Détection langage    │  manifest : composer.json, package.json,
│    / framework           │  requirements.txt, pom.xml, go.mod,
└──────┬────────────────┘  Cargo.toml, .csproj
       │
       ▼
┌───────────────────────┐
│ 2. Chargement des        │  rules/sast/<lang>/*.yaml
│    règles SAST           │  (uniquement les règles pertinentes,
│    pertinentes           │  jamais toute la base — progressive disclosure)
└──────┬────────────────┘
       │
       ▼
┌───────────────────────┐
│ 3. Matching / détection │  prompts/detect.md
│    → findings[]          │  sortie conforme à finding.schema.json
└──────┬────────────────┘
       │
       ▼
┌───────────────────────┐
│ 4. Corrélation           │  dédoublonnage, regroupement par
│    / dédoublonnage       │  fichier+ligne+rule_id (docs/Engine.md)
└──────┬────────────────┘
       │
       ▼
┌───────────────────────┐
│ 5. Lookup connaissance   │  knowledge/<categorie>/<slug>.md
│    (par finding)         │  chargé un fichier à la fois
└──────┬────────────────┘
       │
       ▼
┌───────────────────────┐
│ 6. Explication            │  prompts/explain.md
│    contextualisée         │
└──────┬────────────────┘
       │
       ├─────────────────────────────────────────┐
       │ (Mode Audit : pause ici, rapport d'abord) │
       ▼                                           │
┌───────────────────────┐                          │
│ 7. Génération de patch   │  rules/remediation/<slug>.md
│    (si autorisé)         │  + prompts/patch.md → patch.schema.json
└──────┬────────────────┘                          │
       │                                           │
       ▼                                           │
┌───────────────────────┐                          │
│ 8. Génération de tests   │  prompts/generate-tests.md
│    de non-régression     │
└──────┬────────────────┘                          │
       │                                           │
       ▼                                           │
┌───────────────────────┐                          │
│ 9. Vérification            │  suite de tests projet
│    (tests passent)        │  + tests générés
└──────┬────────────────┘                          │
       │◄──────────────────────────────────────────┘
       ▼
┌───────────────────────┐
│ 10. Scoring               │  prompts/score.md → score.schema.json
│    (pondération §7 SPEC) │
└──────┬────────────────┘
       │
       ▼
┌───────────────────────┐
│ 11. Rapport                │  templates/report-owasp-top10.md
│                             │  templates/report-asvs.md
│                             │  → report.schema.json
└───────────────────────┘
```

## 4. Contrat entre étapes

Chaque flèche du diagramme correspond à un objet JSON validé par un schéma de `schemas/` :

- Étape 3 → 4 : `finding.schema.json` (liste de findings bruts).
- Étape 4 → 7 : `finding.schema.json` (findings dédoublonnés, `status: open`).
- Étape 7 : `patch.schema.json`, référence `finding_id`.
- Étape 8 : `tests_ref` dans `patch.schema.json` pointe vers les tests générés.
- Étape 10 : `score.schema.json`.
- Étape 11 : `report.schema.json`, qui embarque `score` et `findings[]`.

Toute action qui modifie un fichier du projet (étape 7 et au-delà) est subordonnée à une **mission** (`mission.schema.json`) explicitement créée : `mode: audit_only` interdit toute modification sans `confirmed_by` renseigné ; `mode: assisted` autorise la correction immédiate uniquement pour un finding `confidence: high` sur du code écrit par l'agent dans la session en cours (voir `SKILL.md` et `docs/Workflow.md`).

## 5. Pas de composant offensif

Aucun composant de cette architecture ne génère de payload d'exploitation fonctionnel ni de technique de contournement de protection. Le pipeline s'arrête à la détection du pattern vulnérable dans le code source et à sa correction — jamais à la démonstration d'exploitation contre une cible réelle. Voir `docs/FAQ.md`.
