# SecureCode AI — Enterprise Skill

Skill de sécurité applicative pour agents de code IA (Claude Code, Cursor, Codex). Il détecte, explique, corrige et note la sécurité d'un projet selon une taxonomie de ~230 vulnérabilités couvrant injections, XSS, SSRF, contrôle d'accès, cryptographie, cloud, conteneurs, CI/CD, supply chain, LLM/IA, logique métier et plus.

## Statut

🚧 En développement. Ce dépôt contient actuellement :
- `SKILL.md` — le fichier d'entrée du skill (fonctionnel, pointe vers les modules ci-dessous)
- `SPEC.md` — le cahier des charges complet à suivre pour finaliser tous les modules (`knowledge/`, `rules/`, `prompts/`, `templates/`, `examples/`, `tests/`, `schemas/`)

Voir `SPEC.md` section 9 (Roadmap) pour l'ordre de construction recommandé.

## Installation (une fois complet)

- **Claude Code / Cowork** : placer le dossier `security-ai-skill/` dans le répertoire de skills du projet ou de l'organisation.
- **Cursor** : référencer `SKILL.md` comme règle/contexte projet (`.cursor/rules` ou équivalent).
- **Codex** : charger `SKILL.md` comme instruction système additionnelle avant une session de revue de code.

## Utilisation typique

```
"Audite ce dossier src/ pour des vulnérabilités de sécurité et donne-moi un rapport OWASP."
"Ce endpoint /api/users/:id est-il vulnérable à un IDOR ?"
"Génère un patch sécurisé pour cette injection SQL, avec un test de non-régression."
"Quel est notre score de sécurité ASVS actuel ?"
```

## Principes

- **Défensif uniquement** : pas d'exploits fonctionnels, pas de contournement de protections.
- **Traçable** : chaque finding référence un CWE, une catégorie OWASP, une sévérité et un fichier de remédiation.
- **Non-intrusif par défaut** : mode audit seul (aucune modification de fichier) sauf confirmation explicite.
- **Complément, pas remplacement** : ne se substitue pas à un audit de sécurité humain ou un test d'intrusion.

## Licence

À définir dans `LICENSE.md` avant publication (usage commercial, redistribution) — voir `SPEC.md` section 10.

## Contribution

Voir `CONTRIBUTING.md` (à créer) pour le format attendu de chaque nouvelle entrée `knowledge/` et `rules/`.
