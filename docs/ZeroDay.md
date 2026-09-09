# Détection de failles zero-day (recherche heuristique)

## 1. Différence fondamentale avec le reste du skill

Tout le reste du skill (`rules/sast/`, `knowledge/`) fonctionne par **correspondance de signature** : un pattern de code connu et déjà catalogué (ex: concaténation dans une requête SQL) déclenche une règle déjà écrite. Cette approche est fiable mais, par construction, **ne peut jamais détecter un pattern qui n'a pas encore de règle** — un défaut de logique métier propre au projet, une combinaison dangereuse de primitives correctes individuellement, une déviation par rapport aux conventions du reste du code.

Le module zero-day comble une partie de cet angle mort par **revue heuristique assistée par raisonnement**, pas par une nouvelle base de règles figée. L'agent raisonne sur le code comme le ferait un auditeur humain qui découvre le projet : "qu'est-ce qui semble anormal ici, même si je ne peux pas pointer une règle connue qui le dit explicitement ?"

**Ce module ne prétend pas rivaliser avec un pentest humain ni avec une chasse aux vulnérabilités outillée (fuzzing, symbolic execution).** C'est une couche de raisonnement supplémentaire, structurée pour rester exploitable et vérifiable par des humains — pas une garantie de détection.

## 2. Méthodes de découverte (`discovery_method`)

