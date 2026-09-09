---
id: missing-audit
category: database
cwe: CWE-778
owasp: A09:2021-Security-Logging-and-Monitoring-Failures
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Missing Audit (absence de journalisation d'audit)

## Description
Ce pattern décrit l'absence de journalisation des opérations sensibles effectuées sur la base de données : accès aux données personnelles, modifications de privilèges, suppressions massives, connexions administratives. Sans piste d'audit fiable, une organisation ne peut ni détecter une compromission en cours, ni déterminer l'étendue d'un incident après coup, ni satisfaire des obligations réglementaires de traçabilité.

## Où ça apparaît typiquement
- Bases de données de production sans journalisation des requêtes d'administration (connexions, modifications de schéma, changement de rôle).
- Accès aux tables contenant des données sensibles (données personnelles, financières) sans trace de qui a consulté quoi et quand.
- Suppressions ou modifications massives de données possibles sans conservation d'un historique ou d'une piste d'audit immuable.
- Logs d'audit stockés au même endroit que les données qu'ils surveillent, modifiables par le même compte technique compromis.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de configuration de journalisation native de la base de données pour les opérations sensibles (DDL, changements de privilèges, connexions).
- Aucune trace applicative des accès en lecture aux données sensibles (qui a consulté quelle donnée personnelle et quand).
- Journaux d'audit stockés dans le même système/compte que les données surveillées, sans séparation des privilèges d'écriture.
- Pas de conservation définie ni de procédure de revue périodique des journaux d'audit existants.

## Remédiation
- Activer la journalisation native de la base de données pour les opérations sensibles (connexions, DDL, changements de privilèges).
- Journaliser au niveau applicatif les accès aux données sensibles avec l'identité de l'accédant, l'horodatage et la nature de l'opération.
- Stocker les journaux d'audit dans un système séparé, en écriture seule pour les comptes applicatifs, afin d'empêcher leur falsification.
- Définir une politique de rétention et une revue périodique des journaux d'audit, alignée sur les obligations réglementaires applicables.
- Voir `rules/remediation/missing-audit.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/missing-audit/`.

## Références
- OWASP Top 10: A09:2021 – Security Logging and Monitoring Failures
- CWE-778: Insufficient Logging
