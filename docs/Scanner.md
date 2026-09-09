# Scanner (moteur SAST)

## 1. Rôle

Le scanner statique (SAST) est le composant qui applique les règles de `rules/sast/<lang>/*.yaml` sur le code source pour produire des `findings[]` conformes à `schemas/finding.schema.json`. C'est l'étape 2-3 du pipeline décrit dans `docs/Workflow.md` (Pipeline A et B).

Le socle SAST reste le moteur principal du skill (détection statique sur le code source). `rules/dast/`, `rules/iast/`, `rules/rasp/`, `rules/waf/` sont des extensions avancées (`SPEC.md` phase 7) couvrant les catégories prioritaires de la taxonomie — voir section 7 ci-dessous pour leur format. Elles complètent le SAST sans le remplacer : un projet peut n'utiliser que le SAST (cas par défaut, aucune dépendance runtime) et activer DAST/IAST/RASP/WAF seulement s'il dispose de l'infrastructure correspondante (environnement de test dynamique, agent d'instrumentation runtime, proxy WAF).

## 2. Chargement des règles

1. Détection du langage/framework via manifest (`composer.json`, `package.json`, `requirements.txt`, `pom.xml`, `go.mod`, `Cargo.toml`, `.csproj`).
2. Chargement de `rules/sast/<lang>/*.yaml` — filtré aux catégories pertinentes pour le fichier examiné (voir `docs/AI.md` section 2-3).
3. Chaque règle est un document YAML indépendant, avec le format défini par `SPEC.md` 4.2 (voir aussi `rules/sast/php/sqli-union.yaml` comme référence) :

```yaml
id: <slug>
language: <lang>
category: <categorie-knowledge>
cwe: CWE-XXX
owasp: <categorie-owasp>
severity: info|low|medium|high|critical
confidence_default: low|medium|high
match:
  type: ast-pattern | regex
  description: <texte>
  patterns: [<liste de patterns>]
  exclude_if: [<conditions textuelles ou vérifiables>]
message: <message affiché dans le finding>
false_positive_notes: [<notes>]
remediation_ref: rules/remediation/<slug>.md
knowledge_ref: knowledge/<categorie>/<slug>.md
```

## 3. Types de matching

### 3.1 `ast-pattern`

