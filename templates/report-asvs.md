# Rapport de sécurité — OWASP ASVS

**Projet :** {{project}}
**Généré le :** {{generated_at}}
**Score global :** {{score.overall_score}}/100 — **Niveau ASVS atteignable :** {{score.asvs_level_achievable}}

---

## Résumé exécutif

{{summary}}

L'ASVS (Application Security Verification Standard) définit trois niveaux d'exigence croissants :
- **Niveau 1** — exigences minimales, applicables à toute application.
- **Niveau 2** — applications traitant des données sensibles (recommandé pour la majorité des applications métier).
- **Niveau 3** — applications critiques (santé, finance, infrastructures sensibles).

---

## Statut par niveau ASVS

### Niveau 1

| Exigence | Statut | Finding(s) associé(s) | Référence |
|---|---|---|---|
| {{req_1_1}} | {{status_1_1}} | {{finding_ref_1_1}} | {{ref_1_1}} |
| {{req_1_2}} | {{status_1_2}} | {{finding_ref_1_2}} | {{ref_1_2}} |
| ... | ... | ... | ... |

### Niveau 2

| Exigence | Statut | Finding(s) associé(s) | Référence |
|---|---|---|---|
| {{req_2_1}} | {{status_2_1}} | {{finding_ref_2_1}} | {{ref_2_1}} |
| ... | ... | ... | ... |

### Niveau 3

| Exigence | Statut | Finding(s) associé(s) | Référence |
|---|---|---|---|
| {{req_3_1}} | {{status_3_1}} | {{finding_ref_3_1}} | {{ref_3_1}} |
| ... | ... | ... | ... |

> Statut possible pour chaque exigence : **conforme** / **non conforme** / **non applicable**. Justifier brièvement chaque "non applicable" (ex: fonctionnalité absente du projet).

---

## Exigences manquantes pour le niveau supérieur

{{score.missing_for_next_level}}

---

## Prochaine étape

Ce rapport est produit en **Mode Audit** : aucune modification n'a été apportée au code source. Pour appliquer les corrections proposées, une confirmation explicite de l'utilisateur est requise (voir `prompts/patch.md`).

> Ce rapport est un outil d'aide à la décision. Il ne remplace pas un audit de sécurité humain ni une certification ASVS formelle.
