# Changelog

Toutes les modifications notables de ce skill sont documentées ici. Format inspiré de [Keep a Changelog](https://keepachangelog.com/), versionnage [SemVer](https://semver.org/).

## [Non publié]
### Ajouté
- **Mode Recherche zero-day** : détection heuristique de patterns non catalogués (sans règle `rules/sast/` existante), avec exigence de reproduction locale obligatoire et cycle de statut (`unconfirmed` → `confirmed`/`false_positive`/`accepted_risk` réservés à un humain → `corrected`) — `docs/ZeroDay.md`, `prompts/zeroday-detect.md`, `prompts/zeroday-verify.md`, `templates/zeroday-report.md`, `schemas/zeroday.schema.json`. Le registre des findings vit dans le projet audité (`.securecode/zeroday-registry.json`), pas dans le skill.
- Exemples vulnérable/corrigé (`examples/`) pour les catégories SSRF, crypto, cloud, LLM sur les 7 stacks de langage (en complément des catégories prioritaires déjà couvertes : injections, XSS, auth, authorization)

### À venir
- Exemples vulnérable/corrigé pour les catégories non prioritaires (SSRF, crypto, cloud, LLM, etc.) sur les 7 stacks
- Règles DAST/IAST/RASP/WAF (Phase 7)

## [0.4.0] — 2026-08-29
### Ajouté
- **Module de détection zero-day** (revue heuristique de patterns non catalogués, complémentaire au scan par signature) :
  - `docs/ZeroDay.md` — méthodologie, cycle de vie du statut, garde-fous
  - `prompts/zeroday-detect.md` — revue heuristique avec exigence de reproduction locale
  - `prompts/zeroday-verify.md` — vérification humaine obligatoire, jamais posée par l'agent seul
  - `templates/zeroday-report.md` — rapport lisible par un expert cybersécurité, reproductible
  - `schemas/zeroday.schema.json` — structure de données du registre projet (`<projet>/.securecode/zeroday-registry.json`)
  - Mode Recherche zero-day ajouté à `SKILL.md` (Pipeline C dans `docs/Workflow.md`) et FAQ correspondante

## [0.3.0] — 2026-08-29
### Ajouté
- Couverture SAST complète sur les 7 langages prioritaires (PHP, JS, Python, Java, C#, Go, Rust) pour toutes les catégories code-level de la taxonomie
- 222 recettes `rules/remediation/` (une par vulnérabilité de la taxonomie)
- Les 7 prompts (`detect`, `explain`, `patch`, `generate-tests`, `score`, `report`, `architecture-review`) et les 4 templates de rapport
- Les 5 schémas JSON, les 11 fichiers `docs/`
- Exemples vulnérable/corrigé pour les catégories prioritaires (injections, XSS, auth, autorisation) sur les 7 stacks (`examples/`)
- 3 fixtures de test volontairement vulnérables (PHP/Laravel, Node/Express, Python/Django) et `tests/eval-set.json`
- `LICENSE.md`, `CHANGELOG.md`, `ROADMAP.md`, `CONTRIBUTING.md`

## [0.2.0] — 2026-08-23
### Ajouté
- Les ~222 fiches de connaissance `knowledge/<categorie>/<slug>.md` couvrant l'intégralité de la taxonomie définie dans `SPEC.md` section 2 (injections, XSS, inclusion de fichiers, SSRF, auth, autorisation, sessions, upload, API, front-end, crypto, headers, serveur, cloud, Docker, Kubernetes, CI/CD, supply chain, LLM/IA, logique métier, base de données, WebSocket, GraphQL, fichiers sensibles, vulnérabilités modernes)
- Les 5 schémas JSON (`finding`, `patch`, `report`, `score`, `mission`) définissant les structures de données du moteur
- `ROADMAP.md`, `CONTRIBUTING.md`, `LICENSE.md`

### Corrigé
- Fusion du doublon `knowledge/injections/sqli-union.md` / `union-sqli.md` (UNION-based SQLi) — `union-sqli.md` redirige désormais vers `sqli-union.md`, gabarit de référence conservé

## [0.1.0] — Structure initiale
### Ajouté
- Structure initiale + première vague de connaissances
- `SKILL.md` (point d'entrée fonctionnel, deux modes opératoires : Développement assisté / Audit)
- `SPEC.md` (cahier des charges complet)
- `README.md`
- Premiers exemples de référence : `knowledge/injections/sqli-union.md`, `rules/sast/php/sqli-union.yaml`, `rules/remediation/sqli-union.md`, `prompts/detect.md`, `schemas/finding.schema.json`
