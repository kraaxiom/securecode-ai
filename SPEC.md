# SecureCode AI — Enterprise Skill
## Cahier des charges complet pour agent de développement IA (Claude Code / Cursor / Codex)

**Statut :** Spécification maîtresse — à soumettre telle quelle à un agent de codage IA pour implémentation complète.
**Auteur produit :** Kra — Cybersécurité (Côte d'Ivoire)
**Objectif commercial :** Package distribué sur Capafy sous forme de "Skill" pour agents de code IA (Claude Code, Cursor, Codex/OpenAI, Windsurf), permettant à ces agents d'écrire, auditer et corriger du code avec un niveau de sécurité professionnel (OWASP / ASVS).

---

## 0. Instructions pour l'agent qui implémente ce projet

Tu es chargé de construire un **produit logiciel complet**, pas une simple documentation. Le produit final doit être :
1. Un **Skill** au format standard (SKILL.md + ressources) directement utilisable par Claude Code, Cursor et Codex.
2. Une **base de connaissance** exhaustive sur les vulnérabilités listées en section 2.
3. Un **moteur de règles** (SAST/DAST/IAST/RASP/WAF) exploitable par du code (pas seulement lisible par un humain).
4. Un **ensemble de prompts** structurés pour chaque tâche IA (détection, explication, patch, tests, scoring, rapport).
5. Une **documentation professionnelle** complète (architecture, workflow, API, CLI).

Procède **module par module**, dans l'ordre de la section 9 (roadmap), et valide chaque module avec des tests avant de passer au suivant. Ne saute aucune vulnérabilité de la taxonomie (section 2) : chaque item doit avoir au minimum un fichier de connaissance (`knowledge/`) et une règle de détection (`rules/`).

---

## 1. Vision produit

**SecureCode AI Enterprise Skill** transforme un agent de code IA généraliste en **assistant de sécurité applicative**. Quand un développeur (ou l'agent lui-même) écrit ou modifie du code, le Skill :

- **Détecte** automatiquement les vulnérabilités introduites ou préexistantes, sur un fichier ou un projet entier (PHP, Laravel, Symfony, Node.js, Python, Java, C#, Go, Rust).
- **Explique** chaque vulnérabilité en langage clair, avec le CWE/OWASP correspondant, la sévérité (CVSS) et l'impact métier.
- **Corrige** en générant un patch sécurisé, applicable automatiquement ou en revue humaine.
- **Teste** en générant des tests de sécurité (unitaires + PoC de non-régression).
- **Score** le projet selon un référentiel proche d'OWASP ASVS (niveaux 1/2/3).
- **Documente** en produisant un rapport de conformité (OWASP Top 10, ASVS, checklist) exportable.
- **Conseille** sur l'architecture (ex : séparation des secrets, moindre privilège, defense in depth).

Positionnement commercial : un "linter de sécurité augmenté par IA", packagé comme extension pour les agents de code, vendu en abonnement/licence sur Capafy.

---

## 2. Taxonomie des vulnérabilités couvertes

Cette taxonomie est la **source de vérité**. Chaque entrée doit être représentée dans `knowledge/<categorie>/<vuln-slug>.md` et dans au moins une règle sous `rules/`.

### 2.1 Injections (`knowledge/injections/`)
SQL Injection, Blind SQLi, Time-based SQLi, Boolean SQLi, UNION SQLi, Stacked Query SQLi, Error-based SQLi, NoSQL Injection, LDAP Injection, XPath Injection, XML Injection, XXE, SSTI, SSI, Command Injection, OS Command Injection, Code Injection, Expression Language Injection, CRLF Injection, Header Injection, SMTP Injection, IMAP Injection, CSV/Formula Injection, ReDoS, HTTP Parameter Pollution, HTTP Request Smuggling, HTTP Response Splitting.

### 2.2 XSS (`knowledge/xss/`)
Reflected, Stored, DOM-based, Mutation XSS (mXSS), Universal XSS (UXSS), Blind XSS, SVG-based XSS, Markdown-based XSS.

### 2.3 Inclusion de fichiers (`knowledge/file-inclusion/`)
LFI, RFI, Path Traversal, Directory Traversal, Zip Slip, Symlink Attack.

### 2.4 SSRF (`knowledge/ssrf/`)
SSRF classique, Blind SSRF, DNS Rebinding, Metadata AWS/Azure/GCP, Redis SSRF, Elasticsearch SSRF.

### 2.5 Authentification (`knowledge/auth/`)
Weak Password Policy, Default Credentials, Password Spray, Credential Stuffing, Brute Force, MFA Bypass, OAuth Misconfiguration, JWT `alg:none`, JWT Weak Secret, JWT Algorithm Confusion, Session ID Prediction.

### 2.6 Autorisation (`knowledge/authorization/`)
IDOR, BOLA, BFLA, Broken Access Control, Privilege Escalation, Forced Browsing, Mass Assignment.

### 2.7 Sessions (`knowledge/sessions/`)
Session Fixation, Session Hijacking, Cookie Poisoning, Weak Session ID, Session Replay, Missing HttpOnly/Secure/SameSite.

### 2.8 Upload (`knowledge/upload/`)
Web Shell Upload, PHP/ASP/JSP Upload, Polyglot Files, Double Extension, MIME Bypass, Magic Byte Bypass, ZIP Bomb, ImageTragick.

### 2.9 API (`knowledge/api/`)
REST, GraphQL, SOAP, gRPC — Missing Rate Limit, Broken Object Level Authorization, Excessive Data Exposure, Mass Assignment, API Key Leakage.

### 2.10 Front-End (`knowledge/frontend/`)
Clickjacking, CORS Misconfiguration, CSP faible/absente, Open Redirect, DOM Clobbering, Prototype Pollution, CSS Injection, Browser Storage Leak, postMessage abuse, Service Worker Abuse.

### 2.11 Cryptographie (`knowledge/crypto/`)
Weak TLS, SSLv2/SSLv3/TLS1.0, RC4, DES, MD5, SHA1, Weak Random (non-CSPRNG), Hardcoded Keys, Weak JWT Secret.

### 2.12 Headers HTTP (`knowledge/headers/`)
Missing CSP/HSTS/X-Frame-Options/X-Content-Type-Options/Referrer-Policy/Permissions-Policy/COOP/COEP.

### 2.13 Serveur (`knowledge/server/`)
Directory Listing, Backup Files, Git/SVN Exposure, `.env` Exposure, phpinfo, Debug Mode, Stack Trace, Verbose Errors.

### 2.14 Cloud (`knowledge/cloud/`)
AWS S3 Public, Azure Blob Public, GCP Bucket Public, Firebase Misconfiguration, Exposed Secrets, IAM Misconfiguration.

### 2.15 Docker (`knowledge/docker/`)
Docker Socket Exposure, Privileged Container, Root Container, Secrets in Image, Docker API Exposure.

### 2.16 Kubernetes (`knowledge/kubernetes/`)
Dashboard Exposure, RBAC Weakness, Privileged Pod, Anonymous API Access, Secret Exposure.

### 2.17 CI/CD (`knowledge/cicd/`)
GitHub/GitLab Secrets Exposure, Jenkins Exposure, Azure Pipeline Leak, CircleCI Secrets.

### 2.18 Supply Chain (`knowledge/supply-chain/`)
Dependency Confusion, Typosquatting, Malicious Packages, Known CVEs, Vulnerable Libraries.

### 2.19 IA / LLM (`knowledge/llm/`)
Prompt Injection, Jailbreak, Indirect Prompt Injection, Data Poisoning, RAG Poisoning, Embedding Poisoning, Model Extraction, Model Inversion, Token Leakage, Secret Leakage, Tool Injection, Agent Hijacking, Memory Poisoning.

### 2.20 Logique métier (`knowledge/business-logic/`)
Price Manipulation, Coupon Abuse, Race Condition, Double Spending, Workflow Bypass, Business Rule Bypass, Account Takeover, Infinite Redemption.

### 2.21 Base de données (`knowledge/database/`)
Exposed Database, Weak Privileges, Dangerous Stored Procedures, Backup Exposure, Weak Encryption at Rest, Missing Audit Trail.

### 2.22 WebSocket (`knowledge/websocket/`)
Missing Authentication, Origin Bypass, CSWSH, Message Injection.

### 2.23 GraphQL avancé (`knowledge/graphql/`)
Introspection Enabled in Production, Alias Abuse, Deep/Nested Query, Batch Query Attack, DoS via Query Complexity.

### 2.24 Fichiers sensibles exposés (`knowledge/sensitive-files/`)
Liste de patterns : `.env`, `.git/`, `.svn/`, `.hg/`, `.DS_Store`, `backup.zip`, `database.sql`, `config.php`, `wp-config.php`, `composer.lock`, `package-lock.json`, `yarn.lock`, `docker-compose.yml`, `Dockerfile`, `id_rsa`, `known_hosts`, `credentials.json`.

### 2.25 Vulnérabilités modernes (`knowledge/modern/`)
Prototype Pollution (côté serveur/client), Cache Poisoning, HTTP/2 Rapid Reset, Browser Cache Poisoning, Request Desynchronization, Web Cache Deception, Dangling Markup Injection, XS-Leaks, XS-Search, DNS Rebinding.

> **Total : ~230 items de connaissance uniques.** Chaque item = 1 fichier `knowledge/<cat>/<slug>.md` de 40 à 150 lignes suivant le gabarit de la section 4.1.

---

## 3. Architecture du dépôt

```
security-ai-skill/
│
├── SKILL.md                      # Point d'entrée du skill (frontmatter + workflow)
├── README.md                     # Présentation, installation, quickstart
├── LICENSE.md
├── CHANGELOG.md
├── ROADMAP.md
├── CONTRIBUTING.md
│
├── docs/
│   ├── Architecture.md            # Schéma technique complet, flux de données
│   ├── Workflow.md                # Pipeline scan → detect → explain → patch → test → score → report
│   ├── AI.md                      # Comment l'IA doit utiliser prompts/ et knowledge/
│   ├── Scanner.md                 # Spécification du moteur de scan statique/dynamique
│   ├── Engine.md                  # Moteur de règles (chargement, priorité, faux positifs)
│   ├── Database.md                # Schéma de la base de vulnérabilités (findings store)
│   ├── API.md                     # API REST/CLI d'intégration (CI/CD, IDE, webhook)
│   ├── CLI.md                     # Commandes CLI (`securecode scan`, `securecode fix`, ...)
│   ├── Performance.md             # Stratégies de perf sur gros repos (cache, incrémental)
│   ├── Plugins.md                 # Comment étendre par langage/framework
│   └── FAQ.md
│
├── knowledge/                     # Base de connaissance (1 fichier par vulnérabilité)
│   ├── injections/
│   ├── xss/
│   ├── file-inclusion/
│   ├── ssrf/
│   ├── auth/
│   ├── authorization/
│   ├── sessions/
│   ├── upload/
│   ├── api/
│   ├── frontend/
│   ├── crypto/
│   ├── headers/
│   ├── server/
│   ├── cloud/
│   ├── docker/
│   ├── kubernetes/
│   ├── cicd/
│   ├── supply-chain/
│   ├── llm/
│   ├── business-logic/
│   ├── database/
│   ├── websocket/
│   ├── graphql/
│   ├── sensitive-files/
│   └── modern/
│
├── rules/                         # Règles machine-readable (YAML/JSON)
│   ├── sast/                      # Analyse statique par langage (php/, js/, python/, java/, csharp/, go/, rust/)
│   ├── dast/                      # Signatures de scan dynamique (requêtes de test, patterns de réponse)
│   ├── iast/                      # Règles d'instrumentation runtime
│   ├── rasp/                      # Règles de blocage runtime
│   ├── waf/                       # Règles ModSecurity/CRS-compatibles génées depuis la taxonomie
│   └── remediation/                # 1 fichier de patch-recipe par vuln (avant/après, par langage)
│
├── prompts/                       # Prompts structurés pour chaque tâche IA (voir section 6)
│   ├── detect.md
│   ├── explain.md
│   ├── patch.md
│   ├── generate-tests.md
│   ├── score.md
│   ├── report.md
│   └── architecture-review.md
│
├── templates/                     # Gabarits de sortie
│   ├── report-owasp-top10.md
│   ├── report-asvs.md
│   ├── compliance-checklist.md
│   └── patch-diff.md
│
├── examples/                      # Exemples réels vulnérable → corrigé, par langage
│   ├── php-laravel/
│   ├── node-express/
│   ├── python-django/
│   ├── java-spring/
│   ├── csharp-dotnet/
│   ├── go/
│   └── rust/
│
├── tests/                         # Tests du skill lui-même (pas du code scanné)
│   ├── fixtures/                  # Mini-projets volontairement vulnérables (par langage)
│   ├── eval-set.json              # Cas de test pour mesurer précision de détection
│   └── regression/
│
└── schemas/                       # JSON Schema pour toutes les structures de données
    ├── finding.schema.json
    ├── patch.schema.json
    ├── report.schema.json
    ├── score.schema.json
    └── mission.schema.json
```

---

## 4. Gabarits de contenu

### 4.1 Gabarit `knowledge/<categorie>/<slug>.md`

```markdown
---
id: sqli-union
category: injections
cwe: CWE-89
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go, rust]
---

# UNION-based SQL Injection

## Description
[Explication claire, 3-5 phrases, sans détails opérationnels d'exploitation]

## Où ça apparaît typiquement
- Concaténation de requêtes SQL à partir d'entrées utilisateur
- ORM mal utilisé (raw queries)

## Indicateurs de détection (pattern-level, pas de payloads d'attaque)
- Concaténation de chaînes dans une requête SQL avec une variable non échappée
- Absence de requêtes préparées / paramètres liés

## Remédiation
- Requêtes préparées / paramètres liés (exemples par langage → voir rules/remediation/sqli-union.yaml)
- Validation stricte des types d'entrée
- Principe du moindre privilège sur le compte SQL

## Exemple avant/après
[Voir examples/<lang>/sqli-union/]

## Références
- OWASP Cheat Sheet: SQL Injection Prevention
- CWE-89
```

> **Important (garde-fous) :** ces fichiers documentent des **patterns défensifs** de reconnaissance et de correction. Ils ne doivent **jamais** contenir de payloads d'exploitation prêts à l'emploi, de scripts d'attaque automatisés contre des cibles réelles, ni de techniques d'évasion de WAF détaillées. L'agent qui rédige ce contenu doit rester au niveau "reconnaître le pattern vulnérable dans le code source" et "corriger", pas "exploiter en production".

### 4.2 Gabarit `rules/sast/<lang>/<slug>.yaml`

```yaml
id: sqli-union
language: php
match:
  type: ast-pattern     # ou regex si l'AST n'est pas dispo pour le langage
  pattern: |
    call($conn, "query", [string_concat(_, $var, _)])
  confidence: high
message: "Concaténation directe dans une requête SQL — utiliser une requête préparée."
cwe: CWE-89
severity: high
remediation_ref: remediation/sqli-union.md
false_positive_notes:
  - "Si $var provient d'une constante interne, vérifier avant de remonter le finding."
```

### 4.3 Gabarit `rules/remediation/<slug>.md`
Contient, par langage, un diff minimal "avant → après" que l'IA peut appliquer automatiquement, plus une checklist de vérification post-patch (tests à lancer, comportements à valider).

### 4.4 Gabarit `schemas/finding.schema.json`
Champs minimum : `id`, `rule_id`, `file`, `line_start`, `line_end`, `cwe`, `owasp_category`, `severity` (`info|low|medium|high|critical`), `confidence`, `snippet`, `message`, `remediation_ref`, `status` (`open|patched|false_positive|accepted_risk`).

---

## 5. Pipeline fonctionnel (Workflow)

1. **Ingestion** — le Skill détecte le langage/framework du projet (fichiers manifest : `composer.json`, `package.json`, `requirements.txt`, `pom.xml`, `go.mod`, `Cargo.toml`, `.csproj`).
2. **Scan statique (SAST)** — application des règles `rules/sast/<lang>/*` sur l'arborescence, génération de `findings[]` conformes à `finding.schema.json`.
3. **Corrélation** — dédoublonnage, regroupement par fichier/type, filtrage des faux positifs connus.
4. **Explication** — pour chaque finding, l'IA charge `knowledge/<cat>/<slug>.md` et rédige une explication contextualisée au code réel.
5. **Génération de patch** — l'IA utilise `rules/remediation/<slug>.md` + le contexte du fichier pour produire un diff minimal.
6. **Application (optionnelle)** — patch appliqué automatiquement seulement si confiance ≥ seuil configurable, sinon proposé en revue humaine (jamais d'application automatique en mode "audit only").
7. **Génération de tests** — un test unitaire ou d'intégration qui échoue avant le patch et passe après (non-régression).
8. **Vérification de non-régression** — relance de la suite de tests existante du projet + des tests générés.
9. **Scoring** — calcul d'un score de sécurité pondéré (section 7).
10. **Rapport** — génération `report-owasp-top10.md` / `report-asvs.md` + `compliance-checklist.md`.

Ce pipeline doit pouvoir tourner en **mode audit seul** (aucune modification de code) ou en **mode assisté** (proposition de patch avec confirmation) — jamais en application silencieuse sans configuration explicite de l'utilisateur.

---

## 6. Prompts IA (`prompts/`)

Chaque fichier de `prompts/` définit un rôle, un format d'entrée, un format de sortie strict (souvent JSON), et les fichiers de `knowledge/`/`rules/` à charger en contexte. Exemple de structure commune :

```markdown
# Prompt: detect.md

## Rôle
Tu es un analyseur de sécurité statique. Tu ne dois identifier que des patterns de vulnérabilité dans le code fourni, sans jamais produire de payload d'exploitation.

## Entrée
- Extrait de code (fichier + numéros de ligne)
- Langage détecté
- Liste des règles rules/sast/<lang>/*.yaml applicables

## Sortie (JSON strict conforme à schemas/finding.schema.json)

## Contraintes
- Ne pas halluciner de CWE non pertinent
- Si incertain, `confidence: low` plutôt que de ne rien remonter
```

Prévoir un prompt distinct pour : `detect`, `explain`, `patch`, `generate-tests`, `score`, `report`, `architecture-review` (recommandations structurelles : gestion des secrets, séparation des couches, principe de moindre privilège, defense-in-depth).

---

## 7. Modèle de score de sécurité

Score sur 100, inspiré d'OWASP ASVS (niveaux 1/2/3) :

| Pondération | Catégorie |
|---|---|
| 25% | Contrôles d'accès & autorisation |
| 20% | Validation des entrées / injections |
| 15% | Cryptographie & gestion des secrets |
| 15% | Configuration & durcissement serveur/cloud |
| 10% | Gestion de session & authentification |
| 10% | Dépendances & supply chain |
| 5% | Observabilité / audit logging |

Barème de sévérité → déduction de points : `critical: -15`, `high: -8`, `medium: -3`, `low: -1`, plafonné à 0. Le score doit être accompagné du **niveau ASVS atteignable** (1, 2, ou 3) et d'une checklist des items manquants pour le niveau supérieur.

---

## 8. Intégrations (API / CLI / CI-CD)

- **CLI** (`docs/CLI.md`) : `securecode scan`, `securecode fix --interactive`, `securecode report --format=owasp|asvs`, `securecode score`.
- **CI/CD** : action GitHub/GitLab qui échoue le pipeline si score < seuil configuré ou si un finding `critical` est introduit dans le diff.
- **API** (`docs/API.md`) : endpoints REST pour lancer un scan asynchrone et récupérer les résultats (utile pour un futur SaaS autour du Skill).
- **IDE / Agents** : le SKILL.md est le point d'entrée pour Claude Code / Cursor / Codex — ces agents chargent le Skill à la demande, pas besoin d'API pour l'usage local.

---

## 9. Roadmap de développement (ordre recommandé pour l'agent constructeur)

1. **Phase 0 — Socle** : `SKILL.md`, `README.md`, `schemas/*.json`, squelette des dossiers.
2. **Phase 1 — Connaissance** : générer les ~230 fichiers `knowledge/*` (peut être fait par lot de catégories).
3. **Phase 2 — Règles SAST** : couvrir PHP + JS/Node en premier (langages prioritaires de l'auteur), puis Python, Java, C#, Go, Rust.
4. **Phase 3 — Remediation** : 1 fichier `remediation/*.md` par vulnérabilité, avec exemples multi-langages.
5. **Phase 4 — Prompts** : les 7 prompts de la section 6, testés sur les fixtures de `tests/fixtures/`.
6. **Phase 5 — Pipeline & CLI** : orchestrer le workflow de la section 5.
7. **Phase 6 — Scoring & Rapports** : templates OWASP/ASVS + checklist de conformité.
8. **Phase 7 — DAST/IAST/RASP/WAF** (optionnel avancé) : règles dynamiques et runtime.
9. **Phase 8 — Packaging Capafy** : versionnage sémantique, `CHANGELOG.md`, licence, documentation de vente (voir section 10).

À chaque phase, valider avec `tests/eval-set.json` (précision de détection, taux de faux positifs) avant de passer à la phase suivante.

---

## 10. Considérations pour la commercialisation sur Capafy

- **Licence** : définir clairement (usage commercial autorisé ou non, redistribution). À trancher par l'auteur avant publication — indiquer dans `LICENSE.md`.
- **Versionnage** : SemVer (`MAJOR.MINOR.PATCH`), `CHANGELOG.md` tenu à jour à chaque mise à jour de la taxonomie ou des règles.
- **Différenciation** : mettre en avant la couverture multi-langages, l'intégration OWASP ASVS, et le fait que le patch est généré + testé (pas seulement détecté).
- **Support/Mises à jour** : prévoir un canal de mise à jour de la base `knowledge/`/`rules/` (nouvelles CVE, nouveaux frameworks) — ce sera un argument de vente récurrent (abonnement plutôt qu'achat unique).
- **Transparence sur les limites** : préciser dans `README.md` que l'outil réduit le risque mais ne remplace pas un audit de sécurité humain / pentest — important pour la crédibilité et pour éviter les promesses excessives à des clients institutionnels.
- **Aucune fonctionnalité offensive** : le produit doit rester strictement défensif (détection, correction, scoring, reporting). Ne pas inclure de modules générant des exploits fonctionnels, des scripts de scan de masse contre des cibles tierces, ou des fonctionnalités de contournement de protections (WAF bypass automatisé, etc.) — cela protège la valeur commerciale (accepté par les marketplaces professionnelles) et la responsabilité légale de l'auteur.

---

## 11. Livrables attendus de l'agent constructeur

1. Le dépôt complet conforme à l'arborescence de la section 3.
2. Un `SKILL.md` fonctionnel testé sur au moins 3 mini-projets vulnérables (`tests/fixtures/`) couvrant PHP, Node.js et Python.
3. Un rapport de couverture : combien d'items de la taxonomie (section 2) ont un fichier `knowledge/` **et** une règle `rules/sast/` fonctionnelle.
4. Une démonstration du pipeline complet sur un projet fixture : scan → findings → explication → patch proposé → test généré → score → rapport OWASP.
