---
id: css-injection
category: frontend
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: medium
languages: [js, ts]
---

# Injection CSS

## Description
L'injection CSS survient quand une application permet à un attaquant d'injecter des règles ou des valeurs CSS arbitraires dans une page, généralement via des styles dynamiques construits à partir d'entrées utilisateur non neutralisées. Bien que perçue comme moins critique qu'une injection JavaScript, une injection CSS peut être exploitée pour exfiltrer des données sensibles via des sélecteurs d'attributs (ex: lecture de la valeur d'un champ caché), modifier l'apparence de la page pour tromper l'utilisateur, ou masquer/afficher des éléments à des fins de phishing visuel.

## Où ça apparaît typiquement
- Fonctionnalités de personnalisation de thème permettant à l'utilisateur de fournir du CSS libre, injecté tel quel dans une balise `<style>`.
- Génération dynamique d'attributs `style="..."` à partir de données utilisateur (couleur, taille, position) sans validation de format.
- Templates rendant des valeurs utilisateur directement dans des propriétés CSS (`background: url(...)`, `content: attr(...)`) sans échappement.
- Éditeurs WYSIWYG ou markdown autorisant des balises `<style>` ou attributs `style` dans le contenu généré.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Insertion d'une valeur utilisateur directement dans un bloc `<style>` ou un attribut `style` sans validation de type/format strict.
- Absence de liste blanche de propriétés CSS autorisées dans les fonctionnalités de personnalisation de thème.
- Utilisation de `innerHTML` pour injecter des balises `<style>` construites dynamiquement.
- Absence de sanitisation dédiée au CSS (bibliothèque de nettoyage CSS) avant rendu dans le DOM.

## Remédiation
- Ne jamais interpoler d'entrée utilisateur brute dans du CSS ; valider strictement le format attendu (ex: couleur hexadécimale via une regex stricte).
- Utiliser une bibliothèque de sanitisation CSS dédiée si le CSS libre est un besoin métier réel.
- Préférer des classes CSS prédéfinies plutôt que des styles inline générés dynamiquement.
- Appliquer une CSP avec `style-src` restrictif en complément.
- Voir `rules/remediation/css-injection.md` pour les diffs par langage/framework.

## Exemple avant/après
Voir `examples/js/css-injection/`.

## Références
- OWASP Cheat Sheet: Cross Site Scripting Prevention Cheat Sheet
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
