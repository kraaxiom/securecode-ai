<div align="center">

# SecureCode AI — Enterprise Skill

**Skill de sécurité applicative pour agents de code IA** · **Security-audit skill for AI coding agents**

[![License](https://img.shields.io/badge/license-see_LICENSE.md-informational)](./LICENSE.md)
[![Vulnerabilities](https://img.shields.io/badge/vulnerabilities-222_documented-blue)](./knowledge)
[![Languages](https://img.shields.io/badge/languages-7_covered-blue)](./rules/sast)
[![Scope](https://img.shields.io/badge/scope-defensive_only-success)](./docs/FAQ.md)

[Français](#français) · [English](#english)

</div>

---

<a id="français"></a>
## 🇫🇷 Français

Transforme un agent de code IA généraliste (Claude Code, Cursor, Codex, Windsurf) en **auditeur + correcteur de sécurité applicative**. Il détecte, explique, corrige et note la sécurité d'un projet selon une taxonomie de 222 vulnérabilités documentées — injections, XSS, SSRF, contrôle d'accès, cryptographie, cloud, conteneurs, CI/CD, supply chain, LLM/IA, logique métier et plus — plus un module de détection heuristique pour les patterns non catalogués (zero-day).

### Ce que contient ce dépôt

| Dossier | Contenu |
|---|---|
| `knowledge/` | 222 fiches vulnérabilité — description, détection, remédiation, CWE/OWASP |
| `rules/sast/` | Règles de détection statique — 7 langages (PHP, JS, Python, Java, C#, Go, Rust) |
| `rules/dast/` `rules/iast/` `rules/rasp/` `rules/waf/` | Extensions dynamiques/runtime/périmètre pour les catégories prioritaires |
| `rules/remediation/` | 222 recettes de correction avant/après |
| `examples/` | Code vulnérable → corrigé, par stack (Laravel, Express, Django, Spring, .NET, Go, Rust) |
| `prompts/` | Gabarits de tâche IA (détection, explication, patch, tests, score, rapport, zero-day) |
| `docs/` | Architecture, workflow, FAQ, méthodologie zero-day |

### Installation rapide

**Claude Code** — global (tous tes projets, sans reconfiguration) :
```
git clone https://github.com/kraaxiom/securecode-ai "%USERPROFILE%\.claude\skills\securecode-ai"
```
Ou par projet : clone dans `.claude/skills/securecode-ai/` à la racine du projet.

**Cursor** — clone dans `vendor/securecode-ai/`, puis crée `.cursor/rules/securecode-ai.mdc` référençant `vendor/securecode-ai/SKILL.md` (`alwaysApply: true`).

**Codex** — clone dans `vendor/securecode-ai/`, puis ajoute une référence à `vendor/securecode-ai/SKILL.md` dans `AGENTS.md`.

**Autre agent** — colle le contenu de `SKILL.md` comme instruction système en début de session.

### Utilisation typique

```
"Audite ce dossier src/ pour des vulnérabilités de sécurité et donne-moi un rapport OWASP."
"Ce endpoint /api/users/:id est-il vulnérable à un IDOR ?"
"Génère un patch sécurisé pour cette injection SQL, avec un test de non-régression."
"Cherche des failles zero-day sur le module de paiement."
```

### Principes

- **Défensif uniquement** — aucun exploit fonctionnel, aucun contournement de protection.
- **Traçable** — chaque finding référence un CWE réel, une catégorie OWASP, une sévérité, un fichier de remédiation.
- **Non-intrusif par défaut** — mode audit seul (aucune modification de fichier) sauf confirmation explicite.
- **Complément, pas remplacement** — ne se substitue pas à un audit de sécurité humain ou un pentest.

### Contribuer

Toute contribution passe par une Pull Request — jamais de modification directe sur `main`, jamais de fusion automatique. Voir [`CONTRIBUTING.md`](./CONTRIBUTING.md) pour le format attendu, et [`SECURITY.md`](./SECURITY.md) pour signaler une faille dans le moteur du skill lui-même.

Idées de première contribution : traduire une fiche `knowledge/` en anglais, ajouter un exemple manquant pour un langage, signaler un CWE mal mappé via une issue.

### Licence

Voir [`LICENSE.md`](./LICENSE.md).

---

<a id="english"></a>
## 🇬🇧 English

Turns a general-purpose AI coding agent (Claude Code, Cursor, Codex, Windsurf) into an **application-security auditor and fixer**. It detects, explains, patches, and scores a project's security against a taxonomy of 222 documented vulnerabilities — injections, XSS, SSRF, access control, cryptography, cloud, containers, CI/CD, supply chain, LLM/AI, business logic and more — plus a heuristic detection module for uncatalogued (zero-day) patterns.

> Most content under `knowledge/`, `rules/`, and `examples/` is currently written in French — see [Contributing](#contributing) if you'd like to help translate.

### What's in this repository

| Folder | Contents |
|---|---|
| `knowledge/` | 222 vulnerability write-ups — description, detection, remediation, CWE/OWASP |
| `rules/sast/` | Static detection rules — 7 languages (PHP, JS, Python, Java, C#, Go, Rust) |
| `rules/dast/` `rules/iast/` `rules/rasp/` `rules/waf/` | Dynamic/runtime/perimeter extensions for priority categories |
| `rules/remediation/` | 222 before/after fix recipes |
| `examples/` | Vulnerable → fixed code, per stack (Laravel, Express, Django, Spring, .NET, Go, Rust) |
| `prompts/` | AI task templates (detect, explain, patch, tests, score, report, zero-day) |
| `docs/` | Architecture, workflow, FAQ, zero-day methodology |

### Quick install

**Claude Code** — global (all your projects, no per-project setup):
```
git clone https://github.com/kraaxiom/securecode-ai "%USERPROFILE%\.claude\skills\securecode-ai"
```
Or per-project: clone into `.claude/skills/securecode-ai/` at the project root.

**Cursor** — clone into `vendor/securecode-ai/`, then create `.cursor/rules/securecode-ai.mdc` referencing `vendor/securecode-ai/SKILL.md` (`alwaysApply: true`).

**Codex** — clone into `vendor/securecode-ai/`, then reference `vendor/securecode-ai/SKILL.md` from `AGENTS.md`.

**Any other agent** — paste the contents of `SKILL.md` as a system instruction at the start of a session.

### Typical usage

```
"Audit the src/ folder for security vulnerabilities and give me an OWASP report."
"Is this /api/users/:id endpoint vulnerable to IDOR?"
"Generate a secure patch for this SQL injection, with a regression test."
"Look for zero-day issues in the payment module."
```

### Principles

- **Defensive only** — no functional exploits, no protection-bypass techniques.
- **Traceable** — every finding references a real CWE, OWASP category, severity, and remediation file.
- **Non-intrusive by default** — audit-only mode (no file changes) unless explicitly confirmed.
- **A complement, not a replacement** — does not substitute for a human security audit or a pentest.

### Contributing

Every contribution goes through a Pull Request — no direct commits to `main`, no automatic merging. See [`CONTRIBUTING.md`](./CONTRIBUTING.md) for the expected format, and [`SECURITY.md`](./SECURITY.md) to report a vulnerability in the skill's engine itself.

Good first contributions: translate a `knowledge/` entry to English, add a missing example for a language, flag a mismapped CWE via an issue.

### License

See [`LICENSE.md`](./LICENSE.md).
