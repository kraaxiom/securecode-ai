# Workflow

`SKILL.md` définit **deux modes opératoires distincts**. Ce ne sont pas deux variantes d'un même pipeline linéaire : ce sont deux pipelines différents, avec deux positions différentes pour la porte de confirmation utilisateur. L'agent doit toujours savoir dans lequel des deux il se trouve avant d'agir (par défaut : Mode Audit, en l'absence de tâche de développement explicite en cours).

---

## Pipeline A — Mode Développement assisté

Actif automatiquement dès que l'agent écrit, modifie ou génère du code, en continu pendant la session — pas seulement en fin de projet.

### Diagramme

```
┌────────────────────────────┐
│ Écriture/modification de     │
│ code par l'agent             │
└──────────┬─────────────────┘
           │
           ▼
┌────────────────────────────┐
│ Application immédiate des      │  knowledge/ + rules/remediation/
│ principes défensifs pendant     │  pertinents pour le type de code
│ l'écriture (préventif)          │  en cours (ex: SQL → sqli-union.md
└──────────┬─────────────────┘  AVANT d'écrire la requête)
           │
           ▼
┌────────────────────────────┐
│ Fin de tâche de dev            │  (une feature, un endpoint,
│ (unité de travail terminée)    │  un fichier, ce que l'utilisateur
└──────────┬─────────────────┘  considère comme "fini")
           │
           ▼
┌────────────────────────────┐
│ Scan général                    │  fichiers modifiés/créés au minimum ;
│ (rules/sast/<lang>/*)           │  projet entier si zone sensible
└──────────┬─────────────────┘  (auth/paiement/upload/données) ou
           │                      1er scan du projet
           ▼
      findings[] ?
           │
   ┌───────┴───────┐
   │ non             │ oui
   ▼                 ▼
 [tâche          ┌─────────────────────────────┐
 terminée]        │ PORTE : confidence du finding │
                   └──────────┬──────────────────┘
                              │
              ┌───────────────┴────────────────┐
              │ high + code écrit par l'agent    │ low/medium
              │ dans CETTE session                │  OU code préexistant
              │ → PAS de confirmation requise      │  du projet (pas écrit
              ▼                                    │  par l'agent ici)
    ┌────────────────────┐                         │  → confirmation requise
    │ Correction immédiate │                        │  (comme en Mode Audit,
    │ rules/remediation/    │                        │  voir Pipeline B étape 4)
    │ <slug>.md              │◄───────────────────────┘
    └──────────┬─────────┘   (si confirmé)
               │
               ▼
    ┌────────────────────┐
    │ Génération du test    │  prompts/generate-tests.md
    │ de non-régression      │
    └──────────┬─────────┘
               │
               ▼
    ┌────────────────────┐
    │ Re-scan de vérif       │  aucun finding medium+ non traité
    └──────────┬─────────┘  → sinon, retour à correction
               │
               ▼
    ┌────────────────────┐
    │ Information brève      │  fichier, ligne, type de vuln,
    │ de l'utilisateur        │  sans bloquer le flux pour du mineur
    └──────────┬─────────┘  déjà corrigé
               │
               ▼
         [tâche terminée,
        passage à la suite]
```

### Points clés

- La porte de confirmation est **conditionnelle**, pas absente : elle ne saute que pour `confidence: high` sur du code que l'agent vient lui-même d'écrire dans la session courante. C'est l'exception explicitement décrite dans `SKILL.md`.
- Pour tout le reste (code préexistant, confiance `low`/`medium`), on retombe sur le principe de confirmation avant patch du Mode Audit — jamais d'application automatique silencieuse en dehors du cas d'exception.
- La tâche de développement n'est considérée terminée que lorsque le scan de fin de tâche revient propre.

---

## Pipeline B — Mode Audit

Actif sur demande explicite de l'utilisateur, ou par défaut en l'absence de tâche de développement en cours. Aucune modification de fichier tant que l'utilisateur n'a pas explicitement confirmé.

### Diagramme

```
┌────────────────────────────┐
│ Demande d'audit / scan         │
└──────────┬─────────────────┘
           │
           ▼
┌────────────────────────────┐
│ 1. Identification langage/     │  manifest : composer.json,
│    framework                   │  package.json, requirements.txt,
└──────────┬─────────────────┘  pom.xml, go.mod, Cargo.toml, .csproj
           │
           ▼
┌────────────────────────────┐
│ 2. Chargement des règles       │  rules/sast/<lang>/*.yaml
│    pertinentes seulement       │  (progressive disclosure)
└──────────┬─────────────────┘
           │
           ▼
┌────────────────────────────┐
│ 3. Détection                    │  prompts/detect.md
│    → findings[]                 │  → finding.schema.json
└──────────┬─────────────────┘
           │
           ▼
┌────────────────────────────┐
│ 4. Explication contextuelle    │  knowledge/<cat>/<slug>.md
│    par finding                 │  chargé un par un
└──────────┬─────────────────┘
           │
           ▼
┌────────────────────────────┐
│ 5. Rapport + score (si demandé)│  templates/report-owasp-top10.md
│                                 │  ou report-asvs.md, prompts/score.md
└──────────┬─────────────────┘
           │
           ▼
   ══════════════════════════════
   ║  PORTE DE CONFIRMATION       ║  "Je peux corriger les vulnérabilités
   ║  (obligatoire, explicite)    ║  détectées — dois-je procéder ?"
   ══════════════┬═══════════════
                 │
      ┌───────────┴────────────┐
      │ refus / pas de réponse  │ confirmation explicite
      │ positive                 │
      ▼                          ▼
┌──────────────┐      ┌────────────────────────┐
│ Arrêt.          │      │ 6. Application des        │  rules/remediation/
│ Aucun fichier   │      │    patchs                  │  <slug>.md
│ modifié.        │      └──────────┬──────────────┘
│ Le rapport est  │                 │
│ la seule sortie.│                 ▼
└──────────────┘      ┌────────────────────────┐
                        │ 7. Génération des tests    │  prompts/generate-tests.md
                        │    de non-régression        │
                        └──────────┬──────────────┘
                                   │
                                   ▼
                        ┌────────────────────────┐
                        │ 8. Confirmation à            │  fichier, ligne, ce qui
                        │    l'utilisateur de ce         │  a été corrigé
                        │    qui a été corrigé           │
                        └────────────────────────┘
```

