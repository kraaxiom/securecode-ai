# Roadmap — SecureCode AI Enterprise Skill

Reprend les 9 phases de `SPEC.md` (section 9), sous forme de checklist cochable. Voir aussi `AGENT_BUILD_GUIDE.md` pour la déclinaison opérationnelle détaillée de ces phases.

## Phase 0 — Socle
- [x] `SKILL.md`
- [x] `README.md`
- [x] `SPEC.md`
- [x] `schemas/finding.schema.json`
- [x] `schemas/patch.schema.json`
- [x] `schemas/report.schema.json`
- [x] `schemas/score.schema.json`
- [x] `schemas/mission.schema.json`
- [x] Squelette complet des dossiers (`knowledge/`, `rules/`, `prompts/`, `templates/`, `examples/`, `tests/`, `docs/`)

## Phase 1 — Connaissance (`knowledge/`)
- [x] Fusion du doublon `sqli-union` / `union-sqli`
- [x] ~222 fiches `knowledge/<categorie>/<slug>.md` rédigées et vérifiées (CWE + OWASP réels, pas de `TODO` restant)

## Phase 2 — Règles SAST (`rules/sast/`)
- [ ] PHP — couverture complète des vulnérabilités de la Phase 1 prioritaire (injections, XSS, authz, auth)
- [ ] JavaScript/Node.js — idem
- [ ] Python
- [ ] Java
- [ ] C#
- [ ] Go
- [ ] Rust

## Phase 3 — Remédiation (`rules/remediation/`)
- [ ] 1 fichier `remediation/<slug>.md` par vulnérabilité, avec diffs avant/après multi-langages et checklist post-patch

## Phase 4 — Prompts IA (`prompts/`)
- [x] `detect.md`
- [ ] `explain.md`
- [ ] `patch.md`
- [ ] `generate-tests.md`
- [ ] `score.md`
- [ ] `report.md`
- [ ] `architecture-review.md`

## Phase 5 — Pipeline & CLI
- [ ] Orchestration du pipeline complet (`docs/Workflow.md` section 4.7 : scan → détection → explication → patch → test → score → rapport)
- [ ] Spécification CLI (`docs/CLI.md`)

## Phase 6 — Scoring & Rapports
- [ ] `templates/report-owasp-top10.md`
- [ ] `templates/report-asvs.md`
- [ ] `templates/compliance-checklist.md`
- [ ] `templates/patch-diff.md`

## Phase 7 — DAST/IAST/RASP/WAF (optionnel avancé)
- [ ] `rules/dast/`
- [ ] `rules/iast/`
- [ ] `rules/rasp/`
- [ ] `rules/waf/`

## Phase 8 — Packaging Capafy
- [x] `LICENSE.md`
- [x] `CHANGELOG.md`
- [x] `CONTRIBUTING.md`
- [ ] Documentation de vente / positionnement commercial finalisé

---

À chaque phase, valider avec `tests/eval-set.json` (précision de détection, taux de faux positifs) avant de passer à la phase suivante — voir `SPEC.md` section 9 et `AGENT_BUILD_GUIDE.md` section 7 (checklist de fin de projet).