- **`heuristic-review`** — lecture du code à la recherche de patterns dangereux non catalogués : logique de contrôle d'accès non centralisée, état partagé mutable accédé sans synchronisation, hypothèses implicites sur la provenance d'une donnée.
- **`anomaly-detection`** — un fichier/fonction s'écarte notablement des conventions établies ailleurs dans le même projet (ex: tous les autres endpoints valident un token CSRF sauf celui-ci) — l'anomalie elle-même est l'indice, pas un pattern universel.
- **`deviation-from-convention`** — le code ne respecte pas une invariant que le projet lui-même établit (ex: une classe de service qui bypasse la couche de repository que tout le reste du projet utilise pour l'accès aux données).
- **`dangerous-primitive-combination`** — deux éléments individuellement sûrs deviennent dangereux combinés (ex: un cache partagé entre tenants + une clé de cache dérivée d'une donnée utilisateur non namespacée par tenant).
- **`manual-expert-review`** — l'agent documente et structure une piste que l'utilisateur lui a lui-même signalée verbalement, sans que ce soit un pattern qu'il aurait détecté seul.

## 3. Registre de projet (où vivent les findings zero-day)

Contrairement à `knowledge/` et `rules/` qui appartiennent au skill lui-même, les findings zero-day appartiennent au **projet audité**, pas au skill. L'agent maintient un registre dans le projet scanné :

```
<racine-du-projet>/.securecode/zeroday-registry.json
```

- Un tableau d'objets conformes à `schemas/zeroday.schema.json`.
- Créé au premier finding zero-day détecté sur ce projet, jamais recréé/écrasé ensuite — toujours mis à jour en ajoutant ou en modifiant une entrée par `id`.
- Ce fichier doit être committé dans le dépôt du projet (pas ignoré via `.gitignore`) : c'est la trace persistante que l'équipe sécurité doit pouvoir consulter et faire évoluer dans le temps, au même titre qu'un ticket de suivi.
- Si le projet utilise déjà un tracker externe (Jira, Linear, GitHub Issues), le registre local reste la source de vérité structurée pour l'agent ; l'utilisateur peut demander au skill de générer le texte d'un ticket à partir d'une entrée (voir `prompts/zeroday-detect.md` section Sortie).

## 4. Exigence de reproductibilité (non négociable)

Un finding zero-day **ne doit jamais être remonté sur la seule base d'un raisonnement théorique**. Avant de créer une entrée dans le registre, l'agent doit :

1. Formuler des **préconditions** claires (état du système nécessaire).
2. Dérouler des **étapes de reproduction** numérotées, exécutables localement (environnement de dev du projet, une fixture de test, ou un environnement de staging isolé — **jamais** un système tiers, jamais la production).
3. **Rejouer effectivement ces étapes** (`reproduction.reproduced_locally: true`) — via un test, un appel local à l'application en cours d'exécution sur la machine de dev, ou un script de démonstration non destructif versionné dans le projet.
4. Documenter le résultat attendu (comportement correct) et le résultat observé (comportement démontrant le problème).

Si l'agent ne peut pas reproduire localement (ex: il faudrait un accès à un service tiers réel, ou la faille dépend d'un état de production inaccessible), il crée quand même l'entrée mais avec `confidence: low` et `reproduction.reproduced_locally: false`, en étant explicite sur ce qui manque pour la confirmer — jamais en affirmant une confiance qu'il n'a pas.

**Cette exigence sert un objectif précis : permettre à un expert cybersécurité de l'entreprise de rejouer le problème lui-même, sans deviner ce que l'agent a voulu dire.** Une faille non reproductible n'est qu'une hypothèse — elle doit être présentée comme telle.

## 5. Cycle de vie du statut

| Statut | Qui le pose | Signification |
|---|---|---|
| `unconfirmed` | L'agent | Détecté et reproduit localement par l'agent, pas encore revu par un humain |
| `confirmed` | Un expert humain | A rejoué la reproduction, valide que le problème est réel et exploitable dans ce contexte |
| `corrected` | L'agent, après validation humaine du patch | Un patch a été appliqué **et** la reproduction a été rejouée sans succès (le comportement anormal ne se manifeste plus) |
| `false_positive` | Un expert humain uniquement | Revu et jugé non exploitable — jamais posé automatiquement par l'agent |
| `accepted_risk` | Un expert humain uniquement | Reconnu réel mais délibérément non corrigé (décision documentée dans `verified_by.notes`) |

L'agent peut faire transiter un finding vers `unconfirmed` → (proposition de patch) → **jamais directement vers `corrected`** sans qu'un humain ait d'abord vu et confirmé le problème (`confirmed`), sauf en Mode Développement assisté pour du code que l'agent vient lui-même d'écrire dans la session en cours (même exception de confiance `high` que `SKILL.md` — voir `docs/Workflow.md` Pipeline A), où l'agent peut corriger immédiatement puis marquer `corrected` en documentant que la reproduction ne se manifeste plus, tout en signalant clairement à l'utilisateur qu'aucune revue humaine indépendante n'a eu lieu.

`false_positive` et `accepted_risk` **ne peuvent jamais être posés par l'agent lui-même** — uniquement par un humain, via `verified_by`. C'est plus strict que `finding.schema.json` (où le skill peut proposer un statut) précisément parce qu'un finding zero-day n'a pas la validation implicite d'une règle déjà éprouvée par la communauté (CWE/OWASP établis) : le risque de faux positif structurel est plus élevé.

## 6. Promotion vers une fiche `knowledge/` permanente

Si un pattern zero-day s'avère suffisamment général (pas spécifique à la logique métier unique de ce projet) et se reproduit dans plusieurs contextes, il mérite d'entrer dans la base de connaissance permanente du skill plutôt que de rester un finding ponctuel :

1. Rédiger `knowledge/<categorie>/<slug>.md` selon le gabarit habituel (voir `CONTRIBUTING.md`), avec un vrai CWE si un mapping existe désormais, sinon documenter que c'est un pattern émergent sans CWE établi.
2. Créer la règle `rules/sast/<lang>/<slug>.yaml` correspondante si le pattern est détectable par signature a posteriori.
3. Renseigner `promoted_to_knowledge_ref` sur l'entrée du registre zero-day d'origine.
4. Documenter le changement dans `CHANGELOG.md`.

Cette promotion transforme un pattern "trouvé une fois par raisonnement" en "détectable systématiquement par signature" — c'est l'objectif à terme pour tout pattern zero-day récurrent.

## 7. Garde-fous (strictement alignés avec `docs/FAQ.md` et `SKILL.md`)

- **Jamais de PoC contre un système tiers ou en production.** La reproduction se fait uniquement en local/dev/staging isolé appartenant à l'utilisateur.
- **Jamais d'affirmation d'exploitation réelle "dans la nature".** Ce module documente des faiblesses découvertes dans le code source du projet de l'utilisateur, pas une threat intelligence sur des attaques en cours ailleurs.
- **Jamais de `confidence: high` sans reproduction locale réussie.**
- **Jamais de changement de statut vers `false_positive`/`accepted_risk` par l'agent seul** — ce sont des décisions humaines qui engagent la responsabilité de l'organisation.
- Le terme "zero-day" est utilisé ici au sens **"pattern non catalogué découvert avant qu'une règle formelle n'existe pour lui"**, pas au sens "vulnérabilité activement exploitée dans la nature sur un logiciel tiers" — le skill ne fait aucune recherche de CVE non divulguées sur des produits tiers, ce qui sortirait strictement de son cadre défensif (voir `docs/FAQ.md`).

## 8. Où trouver quoi

| Besoin | Fichier |
|---|---|
| Lancer une revue heuristique zero-day | `prompts/zeroday-detect.md` |
| Faire rejouer une reproduction par un humain | `prompts/zeroday-verify.md` |
| Format de rapport pour un finding zero-day | `templates/zeroday-report.md` |
| Structure de données | `schemas/zeroday.schema.json` |
| Emplacement du registre projet | `<racine-du-projet>/.securecode/zeroday-registry.json` (section 3) |
