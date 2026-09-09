# CLI

Interface en ligne de commande pour piloter le pipeline (`docs/Workflow.md`) hors d'une session d'agent interactive — typiquement en CI/CD, en pre-commit, ou en usage manuel local. Les sous-commandes reflètent directement les étapes du pipeline et respectent les mêmes garde-fous que `SKILL.md` (jamais de patch silencieux sans mode explicite).

## Commandes

### `securecode scan`

Exécute le Pipeline B (Mode Audit) jusqu'au rapport, sans jamais modifier de fichier.

```
securecode scan [chemin] [options]

--lang <liste>            Force la liste de langages (sinon auto-détection par manifest)
--scope <glob>             Restreint le scan à un sous-ensemble de fichiers
--format <json|table>      Format de sortie des findings (défaut: table)
--fail-on <severity>       Sévérité minimale qui fait échouer la commande (exit code non-zéro)
--diff-only                Ne scanner que les fichiers modifiés par rapport à la branche de base (docs/Performance.md)
--store <path>              Chemin du findings store (docs/Database.md), défaut: .securecode/findings.json
```

Sortie : liste de findings (`finding.schema.json`), écrite dans `--store` avec mise à jour du statut selon `docs/Database.md` section 3.

**Ne demande jamais de correction** — c'est une commande de scan seul, équivalente aux étapes 1-5 du Pipeline B.

### `securecode fix --interactive`

Propose des patchs pour les findings `open` du dernier scan, un par un, avec confirmation explicite par finding (ou par lot avec `--all`, qui redemande néanmoins une confirmation globale unique avant application). Reproduit fidèlement la porte de confirmation du Pipeline B — cette commande n'existe pas en variante non interactive silencieuse.

```
securecode fix --interactive [options]

--store <path>              Findings store à utiliser (défaut: .securecode/findings.json)
--severity <min>            Ne proposer de patch que pour cette sévérité minimale
--rule <rule_id>             Limiter à un rule_id précis
--dry-run                    Affiche les diffs proposés sans rien appliquer, même avec confirmation
```

Chaque patch appliqué génère aussi son test de non-régression (`prompts/generate-tests.md`) et met à jour le finding correspondant en `status: patched` dans le store.

Il n'existe volontairement pas de `securecode fix --yes-to-all` non interactif qui appliquerait des patchs sans confirmation humaine explicite — cela violerait le principe de `SKILL.md` selon lequel le Mode Audit ne corrige jamais sans confirmation. La seule voie d'application automatique sans confirmation par finding est celle du Mode Développement assisté, qui vit dans la session de l'agent (pas dans cette CLI) sur du code de confiance `high` écrit dans la session en cours.

### `securecode report --format=owasp|asvs`

Génère le rapport à partir du dernier scan (ou d'un `--store` donné), sans exécuter de nouveau scan.

```
securecode report --format=<owasp|asvs|checklist> [options]

--store <path>               Findings store source
--out <fichier>               Chemin de sortie (défaut: stdout)
--include-resolved             Inclure les findings false_positive/accepted_risk dans une section annexe
```

Sortie conforme à `report.schema.json`, rendue avec `templates/report-owasp-top10.md`, `templates/report-asvs.md` ou `templates/compliance-checklist.md`.

### `securecode score`

Calcule uniquement le score (`score.schema.json`), sans régénérer le rapport complet — utile pour un badge ou un check CI léger.

```
securecode score [options]

--store <path>                Findings store source
--format <json|text>          Format de sortie
--min <0-100>                  Si fourni, exit code non-zéro si le score est strictement inférieur
```

## Codes de sortie

| Code | Signification |
|---|---|
| 0 | Succès, aucun seuil dépassé |
| 1 | Erreur d'exécution (projet illisible, langage non supporté, store corrompu) |
| 2 | Seuil de sévérité dépassé (`scan --fail-on`) |
| 3 | Score sous le seuil configuré (`score --min`) |
| 4 | Nouveau finding `critical` introduit dans le diff (CI/CD, voir ci-dessous) |

## Intégration CI/CD

Conformément à `SPEC.md` section 8, le pipeline CI/CD échoue si :
- le score global (`score.schema.json#overall_score`) est strictement inférieur au seuil configuré (`--min`, exit code 3), **ou**
- un finding `severity: critical` nouvellement introduit apparaît dans le diff de la merge request/pull request (comparaison des findings `open` entre la branche de base et la branche courante, exit code 4) — un finding `critical` déjà présent avant la MR et non touché par le diff ne doit pas, à lui seul, faire échouer cette MR précise (il reste visible dans le rapport global, mais ne bloque pas un changement qui ne l'introduit ni ne l'aggrave).

Exemple d'étape CI :

```
securecode scan --diff-only --format json --store .securecode/findings.json
securecode score --min 70
```

`fix --interactive` n'a pas vocation à tourner en CI/CD non supervisée — elle est réservée à un usage local ou à une étape manuelle explicitement approuvée dans le pipeline (ex: un job manuel déclenché par un mainteneur).
