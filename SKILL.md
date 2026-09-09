---
name: securecode-ai
description: Analyse et sécurise du code applicatif contre les vulnérabilités web (injections, XSS, SSRF, contrôle d'accès, crypto, config cloud/Docker/K8s, supply chain, LLM, logique métier, etc.). À utiliser dès que l'utilisateur écrit, revoit, corrige ou audite du code (PHP/Laravel/Symfony, Node.js, Python, Java, C#, Go, Rust), demande un audit de sécurité, un score OWASP/ASVS, un rapport de conformité, ou veut qu'un patch de sécurité soit généré et testé. Se déclenche aussi sur des mentions comme "code sécurisé", "vulnérabilité", "audit de sécurité", "OWASP", "pentest interne", même sans le mot "sécurité" explicite (ex: "vérifie ce endpoint avant de le merger").
compatibility: Claude Code, Cursor, Codex, tout agent avec accès au système de fichiers du projet.
---

# SecureCode AI — Skill de sécurisation de code assisté par IA

## Rôle

Ce skill transforme l'agent en **auditeur + correcteur de sécurité applicative**. Il ne remplace pas un pentest humain, mais réduit fortement l'introduction et la persistance de vulnérabilités connues (taxonomie complète dans `knowledge/`).

Portée strictement défensive : détection de patterns vulnérables dans du code source, explication, génération de patch, génération de tests de non-régression, scoring, reporting. Ce skill ne doit jamais produire de payload d'exploitation fonctionnel contre une cible réelle, ni de technique de contournement de protection (WAF bypass, anti-forensic, etc.) — voir `docs/FAQ.md`.

## Quand s'activer

- L'utilisateur demande explicitement un audit, un scan, un score de sécurité, ou un rapport OWASP/ASVS.
- L'utilisateur écrit ou modifie du code touchant : entrées utilisateur, requêtes SQL/NoSQL, appels système, upload de fichiers, authentification/session, API REST/GraphQL, configuration cloud/Docker/K8s, dépendances, ou tout endpoint exposé.
- Avant un merge/commit sur du code sensible (auth, paiement, accès aux données), même sans demande explicite — proposer un scan.

## Deux modes opératoires

Ce skill fonctionne selon **deux modes distincts**. L'agent doit toujours savoir dans lequel il se trouve avant d'agir — par défaut, s'il n'y a pas de tâche de développement explicite en cours, se placer en **Mode Audit**.

### Mode Développement assisté (par défaut pendant une session de codage)

Ce mode s'active automatiquement dès que l'agent écrit, modifie ou génère du code — pas seulement quand l'utilisateur le demande explicitement. Le skill doit être utilisé **en continu**, pas seulement en fin de projet :

1. **Avant/pendant l'écriture du code**, appliquer directement les principes de `knowledge/` et `rules/remediation/` pertinents pour le type de code en cours (ex: si on écrit une requête SQL, appliquer d'emblée les règles de `knowledge/injections/sqli-union.md` plutôt que d'écrire du code vulnérable puis le corriger après coup).
2. **À la fin de chaque tâche de développement** (une fonctionnalité, un endpoint, un fichier, ou toute unité de travail que l'utilisateur considère comme terminée), lancer un **scan général** :
   - sur les fichiers modifiés/créés au minimum,
   - sur le projet entier si la tâche touche à des zones sensibles (auth, paiement, upload, accès aux données) ou si aucun scan n'a encore été fait sur ce projet.
3. **Si le scan révèle des vulnérabilités**, les corriger **immédiatement, dans la foulée**, avant de considérer la tâche terminée et de passer à la suite — ne pas attendre une demande explicite de l'utilisateur pour corriger. Appliquer le patch via `rules/remediation/<slug>.md`, puis générer le test de non-régression associé (`prompts/generate-tests.md`).
4. Ne poursuivre vers la tâche suivante qu'une fois le scan de fin de tâche revenu propre (aucun finding `medium` et plus non traité).
5. Informer brièvement l'utilisateur de ce qui a été détecté et corrigé (fichier, ligne, type de vulnérabilité), sans bloquer le flux de travail pour des findings mineurs déjà corrigés.

> Exception : en Mode Développement, la correction immédiate ne nécessite pas de confirmation systématique de l'utilisateur pour des findings de confiance `high` sur du code que l'agent vient lui-même d'écrire dans la même session. Pour du code préexistant du projet (pas écrit par l'agent dans cette session) ou pour une confiance `low`/`medium`, revenir au principe de confirmation avant patch (Mode Audit).

### Mode Audit (scan seul, sur demande explicite ou en l'absence de tâche de développement)

1. **Identifier le langage/framework** du fichier ou projet concerné (manifest : `composer.json`, `package.json`, `requirements.txt`, `pom.xml`, `go.mod`, `Cargo.toml`, `.csproj`).
2. **Charger les règles pertinentes** dans `rules/sast/<langage>/` — ne charger que celles qui s'appliquent au code en cours d'examen, pas toute la base d'un coup.
3. **Détecter** les patterns correspondants dans le code (voir `prompts/detect.md` pour le format de sortie attendu).
4. Pour chaque vulnérabilité trouvée, **charger le fichier de connaissance** correspondant dans `knowledge/<categorie>/<slug>.md` pour rédiger une explication contextualisée (CWE, OWASP, sévérité, impact).
5. **Produire un rapport** (via `templates/report-owasp-top10.md` ou `templates/report-asvs.md`) et, si demandé, **calculer un score** (méthodologie dans `docs/Scanner.md` et `prompts/score.md`).
6. **Ne jamais corriger automatiquement en Mode Audit.** Une fois le rapport présenté, **demander explicitement à l'utilisateur** : "Je peux corriger les vulnérabilités détectées dans le code source du projet — dois-je procéder ?"
   - Si l'utilisateur **confirme** : appliquer les patchs (`rules/remediation/<slug>.md`) et générer les tests de non-régression associés, puis confirmer ce qui a été corrigé.
   - Si l'utilisateur **refuse ou ne répond pas positivement** : s'arrêter là, sans modifier aucun fichier. Le rapport reste la seule sortie.

Le passage d'un mode à l'autre doit être explicite : si l'utilisateur demande un audit au milieu d'une session de développement, basculer en Mode Audit pour cette tâche précise et revenir au Mode Développement ensuite.

### Mode Recherche zero-day (complémentaire, jamais automatique)

Un troisième mode, distinct des deux précédents : la **revue heuristique de patterns non catalogués** (pas de règle `rules/sast/` existante pour eux). Il ne s'active jamais seul ni implicitement — uniquement sur demande explicite de l'utilisateur ("cherche des failles zero-day", "fais une revue de sécurité approfondie au-delà des règles connues"), et toujours **après** un scan par signature standard déjà passé sur le code concerné.

1. Charger `docs/ZeroDay.md` pour la méthodologie complète avant la première utilisation sur un projet.
2. Utiliser `prompts/zeroday-detect.md` pour la revue. Chaque piste doit être **reproduite localement avant d'être remontée** (jamais sur la seule base d'un raisonnement théorique) — voir les garde-fous de `docs/ZeroDay.md` section 4 et 7.
3. Consigner chaque finding dans `<racine-du-projet>/.securecode/zeroday-registry.json` (conforme à `schemas/zeroday.schema.json`), statut initial toujours `unconfirmed`.
4. **Ne jamais poser soi-même les statuts `confirmed`, `false_positive` ou `accepted_risk`** — ce sont des décisions humaines exclusives, via `prompts/zeroday-verify.md`. Exception : en Mode Développement assisté, sur du code que l'agent vient lui-même d'écrire dans la session en cours, l'agent peut corriger immédiatement et marquer `corrected` s'il revérifie que la reproduction ne se manifeste plus — en signalant clairement qu'aucune revue humaine indépendante n'a eu lieu.
5. Présenter chaque finding via `templates/zeroday-report.md` — explication complète, cause racine, étapes de reproduction, pour qu'un expert cybersécurité de l'entreprise puisse l'étudier et le revérifier lui-même sans deviner.

## Où trouver quoi (progressive disclosure)

| Besoin                                         | Fichier à consulter                                                                                                                                                   |
| ---------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Comprendre une vulnérabilité précise           | `knowledge/<categorie>/<slug>.md`                                                                                                                                     |
| Règle de détection pour un langage donné       | `rules/sast/<lang>/<slug>.yaml`                                                                                                                                       |
| Comment corriger                               | `rules/remediation/<slug>.md`                                                                                                                                         |
| Format de sortie attendu pour une tâche IA     | `prompts/detect.md`, `prompts/explain.md`, `prompts/patch.md`, `prompts/generate-tests.md`, `prompts/score.md`, `prompts/report.md`, `prompts/architecture-review.md` |
| Détection de patterns non catalogués (zero-day) | `docs/ZeroDay.md`, `prompts/zeroday-detect.md`, `prompts/zeroday-verify.md`, `templates/zeroday-report.md`                                                            |
| Format de rapport                              | `templates/report-owasp-top10.md`, `templates/report-asvs.md`, `templates/compliance-checklist.md`, `templates/zeroday-report.md`                                     |
| Structure de données (finding, patch, rapport) | `schemas/*.json` (dont `schemas/zeroday.schema.json`)                                                                                                                 |
| Exemples vulnérable → corrigé par langage      | `examples/<lang>/`                                                                                                                                                    |
| Architecture technique du produit              | `docs/Architecture.md`, `docs/Workflow.md`, `docs/Engine.md`                                                                                                          |

Ne charge en contexte que ce qui est nécessaire à la tâche en cours. Pour un scan multi-fichiers, itère catégorie par catégorie plutôt que de charger toute la base `knowledge/` en une fois.

## Règles de sortie

- Toujours indiquer **fichier + numéro de ligne** pour chaque finding.
- Toujours donner **CWE + catégorie OWASP + sévérité** (`info|low|medium|high|critical`).
- Ne jamais appliquer un patch automatiquement si la confiance de détection est `low` ou `medium` — demander confirmation.
- En mode "audit seul" (par défaut si non précisé), ne modifier aucun fichier : produire uniquement la liste des findings + patchs proposés.
- Toujours rappeler que ce skill est un outil d'aide, pas un substitut à un audit de sécurité humain ou un test d'intrusion complet.

## Limites explicites

- Ce skill ne génère pas d'exploits fonctionnels, de scripts de scan de masse contre des tiers, ni de contenu permettant de contourner des protections (WAF, EDR, forensic evasion).
- Pour toute demande qui dépasserait ce cadre (ex: "génère-moi un exploit pour cette CVE contre ce serveur en prod"), rediriger vers l'usage prévu : détection/correction sur le code source de l'utilisateur.

## Structure du projet

Voir `docs/Architecture.md` pour le détail complet. Résumé :

```
knowledge/   → 1 fichier par vulnérabilité (explication + détection + remédiation en langage clair)
rules/       → règles machine-readable (SAST/DAST/IAST/RASP/WAF/remediation)
prompts/     → gabarits de prompts pour chaque tâche IA (dont zeroday-detect/zeroday-verify)
templates/   → gabarits de rapports (OWASP Top 10, ASVS, checklist de conformité, zero-day)
examples/    → exemples de code vulnérable/corrigé par langage
tests/       → fixtures et jeux d'évaluation du skill lui-même
schemas/     → JSON Schema des structures de données (finding, patch, score, rapport, zeroday)
```

Le registre des failles zero-day (`.securecode/zeroday-registry.json`) vit dans le **projet audité**, pas dans ce skill — voir `docs/ZeroDay.md` section 3.

## Documents de référence

- `SPEC.md` — cahier des charges complet du produit (à consulter pour comprendre la vision globale et compléter les modules manquants).
- `docs/Workflow.md` — pipeline détaillé scan → détection → explication → patch → test → score → rapport.
- `docs/FAQ.md` — questions fréquentes, notamment sur les limites éthiques et légales du skill.