Mode privilégié quand un parseur AST est disponible pour le langage (PHP, JS/TS, Python, Java, C#, Go via leurs outillages respectifs). Le pattern décrit une forme structurelle du code (appel de fonction, arguments, opérateurs de concaténation) plutôt qu'une chaîne de caractères — ce qui réduit fortement les faux positifs liés au formatage, aux commentaires, ou aux chaînes de caractères non exécutables qui ressembleraient au pattern.

Exemple (`sqli-union.yaml`) : `mysqli_query($_, string_concat(_, $var, _))` matche un appel `mysqli_query` dont le deuxième argument est une concaténation de chaîne incluant une variable, indépendamment de la mise en forme exacte du code.

### 3.2 `regex`

Utilisé quand aucun AST fiable n'est disponible pour le langage/format ciblé (ex: fichiers de configuration, Dockerfile, YAML Kubernetes, `.env`, certains patterns de fichiers sensibles de `knowledge/sensitive-files/`). Le regex opère sur le texte brut du fichier. Plus sujet aux faux positifs que `ast-pattern` — les règles regex doivent documenter des `exclude_if` plus larges et privilégier une `confidence_default` plus basse (`medium` au lieu de `high`) sauf cas très spécifiques (secret hardcodé avec format reconnaissable, par exemple).

## 4. Confiance (`confidence`)

La confiance n'est pas un simple attribut statique de la règle : `confidence_default` fixe une valeur de départ, mais le moteur peut l'ajuster à la baisse selon le contexte du match (variable dont l'origine est ambiguë, présence d'une bibliothèque d'échappement en amont non modélisée par le pattern). Elle ne doit jamais être ajustée à la hausse sans justification explicite dans le finding.

Conséquence directe sur le pipeline (`SKILL.md`, `docs/Workflow.md`) : `confidence: low` ou `medium` interdit toute application automatique de patch, quel que soit le mode — seule `confidence: high` sur du code écrit par l'agent dans la session courante (Mode Développement assisté) permet de sauter la confirmation utilisateur.

## 5. `exclude_if`

Chaque règle documente des conditions qui, si observées, annulent ou déclassent le finding plutôt que de le supprimer silencieusement :

- Une condition `exclude_if` vérifiable mécaniquement (ex: "la variable passe par `(int)`/`intval()` avant la requête") doit être évaluée par le scanner et empêcher l'émission du finding.
- Une condition non vérifiable mécaniquement (ex: "si un framework applique déjà un échappement systématique en amont, vérifier manuellement") doit être laissée à l'appréciation de l'agent au moment de l'explication (`prompts/explain.md`) — elle ne doit pas être ignorée, mais elle ne doit pas non plus supprimer un finding réel par excès de confiance automatique. Le comportement par défaut en cas de doute reste conforme à `prompts/detect.md` : préférer `confidence: low` à l'absence de finding.

`false_positive_notes` documente des cas connus mais n'a pas de valeur d'exécution — c'est un aide-mémoire pour l'étape d'explication, pas une condition d'`exclude_if`.

## 6. Sortie

La sortie du scanner est un tableau de findings au format `schemas/finding.schema.json`, avec `status: open` par défaut (voir `docs/Database.md` pour le cycle de vie du `status` au-delà du scan initial). Ce tableau passe ensuite à l'étape de corrélation/dédoublonnage décrite dans `docs/Engine.md` avant explication.

## 7. Extensions avancées — DAST / IAST / RASP / WAF

Ces quatre familles de règles ciblent le même `<slug>` que la règle SAST correspondante (même `knowledge_ref`/`remediation_ref`), mais à un stade différent du cycle de vie de l'application. Chaque fichier reste conforme au principe défensif du skill : jamais de payload d'exploitation littéral, seulement des indicateurs de détection ou des conditions de blocage.

### 7.1 DAST (`rules/dast/<slug>.yaml`) — scan dynamique boîte noire

Décrit **comment observer une différence de comportement** de l'application en test (pas un payload d'attaque prêt à l'emploi) :

```yaml
id: <slug>
category: <categorie-knowledge>
cwe: CWE-XXX
owasp: <categorie-owasp>
severity: info|low|medium|high|critical
type: dast
test:
  method: GET|POST|PUT|...
  target: <où l'entrée testée est injectée : paramètre de requête, champ JSON, en-tête HTTP>
  technique: boundary-value | differential-response | timing-anomaly | error-signature | out-of-band
  indicators:
    - <description du signal observable indiquant la vulnérabilité — ex: "temps de réponse significativement supérieur à la baseline pour une valeur induisant un délai conditionnel", jamais la valeur d'entrée elle-même si elle constituerait un payload fonctionnel>
false_positive_notes: [...]
remediation_ref: rules/remediation/<slug>.md
knowledge_ref: knowledge/<categorie>/<slug>.md
```

### 7.2 IAST (`rules/iast/<slug>.yaml`) — instrumentation runtime

Décrit **où instrumenter** le code en exécution (sources/sinks à observer) pour tracer un flux de données non fiable, dans un environnement de test avec l'agent d'instrumentation actif :

```yaml
id: <slug>
category: <categorie-knowledge>
cwe: CWE-XXX
type: iast
instrumentation:
  hook_points:
    - "<langage>: <fonction/API à instrumenter comme source ou sink>"
  taint_check: <description du flux source → sink tracé>
  trigger_condition: <condition runtime qui déclenche une alerte, ex: donnée marquée non fiable atteignant le sink sans passer par une fonction d'échappement connue>
severity: info|low|medium|high|critical
remediation_ref: rules/remediation/<slug>.md
knowledge_ref: knowledge/<categorie>/<slug>.md
```

### 7.3 RASP (`rules/rasp/<slug>.yaml`) — blocage runtime en production

Décrit une **condition de blocage** à l'exécution, pour un agent RASP déployé en production. `action: block` doit toujours documenter le mode d'échec recommandé (`fail_mode`) pour que l'activation reste un choix conscient de l'opérateur :

```yaml
id: <slug>
category: <categorie-knowledge>
cwe: CWE-XXX
type: rasp
guard:
  sink: <fonction/API interceptée à l'exécution>
  block_condition: <condition runtime déclenchant le blocage>
  action: block | sanitize | alert-only
  fail_mode: fail-closed | fail-open   # recommandation, jamais imposée silencieusement
severity: info|low|medium|high|critical
remediation_ref: rules/remediation/<slug>.md
knowledge_ref: knowledge/<categorie>/<slug>.md
```

### 7.4 WAF (`rules/waf/<slug>.conf`) — règle de filtrage périmétrique

Contrairement aux trois précédentes, une règle WAF est **intrinsèquement défensive** (elle bloque, elle n'exploite jamais) — elle peut donc contenir des patterns de détection complets, au format compatible ModSecurity/OWASP CRS (`SecRule`). Chaque fichier documente en commentaire le `<slug>` et le CWE associés, et reste un filtre générique de périmètre (défense en profondeur), jamais un substitut à la correction du code source (`remediation_ref`).

### 7.5 Portée

Ces quatre familles ne couvrent, à ce stade, que les catégories prioritaires de la taxonomie (injections, XSS, authentification, autorisation, SSRF) — pas l'intégralité des ~222 entrées de `knowledge/`. L'absence de règle DAST/IAST/RASP/WAF pour un `<slug>` donné ne signifie pas l'absence de risque, seulement l'absence de couverture à ce stade (voir `docs/Plugins.md` pour étendre).