### Points clés

- La porte de confirmation est **inconditionnelle** en Mode Audit : elle intervient toujours après le rapport, jamais avant, et jamais implicitement.
- Sans confirmation explicite positive, aucun fichier du projet n'est modifié — le rapport est la seule sortie possible.
- Si l'utilisateur demande un audit au milieu d'une session de développement (Pipeline A en cours), l'agent bascule sur le Pipeline B pour cette tâche précise puis revient au Pipeline A ensuite.

---

---

## Pipeline C — Mode Recherche zero-day (complémentaire, sur demande explicite uniquement)

Ne s'active jamais seul — toujours après un Pipeline A ou B déjà exécuté sur le code concerné, et uniquement si l'utilisateur le demande explicitement (voir `docs/ZeroDay.md`).

```
┌────────────────────────────┐
│ Scan par signature déjà        │  Pipeline A ou B déjà passé
│ effectué sur le code concerné  │
└──────────┬─────────────────┘
           │
           ▼
┌────────────────────────────┐
│ Demande explicite de revue     │  jamais automatique
│ heuristique zero-day            │
└──────────┬─────────────────┘
           │
           ▼
┌────────────────────────────┐
│ Revue heuristique               │  prompts/zeroday-detect.md
│ (5 méthodes, docs/ZeroDay.md §2)│
└──────────┬─────────────────┘
           │
           ▼
      piste sérieuse ?
           │
   ┌───────┴───────┐
   │ non             │ oui
   ▼                 ▼
[rien à           ┌─────────────────────────────┐
 remonter]         │ Tentative de reproduction       │  local-dev / test-fixture /
                    │ locale effective                │  staging-isolated UNIQUEMENT
                    └──────────┬──────────────────┘
                               │
                    ┌──────────┴──────────┐
                    │ échec de repro        │ succès
                    │ → confidence: low      │
                    │ → reproduced_locally:  │
                    │   false, documenté      │
                    ▼                         ▼
        ┌────────────────────┐   ┌────────────────────┐
        │ Entrée créée quand    │   │ Entrée créée,          │
        │ même, transparente     │   │ confidence conforme     │
        │ sur la limite            │   │ au succès de repro       │
        └──────────┬─────────┘   └──────────┬─────────┘
                   └────────────┬────────────┘
                                ▼
                  ┌────────────────────────┐
                  │ Registre projet            │  <projet>/.securecode/
                  │ status: unconfirmed         │  zeroday-registry.json
                  └──────────┬──────────────┘
                             ▼
                  ┌────────────────────────┐
                  │ Rapport lisible            │  templates/zeroday-report.md
                  │ pour expert humain          │
                  └──────────┬──────────────┘
                             ▼
        ══════════════════════════════════════
        ║  Revue humaine obligatoire            ║  prompts/zeroday-verify.md
        ║  (confirmed / false_positive /        ║  jamais posé par l'agent seul
        ║   accepted_risk)                       ║
        ══════════════════════════════════════
```

### Points clés

- La porte de confirmation humaine est **obligatoire** pour passer de `unconfirmed` à `confirmed`/`false_positive`/`accepted_risk` — jamais posée par l'agent seul, à l'exception du code écrit par l'agent dans la session en cours (même exception de confiance `high` que le Pipeline A), où l'agent peut corriger et marquer `corrected` en revérifiant la reproduction, tout en signalant l'absence de revue indépendante.
- Une entrée du registre reste `unconfirmed` indéfiniment tant qu'aucun humain ne l'a revue — ce n'est pas un blocage du pipeline, juste un état d'attente légitime.
- `corrected` exige toujours un patch appliqué **et** une reproduction rejouée sans succès après coup (`docs/ZeroDay.md` section 5) — jamais présumé.

---

## Étapes communes aux deux pipelines

Certaines briques sont réutilisées identiquement dans les deux pipelines, avec les mêmes contrats de données :

- **Corrélation/dédoublonnage** des findings avant toute explication (`docs/Engine.md`).
- **Génération de patch** toujours basée sur `rules/remediation/<slug>.md` + `prompts/patch.md`, jamais improvisée hors de ce cadre.
- **Génération de tests** toujours via `prompts/generate-tests.md`, produisant un test qui échoue avant patch et passe après.
- **Aucune application silencieuse sans configuration explicite** : dans les deux pipelines, une modification de fichier suppose soit l'exception de confiance `high` en session (Pipeline A), soit une confirmation explicite (Pipeline B, ou Pipeline A hors exception).
