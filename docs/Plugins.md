# Ajouter un langage ou framework (Plugins)

## 1. Principe

Le moteur (`docs/Engine.md`, `docs/Scanner.md`) ne contient aucune logique spécifique à un langage : il consomme des règles YAML sous `rules/sast/<lang>/` et de la connaissance sous `knowledge/<categorie>/`. Ajouter un nouveau langage ou framework consiste à ajouter du **contenu**, pas à modifier le moteur — c'est le mécanisme d'extension prévu par `SPEC.md` (les langages listés en section 1-2 ne sont pas exhaustifs par construction du format).

## 2. Étapes pour ajouter un nouveau langage

### 2.1 Détection du manifest

Ajouter le fichier manifest caractéristique du langage/framework à la liste utilisée à l'étape 1 du pipeline (`docs/Workflow.md`) — ex: `Gemfile` pour Ruby, `mix.exs` pour Elixir. Cette liste vit dans la logique d'ingestion du pipeline (actuellement énumérée dans `SKILL.md` et `SPEC.md` section 5) ; l'ajouter là suffit, aucun autre composant n'a besoin de connaître la liste des langages supportés.

### 2.2 Créer le dossier de règles

```
rules/sast/<newlang>/
├── sqli-union.yaml
├── xss-reflected.yaml
├── ...
```

Chaque règle suit le format défini par `SPEC.md` 4.2 et illustré par `rules/sast/php/sqli-union.yaml` : `id`, `language`, `category`, `cwe`, `owasp`, `severity`, `confidence_default`, bloc `match` (`type: ast-pattern` si un parseur AST est disponible pour le langage, sinon `type: regex` — voir `docs/Scanner.md` section 3), `message`, `false_positive_notes`, `remediation_ref`, `knowledge_ref`.

Il n'est pas nécessaire de couvrir toute la taxonomie (`SPEC.md` section 2, ~230 items) dès l'ajout d'un langage : commencer par les catégories les plus critiques et fréquentes pour ce langage (typiquement injections, XSS si applicable, auth) puis étendre — chaque règle ajoutée doit référencer un `knowledge_ref` existant ou nouvellement créé (2.3).

### 2.3 Connaissance associée (`knowledge_ref`)

Une règle SAST référence toujours un fichier `knowledge/<categorie>/<slug>.md` via `knowledge_ref`. Deux cas :

- **Le slug existe déjà** (ex: `sqli-union` couvre déjà plusieurs langages) : la nouvelle règle peut réutiliser le même `knowledge_ref` — le fichier de connaissance est volontairement neutre vis-à-vis du langage (`languages: [...]` dans son frontmatter, voir `SPEC.md` 4.1), il suffit d'y ajouter le nouveau langage à cette liste et, si utile, une note spécifique dans la section "Où ça apparaît typiquement".
- **Le slug n'existe pas** (vulnérabilité spécifique au framework, ex: une mauvaise pratique propre à un ORM particulier) : créer un nouveau fichier `knowledge/<categorie>/<slug>.md` suivant le gabarit `SPEC.md` 4.1, avec la même exigence de garde-fou (pas de payload d'exploitation, pattern défensif uniquement).

### 2.4 Remédiation

Ajouter la section correspondant au nouveau langage dans `rules/remediation/<slug>.md` (fichier existant, multi-langages) plutôt que de créer un fichier de remédiation séparé par langage — `SPEC.md` 4.3 prévoit explicitement un fichier de remédiation par vulnérabilité couvrant plusieurs langages en son sein.

### 2.5 Exemples

Ajouter un dossier `examples/<newlang-stack>/<slug>/` avec le couple vulnérable → corrigé, cohérent avec le diff documenté dans `rules/remediation/<slug>.md`.

### 2.6 Tests

Ajouter des fixtures dans `tests/fixtures/<newlang>/` couvrant les nouvelles règles, et des cas correspondants dans `tests/eval-set.json` pour que la précision de détection sur ce langage soit mesurable comme les autres (`SPEC.md` section 9, validation par phase).

## 3. Ce qu'il ne faut jamais faire

- Ne pas introduire de branchement spécifique à un langage dans le moteur (`docs/Engine.md`) ou dans les prompts (`prompts/*.md`) — ces composants doivent rester agnostiques du langage, toute la spécificité vit dans `rules/sast/<lang>/` et `knowledge/`.
- Ne pas dupliquer un fichier de connaissance par langage quand le slug est générique — cela casse la source de vérité unique et complique la maintenance (mise à jour d'une vulnérabilité = un seul fichier à modifier).
- Ne pas introduire, à l'occasion d'un nouveau langage, de contenu qui documenterait des techniques d'exploitation ou de contournement de protection spécifiques à ce langage/framework — les mêmes garde-fous que `SKILL.md`/`SPEC.md` 4.1 s'appliquent à tout nouveau contenu, quel que soit le langage.

## 4. Frameworks au sein d'un langage déjà supporté

Un nouveau framework sur un langage déjà couvert (ex: ajouter Symfony alors que PHP/Laravel existe déjà) suit la même logique à l'échelle du fichier de règle : ajouter les patterns spécifiques au framework dans une règle existante (liste `patterns:` étendue) plutôt que de dupliquer tout le fichier, sauf si la sévérité/confiance par défaut diffère significativement entre les deux frameworks, auquel cas une règle distincte (`<slug>-<framework>.yaml`) est justifiée.
