---
id: ssti
category: injections
cwe: CWE-1336
owasp: A03:2021-Injection
severity_default: critical
languages: [php, js, python, java]
---

# Server-Side Template Injection (SSTI)

## Description
L'injection de template côté serveur survient lorsqu'une entrée utilisateur non neutralisée est intégrée directement dans un template avant sa compilation/rendu par le moteur de templates (Jinja2, Twig, FreeMarker, Handlebars côté serveur, etc.), plutôt que d'être passée comme simple variable de contexte. Le moteur interprète alors la syntaxe injectée comme du code de template légitime, ce qui peut mener à une divulgation d'information, voire à une exécution de code arbitraire selon la puissance du langage de template.

## Où ça apparaît typiquement
- Génération de pages ou d'emails personnalisés où une partie du template lui-même (pas seulement les variables) provient de l'utilisateur.
- Fonctionnalités de personnalisation avancée (thèmes, messages templatisables) permettant à un utilisateur de saisir une syntaxe de template.
- Concaténation d'une chaîne utilisateur directement dans le texte du template avant l'appel à `render_template_string` ou équivalent, au lieu de la passer en contexte.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Appel à une fonction de rendu de template avec une chaîne de template construite dynamiquement à partir d'une entrée utilisateur (`render_template_string(user_input)`).
- Absence de séparation claire entre le template statique (contrôlé par le développeur) et les données utilisateur (passées en contexte).
- Utilisation d'un moteur de template dont le "sandboxing" n'est pas activé alors que du contenu utilisateur peut atteindre le template.

## Remédiation
- Ne jamais construire dynamiquement le texte du template à partir d'entrée utilisateur; toujours passer les données utilisateur comme variables de contexte, jamais comme structure de template.
- Si un rendu dynamique est indispensable, utiliser un moteur en mode sandbox avec une liste blanche stricte de fonctions/filtres autorisés.
- Séparer clairement templates de confiance (fichiers statiques versionnés) et données utilisateur.
- Voir `rules/remediation/ssti.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/python-flask/ssti/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Server Side Template Injection Prevention (portail OWASP)
- CWE-1336: Improper Neutralization of Special Elements Used in a Template Engine
