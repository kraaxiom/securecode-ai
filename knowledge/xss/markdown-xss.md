---
id: markdown-xss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: medium
languages: [js, php, python, java, csharp, go]
---

# Markdown-based Cross-Site Scripting

## Description
Le XSS via Markdown survient lorsqu'un contenu utilisateur écrit en Markdown est converti en HTML par un moteur de rendu, puis inséré dans une page sans neutralisation supplémentaire. De nombreux moteurs Markdown autorisent par défaut l'inclusion de HTML brut ou de constructions (liens `javascript:`, images avec gestionnaires d'événements) qui, une fois converties, produisent du HTML actif exécutable par le navigateur.

## Où ça apparaît typiquement
- Éditeurs de commentaires, wikis, tickets ou README acceptant du Markdown affiché ensuite à d'autres utilisateurs.
- Rendu côté client d'un contenu Markdown stocké, via une bibliothèque de conversion Markdown → HTML.
- Prévisualisation en temps réel de contenu Markdown avant publication.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Utilisation d'une bibliothèque de rendu Markdown avec le support du HTML brut activé (option souvent nommée `html: true` ou équivalente) sur du contenu utilisateur non fiable.
- Insertion du HTML généré par le moteur Markdown directement via un sink dangereux (`innerHTML`, `dangerouslySetInnerHTML`) sans passage par une étape de sanitisation dédiée.
- Absence de restriction sur les schémas d'URL autorisés dans les liens/images Markdown (`javascript:`, `data:`).

## Remédiation
- Désactiver le support du HTML brut dans le moteur Markdown lorsque le contenu provient d'utilisateurs non fiables.
- Faire passer systématiquement le HTML généré par le rendu Markdown dans une bibliothèque de sanitisation HTML reconnue avant insertion dans le DOM.
- Restreindre les schémas d'URL autorisés dans les liens et images à une liste blanche (`http`, `https`, `mailto`).
- Voir `rules/remediation/markdown-xss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/markdown-xss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Cross Site Scripting Prevention / Sanitization
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
