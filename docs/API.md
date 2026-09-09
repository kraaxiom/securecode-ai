# API REST (direction future SaaS)

## 1. Statut

Comme indiqué dans `SPEC.md` section 8, cette API n'est **pas nécessaire à l'usage local du skill** : Claude Code, Cursor et Codex chargent le skill directement via `SKILL.md` et le système de fichiers du projet, sans passer par un réseau. Cette API décrit la direction envisagée pour un futur service distant (scan asynchrone, intégration webhook, offre SaaS autour du skill), afin que l'implémentation actuelle ne ferme pas cette porte. Elle n'est pas un composant livré dans le socle actuel.

## 2. Principes

- Strictement défensive, au même titre que le reste du skill : aucun endpoint ne retourne de payload d'exploitation ni de technique de contournement — seulement des findings, patchs proposés, scores et rapports.
- Toute opération de modification de code reste soumise aux mêmes garde-fous que le pipeline local : un scan asynchrone ne patch jamais automatiquement sans mission confirmée (`mission.schema.json`).
- Les corps de requête/réponse réutilisent directement `schemas/*.json` — pas de format parallèle.

## 3. Authentification

`Authorization: Bearer <token>` sur tous les endpoints. Hors périmètre de ce document : gestion des comptes/licences (sujet commercial, voir `SPEC.md` section 10).

## 4. Endpoints

### `POST /v1/scans`

Démarre un scan asynchrone sur un projet fourni (archive ou référence à un dépôt accessible).

Requête :
```json
{
  "project": "string (nom ou identifiant du projet)",
  "mode": "audit_only | assisted",
  "scope": ["chemins ou glob à inclure"],
  "languages": ["php", "js"]
}
```

`mode` reprend directement `mission.schema.json#mode` — une requête `assisted` ne dispense pas des garde-fous de confiance décrits dans `SKILL.md` (correction immédiate seulement sur du code écrit dans la session, non applicable tel quel à un scan asynchrone sans "session" ; en pratique un scan API en `assisted` doit être traité comme `audit_only` côté patch tant qu'aucune notion de session persistante n'est définie pour l'API).

Réponse `202 Accepted` :
```json
{
  "scan_id": "uuid",
  "status": "queued",
  "mission": { "...": "mission.schema.json" }
}
```

### `GET /v1/scans/{scan_id}`

Poll du statut d'un scan.

Réponse :
```json
{
  "scan_id": "uuid",
  "status": "queued | running | done | failed",
  "progress": { "files_scanned": 42, "files_total": 120 },
  "created_at": "date-time",
  "finished_at": "date-time | null"
}
```

### `GET /v1/scans/{scan_id}/findings`

Liste les findings d'un scan terminé.

Réponse : tableau de `finding.schema.json`. Pagination via `?page=`/`?per_page=`.

### `GET /v1/scans/{scan_id}/report?format=owasp-top10|asvs|checklist`

Retourne le rapport complet, conforme à `report.schema.json`.

### `GET /v1/scans/{scan_id}/score`

Retourne `score.schema.json` seul (utile pour un badge CI/CD sans charger tout le rapport).

### `POST /v1/scans/{scan_id}/findings/{finding_id}/status`

Changement manuel de statut d'un finding (`false_positive`, `accepted_risk`) — voir `docs/Database.md`. Requiert un motif.

```json
{ "status": "false_positive | accepted_risk", "reason": "string", "changed_by": "string" }
```

### `POST /v1/scans/{scan_id}/patches`

Demande la génération de patchs pour une liste de `finding_id`. **Ne les applique jamais automatiquement** : retourne des `patch.schema.json` avec `applied: false`, à approuver via l'endpoint suivant — reflet strict de la porte de confirmation du Mode Audit (`docs/Workflow.md` Pipeline B).

### `POST /v1/patches/{patch_id}/apply`

Applique un patch précédemment généré, uniquement si la mission associée porte un `confirmed_by` non vide. Retourne le `patch.schema.json` mis à jour (`applied: true`, `applied_at` renseigné) et les `tests_ref` générés.

## 5. Webhooks (CI/CD)

`POST` configurable par l'utilisateur, déclenché à la fin d'un scan (`scan.completed`) avec le corps équivalent à `GET /v1/scans/{id}/report`. Utilisé par l'intégration CI/CD décrite dans `docs/CLI.md` pour faire échouer un pipeline sans polling actif.

## 6. Codes d'erreur

| Code | Signification |
|---|---|
| 400 | Requête malformée / scope invalide |
| 401 | Token absent ou invalide |
| 403 | Tentative d'application de patch sans mission confirmée |
| 404 | `scan_id`/`finding_id`/`patch_id` inconnu |
| 409 | Statut de finding déjà fixé de façon incompatible (ex: re-confirmer un patch déjà appliqué) |
| 422 | Corps de requête ne respectant pas le schéma attendu |
| 500 | Erreur interne du service de scan |
