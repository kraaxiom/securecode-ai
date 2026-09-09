# Contribuer à SecureCode AI

Ce document explique le format attendu pour ajouter ou modifier une entrée de la base de connaissance ou du moteur de règles. Il complète `AGENT_BUILD_GUIDE.md` (destiné aux agents IA qui construisent le contenu) et `SPEC.md` (vision produit complète).

## Principes non négociables

1. **Portée strictement défensive.** Aucune contribution ne doit contenir de payload d'exploitation fonctionnel, de script d'attaque automatisée contre des cibles tierces, ni de technique de contournement de protection (WAF bypass, anti-forensic, évasion EDR). Le contenu reste au niveau "reconnaître un pattern vulnérable dans le code source" et "corriger".
2. **Traçabilité.** Chaque vulnérabilité doit référencer un CWE réel et une catégorie OWASP réelle (Top 10 Web, API Top 10, ou LLM Top 10 selon la catégorie). Ne jamais inventer un numéro CWE ou une catégorie — si le numéro exact n'est pas connu, chercher plutôt que d'approximer.
3. **Cohérence de schéma.** Tout `finding` produit doit rester conforme à `schemas/finding.schema.json`. Toute évolution de structure nécessite une mise à jour synchronisée des schémas concernés.
4. **Pas de duplication.** Vérifier qu'une entrée thématiquement équivalente n'existe pas déjà avant d'en créer une nouvelle (voir la fusion `sqli-union`/`union-sqli` comme précédent).

## Ajouter une fiche `knowledge/<categorie>/<slug>.md`

Utiliser `knowledge/injections/sqli-union.md` comme gabarit de référence. Structure obligatoire :

```markdown
---
id: <slug>
category: <categorie>
cwe: CWE-XXX
owasp: A0X:2021-XXX
severity_default: info|low|medium|high|critical
languages: [php, js, python, java, csharp, go, rust]
---

# <Nom lisible>

## Description
## Où ça apparaît typiquement
## Indicateurs de détection (niveau pattern, pas d'exploitation)
## Remédiation
## Exemple avant/après
## Références
```

- Rédiger en français, langage clair, 3 à 5 phrases pour la description.
- Les "indicateurs de détection" décrivent des patterns de code statiquement repérables — jamais un payload d'attaque.
- Toujours renvoyer vers `rules/remediation/<slug>.md` et `examples/<lang>/<slug>/`.

## Ajouter une règle `rules/sast/<lang>/<slug>.yaml`

Utiliser `rules/sast/php/sqli-union.yaml` comme gabarit. Champs obligatoires : `id`, `language`, `match` (type + pattern + confidence), `message`, `cwe`, `severity`, `remediation_ref`, `false_positive_notes`. Prioriser PHP et JavaScript/Node.js avant les autres langages.

## Ajouter une recette `rules/remediation/<slug>.md`

Utiliser `rules/remediation/sqli-union.md` comme gabarit : un bloc avant/après par langage pertinent, plus une checklist de vérification post-patch.

## Revue avant merge

- [ ] Aucun `TODO` restant dans le fichier.
- [ ] CWE et catégorie OWASP vérifiés (pas de valeur générique par erreur).
- [ ] Références croisées valides (`remediation_ref`, `knowledge_ref` pointent vers des fichiers existants).
- [ ] Aucun contenu offensif (exploit fonctionnel, script de scan de masse, bypass de protection).
- [ ] Cohérent avec le format des fichiers `schemas/*.json` si une structure de données est modifiée.
