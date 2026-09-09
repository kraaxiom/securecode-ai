---
id: svg-xss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: high
languages: [js, php, python, java, csharp, go]
---

# XSS via SVG

## Description
Le XSS via SVG exploite le fait que le format SVG est un document XML pouvant contenir des éléments actifs (`<script>`, gestionnaires d'événements `onload`, `onerror`, liens `javascript:` dans des attributs `xlink:href`). Lorsqu'un fichier SVG uploadé par un utilisateur est servi directement au navigateur ou affiché inline dans une page, le code qu'il contient peut s'exécuter dans le contexte d'origine de l'application.

## Où ça apparaît typiquement
- Fonctionnalités d'upload d'image acceptant le format SVG et le servant tel quel (avatar, logo, pièce jointe).
- Affichage inline de SVG fourni par l'utilisateur directement dans le DOM (`innerHTML` avec du SVG).
- Éditeurs graphiques ou d'icônes permettant l'import de fichiers SVG externes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Acceptation de fichiers `.svg` dans une fonctionnalité d'upload sans sanitisation du contenu XML avant stockage/affichage.
- Service direct de fichiers SVG uploadés avec un `Content-Type: image/svg+xml` depuis la même origine que l'application (au lieu d'une origine isolée).
- Insertion de SVG utilisateur inline dans le DOM sans passage par un sanitiseur HTML/SVG dédié.

## Remédiation
- Sanitiser le contenu SVG (suppression des balises `script`, gestionnaires d'événements, schémas `javascript:`) avant stockage ou affichage.
- Servir les fichiers uploadés, y compris SVG, depuis une origine séparée (sous-domaine dédié sans cookies) pour limiter l'impact d'une éventuelle exécution.
- Forcer le téléchargement (`Content-Disposition: attachment`) plutôt que l'affichage inline pour les fichiers non indispensables au rendu direct.
- Voir `rules/remediation/svg-xss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/nodejs-express/svg-xss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet / Cross Site Scripting Prevention
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
