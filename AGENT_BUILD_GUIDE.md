# Guide de développement — SecureCode AI Enterprise Skill

## À suivre par un agent de codage IA (Claude Code) pour terminer le projet

Tu reçois un dépôt **partiellement scaffoldé** : l'arborescence complète existe déjà, avec quelques fichiers réels (exemples) et ~275 fichiers vides ou placeholders à compléter. Ta mission : **compléter tout le contenu manquant**, catégorie par catégorie, en respectant strictement les gabarits ci-dessous. Ne restructure pas les dossiers existants — remplis-les.

Travaille dans l'ordre des phases (section 6). Ne passe à la phase suivante qu'une fois la phase courante entièrement terminée et vérifiée (checklist section 7).

---

## 1. Contraintes non négociables

1. **Portée strictement défensive.** Ce produit détecte, explique, corrige et note la sécurité de code source. Il ne doit **jamais** contenir :
   - de payload d'exploitation fonctionnel prêt à l'emploi contre un système réel,
   - de script d'attaque automatisée contre des cibles tierces,
   - de technique de contournement de protection (bypass WAF, anti-forensic, évasion d'EDR).
     Reste au niveau "reconnaître un pattern vulnérable dans du code source" et "corriger" — jamais "exploiter en production".
2. **Cohérence de schéma.** Tout `finding` généré par le moteur doit être conforme à `schemas/finding.schema.json`. Ne dévie pas de ce format sans mettre à jour le schéma en conséquence.
3. **Traçabilité.** Chaque fichier `knowledge/*.md` doit avoir un CWE réel et cohérent, une catégorie OWASP réelle (pas de valeur inventée). Si tu ne connais pas le CWE exact, cherche-le plutôt que d'inventer un numéro.
4. **Pas de duplication silencieuse.** Le projet contient déjà un doublon connu à corriger dès la Phase 1 : `knowledge/injections/sqli-union.md` et `knowledge/injections/union-sqli.md` couvrent le même sujet (UNION-based SQLi). Fusionne-les : garde `sqli-union.md` (déjà rédigé, sert de gabarit officiel) et supprime `union-sqli.md`, ou transforme ce dernier en simple redirection ("voir sqli-union.md"). Vérifie s'il existe d'autres recoupements du même type ailleurs dans `knowledge/` avant de commencer à écrire (ex: catégories qui se chevauchent entre `server/` et `sensitive-files/`) et propose une fusion cohérente plutôt que deux fiches redondantes.

---

## 1 bis. Modes opératoires du skill (déjà posés dans `SKILL.md`, à respecter partout ailleurs)

`SKILL.md` définit désormais deux modes obligatoires. Tout le reste du dépôt (`docs/Workflow.md`, les prompts, les schémas) doit rester cohérent avec cette définition — ne la réécris pas différemment ailleurs.

**Mode Développement assisté** (par défaut pendant une session de codage) :

- Le skill s'applique en continu pendant l'écriture du code, pas seulement sur demande.
- À la fin de chaque tâche de développement, un scan général est lancé (fichiers modifiés au minimum, projet entier si zone sensible ou premier scan).
- Les vulnérabilités trouvées sont **corrigées immédiatement**, avant de passer à la tâche suivante — pas d'attente de confirmation pour du code de confiance `high` que l'agent vient d'écrire dans la session en cours.
- Un test de non-régression est généré pour chaque correctif appliqué.

**Mode Audit** (scan seul, sur demande explicite ou par défaut hors contexte de développement) :

- Scan → findings → rapport (OWASP Top 10 / ASVS), aucune modification de fichier.
- Une fois le rapport livré, l'agent **demande explicitement la permission** de corriger.
- Si oui → patchs + tests de non-régression appliqués. Si non/pas de réponse positive → arrêt, aucun fichier modifié.

Répercussions sur les livrables à produire :

- `docs/Workflow.md` (section 4.7) doit documenter ces deux modes comme deux pipelines distincts, pas un seul pipeline linéaire — précise clairement où se situe la porte de confirmation dans chacun.
- `schemas/mission.schema.json` (section 4.6) a déjà un champ `mode: audit_only|assisted` — vérifie qu'il correspond exactement à ces deux modes (`audit_only` = Mode Audit, `assisted` = Mode Développement assisté) et que le champ `confirmed_by` n'est requis que pour la bascule audit → correction.
- `prompts/patch.md` et `prompts/report.md` doivent chacun préciser dans quel mode ils s'utilisent et rappeler la règle de confirmation correspondante.
- `docs/FAQ.md` doit inclure une question "Le skill corrige-t-il automatiquement mon code ?" avec la réponse nuancée selon le mode.

---

## 2. Structure actuelle du dépôt (déjà en place — ne pas recréer)

```
security-ai-skill/
├── SKILL.md            ✅ rédigé — point d'entrée du skill, ne pas modifier sans raison
├── SPEC.md              ✅ rédigé — vision produit, architecture, workflow, roadmap
├── README.md            ✅ rédigé
├── LICENSE.md            ⬜ vide
├── CHANGELOG.md          ⬜ vide
├── ROADMAP.md            ⬜ vide
├── CONTRIBUTING.md       ⬜ vide
│
├── docs/                 ⬜ 11 fichiers vides (Architecture.md, Workflow.md, AI.md, Scanner.md,
│                            Engine.md, Database.md, API.md, CLI.md, Performance.md, Plugins.md, FAQ.md)
│
├── knowledge/            🟡 222 fichiers placeholder (frontmatter + TODO), 1 déjà rédigé (sqli-union.md)
│   └── <25 sous-dossiers, liste exacte section 3>
│
├── rules/
│   ├── sast/{php,js,python,java,csharp,go,rust}/   🟡 1 exemple (php/sqli-union.yaml), reste vide
│   ├── dast/, iast/, rasp/, waf/                     ⬜ vides (voir Phase 5, optionnel/avancé)
│   └── remediation/                                  🟡 1 exemple (sqli-union.md), reste vide
│
├── prompts/               🟡 detect.md rédigé, 6 fichiers vides (explain, patch, generate-tests,
│                             score, report, architecture-review)
│
├── templates/             ⬜ 4 fichiers vides (report-owasp-top10.md, report-asvs.md,
│                             compliance-checklist.md, patch-diff.md)
│
├── examples/{php-laravel,node-express,python-django,java-spring,csharp-dotnet,go,rust}/  ⬜ vides
│
├── tests/{fixtures,regression}/, eval-set.json        ⬜ vides
│
└── schemas/               🟡 finding.schema.json rédigé, 4 fichiers vides
    (patch.schema.json, report.schema.json, score.schema.json, mission.schema.json)
```

---

## 3. Manifeste exact des fichiers `knowledge/` à compléter

Chaque fichier ci-dessous existe déjà en placeholder (frontmatter + sections `TODO`). Ta tâche : remplacer le contenu par une fiche complète suivant le gabarit de la section 4.1. Ne renomme pas les fichiers (sauf le doublon signalé en section 1).

- **injections/** (29) : blind-sql-injection, boolean-sqli, code-injection, command-injection, crlf-injection, csv-injection, error-sqli, expression-language-injection, formula-injection, header-injection, http-parameter-pollution, http-request-smuggling, http-response-splitting, imap-injection, ldap-injection, nosql-injection, os-command-injection, regex-dos, smtp-injection, sql-injection, **sqli-union** (déjà rédigé), ssi, ssti, stacked-query-sqli, time-based-sqli, ~~union-sqli~~ (à fusionner), xml-injection, xpath-injection, xxe
- **xss/** (8) : blind-xss, dom-xss, markdown-xss, mutation-xss, reflected-xss, stored-xss, svg-xss, uxss
- **file-inclusion/** (6) : directory-traversal, lfi, path-traversal, rfi, symlink-attack, zip-slip
- **ssrf/** (8) : blind-ssrf, dns-rebinding, elasticsearch-ssrf, metadata-aws, metadata-azure, metadata-gcp, redis-ssrf, ssrf-classique
- **auth/** (11) : brute-force, credential-stuffing, default-credentials, jwt-algorithm-confusion, jwt-none, jwt-weak-secret, mfa-bypass, oauth-misconfiguration, password-spray, session-prediction, weak-password
- **authorization/** (7) : bfla, bola, broken-access-control, forced-browsing, idor, mass-assignment, privilege-escalation
- **sessions/** (8) : cookie-poisoning, missing-httponly, missing-samesite, missing-secure, session-fixation, session-hijacking, session-replay, weak-session-id
- **upload/** (10) : asp-upload, double-extension, imagetragick, jsp-upload, magic-byte-bypass, mime-bypass, php-upload, polyglot-files, web-shell-upload, zip-bomb
- **api/** (9) : api-key-leakage, api-mass-assignment, broken-object-level-authorization, excessive-data-exposure, graphql-api, grpc, missing-rate-limit, rest, soap
- **frontend/** (10) : browser-storage-leak, clickjacking, cors-misconfiguration, csp-weakness, css-injection, dom-clobbering, open-redirect, postmessage-abuse, prototype-pollution, service-worker-abuse
- **crypto/** (11) : des, hardcoded-keys, md5, rc4, sha1, sslv2, sslv3, tls1-0, weak-jwt-secret, weak-random, weak-tls
- **headers/** (8) : missing-coep, missing-coop, missing-csp, missing-hsts, missing-permissions-policy, missing-referrer-policy, missing-x-content-type-options, missing-x-frame-options
- **server/** (9) : backup-files, debug-mode, directory-listing, env-exposure, git-exposure, phpinfo-exposure, stack-trace, svn-exposure, verbose-errors
- **cloud/** (6) : aws-s3-public, azure-blob-public, exposed-secrets, firebase-misconfiguration, gcp-bucket-public, iam-misconfiguration
- **docker/** (5) : docker-api-exposure, docker-socket-exposure, privileged-container, root-container, secrets-in-image
- **kubernetes/** (5) : anonymous-api, dashboard-exposure, k8s-secret-exposure, privileged-pod, rbac-weakness
- **cicd/** (5) : azure-pipeline-leak, circleci-secrets, github-secrets, gitlab-secrets, jenkins-exposure
- **supply-chain/** (5) : dependency-confusion, known-cves, malicious-packages, typosquatting, vulnerable-libraries
- **llm/** (13) : agent-hijacking, data-poisoning, embedding-poisoning, indirect-prompt-injection, jailbreak, memory-poisoning, model-extraction, model-inversion, prompt-injection, rag-poisoning, secret-leakage, token-leakage, tool-injection
- **business-logic/** (8) : account-takeover, business-rule-bypass, coupon-abuse, double-spending, infinite-redemption, price-manipulation, race-condition, workflow-bypass
- **database/** (6) : backup-exposure, dangerous-stored-procedures, exposed-database, missing-audit, weak-encryption, weak-privileges
- **websocket/** (4) : cswsh, message-injection, missing-authentication, origin-bypass
- **graphql/** (5) : alias-abuse, batch-attack, deep-query, dos-query, introspection-enabled
- **sensitive-files/** (17) : backup-zip, composer-lock, config-php, credentials-json, database-sql, docker-compose-yml, dockerfile, ds-store, env, git, hg, id-rsa, known-hosts, package-lock-json, svn, wp-config-php, yarn-lock
- **modern/** (10) : browser-cache-poisoning, cache-poisoning, dangling-markup, dns-rebinding-modern, http2-rapid-reset, prototype-pollution-advanced, request-desynchronization, web-cache-deception, xs-leaks, xs-search

**Total : 222 fiches à rédiger (221 après fusion du doublon).**

---

## 4. Gabarits obligatoires

### 4.1 `knowledge/<categorie>/<slug>.md`

Remplace le placeholder existant en gardant le frontmatter mais en renseignant les vraies valeurs :

```markdown
---
id: <slug>
category: <categorie>
cwe: CWE-XXX # numéro réel, vérifié
owasp: A0X:2021-XXX # ou référence OWASP la plus pertinente (API Top 10, LLM Top 10 pour la catégorie llm/, etc.)
severity_default: info|low|medium|high|critical
languages: [php, js, python, java, csharp, go, rust] # uniquement ceux réellement concernés
---

# <Nom lisible de la vulnérabilité>

## Description

3 à 5 phrases, langage clair, pas de jargon inutile.

## Où ça apparaît typiquement

Liste à puces des contextes de code où ce pattern survient.

## Indicateurs de détection (niveau pattern, pas d'exploitation)

Liste à puces de patterns de code repérables statiquement — jamais un payload d'attaque.

## Remédiation

Résumé des principes de correction + renvoi vers `rules/remediation/<slug>.md`.

## Exemple avant/après

Renvoi vers `examples/<lang>/<slug>/` (à créer en Phase 4).

## Références

OWASP Cheat Sheet pertinent + CWE. Pas de lien inventé — si tu n'es pas sûr de l'URL exacte, cite le nom du document sans lien.
```

Utilise `knowledge/injections/sqli-union.md` comme référence de qualité — reproduis exactement ce niveau de détail.

### 4.2 `rules/sast/<lang>/<slug>.yaml`

Utilise `rules/sast/php/sqli-union.yaml` comme gabarit exact. Pour chaque vulnérabilité et chaque langage listé dans son `languages:`, crée un fichier avec :

- `id`, `language`, `category`, `cwe`, `owasp`, `severity`, `confidence_default`
- `match` : description du pattern + liste de patterns (AST-pattern si possible, sinon regex documentée)
- `exclude_if` : conditions qui invalident le finding (réduire les faux positifs)
- `message`, `false_positive_notes`, `remediation_ref`, `knowledge_ref`

Priorité : commence par **PHP et JavaScript/Node.js** (langages prioritaires de l'auteur du projet), puis Python, Java, C#, Go, Rust.

### 4.3 `rules/remediation/<slug>.md`

Utilise `rules/remediation/sqli-union.md` comme gabarit. Structure : un bloc de code avant/après par langage pertinent, plus une checklist de vérification post-patch en fin de fichier.

### 4.4 `prompts/*.md`

Chaque fichier suit la même structure que `prompts/detect.md` : Rôle → Entrée → Tâche → Sortie attendue (JSON si applicable) → Contraintes. Contenu attendu par fichier :

- `explain.md` : transformer un finding brut en explication pédagogique pour un développeur, en s'appuyant sur `knowledge/<slug>.md`.
- `patch.md` : générer un diff minimal à partir de `rules/remediation/<slug>.md` + du contexte réel du fichier, avec un niveau de confiance explicite.
- `generate-tests.md` : produire un test qui échoue avant patch et passe après, format adapté au framework de test détecté (PHPUnit, Jest, pytest, JUnit, xUnit, go test, cargo test).
- `score.md` : appliquer la méthodologie de scoring de `SPEC.md` section 7 à un ensemble de findings.
- `report.md` : assembler un rapport final à partir des templates de `templates/`.
- `architecture-review.md` : recommandations structurelles (gestion des secrets, moindre privilège, defense-in-depth) — reste au niveau conseil, pas de scan automatique.

### 4.5 `templates/*.md`

- `report-owasp-top10.md` : gabarit de rapport organisé par les 10 catégories OWASP Top 10 2021, avec sections "Résumé exécutif", "Findings par catégorie", "Recommandations prioritaires".
- `report-asvs.md` : gabarit organisé par niveau ASVS (1/2/3) avec statut par exigence (conforme / non conforme / non applicable).
- `compliance-checklist.md` : checklist cochable dérivée des deux rapports ci-dessus.
- `patch-diff.md` : format standard pour présenter un patch proposé (fichier, avant, après, justification, tests associés).

### 4.6 `schemas/*.schema.json`

Complète les 4 schémas manquants en cohérence avec `finding.schema.json` déjà écrit :

- `patch.schema.json` : `id`, `finding_id`, `file`, `diff`, `confidence`, `applied` (bool), `applied_at`.
- `report.schema.json` : `id`, `generated_at`, `project`, `score`, `findings[]` (référence finding.schema.json), `format` (`owasp-top10|asvs|checklist`).
- `score.schema.json` : `overall_score` (0-100), `asvs_level_achievable` (1|2|3), `breakdown` (par catégorie pondérée selon SPEC.md section 7), `missing_for_next_level[]`.
- `mission.schema.json` : `id`, `scope` (fichiers/dossiers autorisés), `mode` (`audit_only|assisted`), `confirmed_by`, `created_at` — sert de garde-fou pour qu'aucune modification de fichier n'ait lieu hors d'une mission explicitement confirmée par l'utilisateur.

### 4.7 `docs/*.md`

Rédige chacun en t'appuyant sur `SPEC.md` (qui contient déjà l'essentiel de chaque sujet) :

- `Architecture.md` : reprendre et détailler la section 3 de SPEC.md avec un schéma de flux de données.
- `Workflow.md` : détailler la section 5 de SPEC.md (pipeline complet), avec un diagramme texte étape par étape.
- `AI.md` : comment un agent doit charger `knowledge/`, `rules/`, `prompts/` selon la tâche (progressive disclosure, ne pas tout charger d'un coup).
- `Scanner.md` : spécification technique du moteur SAST (comment les règles YAML sont chargées et évaluées).
- `Engine.md` : moteur de règles — ordre de priorité, gestion des faux positifs, dédoublonnage des findings.
- `Database.md` : schéma du store de findings (persistant entre scans, pour suivi de statut `open/patched/false_positive/accepted_risk`).
- `API.md` : endpoints REST pour scan asynchrone (si le produit évolue vers un SaaS).
- `CLI.md` : commandes (`securecode scan`, `securecode fix --interactive`, `securecode report --format=owasp|asvs`, `securecode score`).
- `Performance.md` : stratégie de scan incrémental (ne réanalyser que les fichiers modifiés) + cache de résultats.
- `Plugins.md` : comment ajouter un nouveau langage/framework sans toucher au cœur du moteur.
- `FAQ.md` : inclure explicitement une question "Ce skill peut-il générer des exploits ?" avec réponse "Non, strictement défensif" et une question sur les limites (ne remplace pas un pentest humain).

---

## 5. Fichiers projet restants

- `LICENSE.md` : rédiger une licence adaptée à un produit commercial vendu sur Capafy (à valider par l'auteur avant publication — proposer une clause d'usage commercial autorisé avec attribution, sans redistribution du code source des règles).
- `CHANGELOG.md` : initialiser avec une entrée `v0.1.0 — Structure initiale + première vague de connaissances`.
- `ROADMAP.md` : reprendre les 8 phases de `SPEC.md` section 9 sous forme de checklist cochable.
- `CONTRIBUTING.md` : expliquer le format attendu pour ajouter une nouvelle entrée `knowledge/` + `rules/` (renvoyer vers les gabarits section 4 de ce guide).
- `tests/eval-set.json` : jeu de cas de test `{ "id", "fixture_path", "expected_findings": [...] }` pour mesurer la précision de détection — à peupler au fur et à mesure que `tests/fixtures/` se remplit.
- `examples/<stack>/` : pour chaque vulnérabilité déjà traitée dans `rules/sast/<lang>/`, ajouter un mini-exemple avant/après dans le dossier du stack correspondant.

---

## 6. Ordre d'exécution recommandé (phases)

1. **Phase 1 — Nettoyage + connaissances prioritaires.** Corriger le doublon `union-sqli`/`sqli-union`. Rédiger en priorité les catégories `injections/`, `xss/`, `authorization/`, `auth/` (les plus fréquentes en audit réel).
2. **Phase 2 — Reste de `knowledge/`.** Compléter les 21 autres catégories, dans l'ordre du manifeste section 3.
3. **Phase 3 — Règles SAST + remédiation.** Pour chaque fiche `knowledge/` rédigée, créer la règle `rules/sast/<lang>/` (PHP + JS d'abord) et le fichier `rules/remediation/`.
4. **Phase 4 — Prompts + templates + schemas.** Compléter les 6 prompts restants, les 4 templates, les 4 schémas restants.
5. **Phase 5 — Exemples + tests.** Peupler `examples/` et `tests/fixtures/` avec des mini-projets volontairement vulnérables, un par vulnérabilité prioritaire (au minimum les catégories Phase 1).
6. **Phase 6 — Docs.** Rédiger les 11 fichiers de `docs/`.
7. **Phase 7 — Fichiers projet.** `LICENSE.md`, `CHANGELOG.md`, `ROADMAP.md`, `CONTRIBUTING.md`.
8. **Phase 8 — Validation finale.** Exécuter la checklist section 7 sur l'ensemble du dépôt.

---

## 7. Checklist de fin de projet (definition of done)

- [ ] Aucun fichier `knowledge/*.md` ne contient encore la mention `TODO`.
- [ ] Aucun doublon thématique dans `knowledge/` (vérifié manuellement catégorie par catégorie).
- [ ] Chaque fichier `knowledge/*.md` a un CWE et une catégorie OWASP réels et vérifiés (pas de valeur générique laissée par erreur).
- [ ] Au moins PHP et JS ont une règle `rules/sast/` pour chaque vulnérabilité de la Phase 1.
- [ ] Chaque règle `rules/sast/` référence un `remediation_ref` et un `knowledge_ref` qui existent réellement (pas de lien mort).
- [ ] Les 7 prompts de `prompts/` sont rédigés et testés sur au moins un cas réel chacun.
- [ ] Les 5 schémas JSON sont valides (`jsonschema` ou équivalent) et cohérents entre eux (mêmes noms de champs pour les mêmes concepts).
- [ ] `docs/FAQ.md` répond explicitement aux questions sur les limites éthiques/légales du produit.
- [ ] Aucun fichier du dépôt ne contient de payload d'exploitation fonctionnel, de script offensif, ou de technique de contournement de protection.
- [ ] Le pipeline complet (section 5 de `SPEC.md`) a été démontré de bout en bout sur au moins un mini-projet fixture par langage prioritaire (PHP, JS).
- [ ] Les deux modes opératoires (section 1 bis) sont démontrés séparément : un run en Mode Développement assisté où une vulnérabilité introduite est détectée et corrigée avant la fin de la tâche, et un run en Mode Audit où le rapport est produit puis la correction est bloquée tant que la permission n'est pas donnée.

---

## 8. Documents de référence déjà fournis

- `SPEC.md` — vision produit complète, section 7 (scoring), section 9 (roadmap détaillée), section 10 (commercialisation Capafy). Ce guide-ci ne remplace pas `SPEC.md`, il en est la déclinaison opérationnelle "à exécuter maintenant".
- `SKILL.md` — point d'entrée déjà fonctionnel, à garder comme référence de style pour toute nouvelle documentation.
- `knowledge/injections/sqli-union.md`, `rules/sast/php/sqli-union.yaml`, `rules/remediation/sqli-union.md`, `prompts/detect.md`, `schemas/finding.schema.json` — les 5 fichiers de référence qualité à reproduire pour tout le reste du dépôt.
