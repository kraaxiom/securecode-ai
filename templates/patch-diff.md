# Format de présentation de patch proposé

Un fichier peut contenir plusieurs entrées si plusieurs patchs sont proposés dans la même session. Une entrée par patch, dans l'ordre de priorité (sévérité du finding corrigé décroissante).

---

## Patch #{{n}} — {{finding.rule_id}}

**Fichier :** `{{patch.file}}`
**Finding associé :** `{{patch.finding_id}}` — {{finding.message}}
**CWE :** {{finding.cwe}} — **OWASP :** {{finding.owasp_category}} — **Sévérité :** {{finding.severity}}
**Confiance du patch :** {{patch.confidence}}
**Statut :** {{patch.applied ? "appliqué" : "proposé, en attente de confirmation"}}

### Avant
```{{lang}}
{{code_avant}}
```

### Après
```{{lang}}
{{code_apres}}
```

### Diff
```diff
{{patch.diff}}
```

### Justification
{{justification}}

> Expliquer en 2-4 phrases pourquoi ce changement corrige le finding, en référence à `rules/remediation/{{slug}}.md`, et signaler tout effet de bord potentiel (ex : changement de type retourné, contrainte supplémentaire sur l'entrée).

### Tests associés
{{#each patch.tests_ref}}
- `{{this}}`
{{/each}}

> Chaque test doit échouer sur le code "avant" et passer sur le code "après" (voir `prompts/generate-tests.md`).

---

> Rappel : en Mode Audit, aucun patch listé ici n'est appliqué sans confirmation explicite de l'utilisateur. En Mode Développement assisté, seuls les patchs de confiance `high` sur du code écrit dans la session en cours peuvent être marqués "appliqué" sans confirmation préalable.
