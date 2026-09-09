# Format de rapport — Finding zero-day

Un finding par section, dans l'ordre de sévérité décroissante. Généré à partir d'une entrée de `<racine-du-projet>/.securecode/zeroday-registry.json` conforme à `schemas/zeroday.schema.json`.

---

## {{title}}

**ID registre :** `{{id}}` — **Statut :** {{status}} — **Sévérité estimée :** {{severity_estimate}} — **Confiance :** {{confidence}}
**Fichier :** `{{file}}`{{#if line_start}} (lignes {{line_start}}–{{line_end}}){{/if}}
**Composant affecté :** {{affected_component}}
**CWE :** {{cwe | "Aucun mapping établi — pattern non catalogué"}} — **OWASP :** {{owasp_category | "N/A"}}
**Méthode de découverte :** {{discovery_method}}
{{#if closest_known_pattern}}**Pattern connu le plus proche :** `{{closest_known_pattern}}`{{/if}}

### Description

{{description}}

### Cause racine

{{root_cause}}

### Reproduction

> Rejouable indépendamment par un expert cybersécurité — chaque étape doit être exécutable telle quelle dans l'environnement indiqué.

**Environnement :** {{reproduction.environment}} — **Reproduit par l'agent :** {{reproduction.reproduced_locally ? "Oui" : "Non — à vérifier manuellement"}}

**Préconditions :** {{reproduction.preconditions}}

**Étapes :**
{{#each reproduction.steps}}
{{@index}}. {{this}}
{{/each}}

**Résultat attendu (comportement correct) :** {{reproduction.expected_result}}
**Résultat observé (démontre le problème) :** {{reproduction.observed_result}}
{{#if reproduction.poc_ref}}**Script de reproduction :** `{{reproduction.poc_ref}}`{{/if}}

### Vérification humaine

{{#if verified_by}}
Vérifié par **{{verified_by.name}}** ({{verified_by.role}}) le {{verified_by.verified_at}}.
{{#if verified_by.notes}}> {{verified_by.notes}}{{/if}}
{{else}}
**Non encore vérifié par un expert humain.** Ce finding reste `unconfirmed` tant qu'aucune revue humaine n'a rejoué la reproduction ci-dessus (voir `prompts/zeroday-verify.md`).
{{/if}}

{{#if promoted_to_knowledge_ref}}
### Promotion

Ce pattern a été jugé suffisamment général pour rejoindre la base de connaissance permanente : voir `{{promoted_to_knowledge_ref}}`.
{{/if}}

---

> Rappel (voir `docs/ZeroDay.md`) : ce rapport documente une **hypothèse de vulnérabilité découverte par raisonnement heuristique** sur le code source du projet de l'utilisateur, jamais une affirmation d'exploitation réelle ni une recherche de faille sur un système tiers. Tant que `status` n'est pas `confirmed` ou `corrected` par un humain nommé, ce finding doit être traité comme une piste à instruire, pas comme un fait établi.
