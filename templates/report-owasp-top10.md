# Rapport de sécurité — OWASP Top 10 (2021)

**Projet :** {{project}}
**Généré le :** {{generated_at}}
**Score global :** {{score.overall_score}}/100 — **Niveau ASVS atteignable :** {{score.asvs_level_achievable}}

---

## Résumé exécutif

{{summary}}

| Sévérité | Nombre de findings ouverts |
|---|---|
| Critical | {{count.critical}} |
| High | {{count.high}} |
| Medium | {{count.medium}} |
| Low | {{count.low}} |
| Info | {{count.info}} |

---

## Findings par catégorie

### A01:2021 — Broken Access Control
{{findings.A01}}

### A02:2021 — Cryptographic Failures
{{findings.A02}}

### A03:2021 — Injection
{{findings.A03}}

### A04:2021 — Insecure Design
{{findings.A04}}

### A05:2021 — Security Misconfiguration
{{findings.A05}}

### A06:2021 — Vulnerable and Outdated Components
{{findings.A06}}

### A07:2021 — Identification and Authentication Failures
{{findings.A07}}

### A08:2021 — Software and Data Integrity Failures
{{findings.A08}}

### A09:2021 — Security Logging and Monitoring Failures
{{findings.A09}}

### A10:2021 — Server-Side Request Forgery (SSRF)
{{findings.A10}}

> Pour chaque sous-section, présenter les findings sous la forme : fichier + ligne, CWE, sévérité, confiance, message, référence `remediation_ref`. Si aucune vulnérabilité n'est détectée dans une catégorie, l'indiquer explicitement ("Aucun finding ouvert dans cette catégorie").

---

## Recommandations prioritaires

1. {{recommendation_1}}
2. {{recommendation_2}}
3. {{recommendation_3}}

---

## Prochaine étape

Ce rapport est produit en **Mode Audit** : aucune modification n'a été apportée au code source. Pour appliquer les corrections proposées, une confirmation explicite de l'utilisateur est requise (voir `prompts/patch.md`).

> Ce rapport est un outil d'aide à la décision. Il ne remplace pas un audit de sécurité humain ni un test d'intrusion complet.
