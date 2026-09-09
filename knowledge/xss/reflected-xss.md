---
id: reflected-xss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Reflected Cross-Site Scripting (Reflected XSS)

## Description
Le XSS réfléchi survient lorsqu'une entrée fournie dans une requête HTTP (paramètre d'URL, en-tête, corps de formulaire) est directement réinsérée dans la réponse HTML générée par le serveur sans encodage contextuel adapté. Contrairement au XSS stocké, la charge n'est pas persistée: elle n'est active que dans la réponse immédiate à la requête contenant l'entrée malveillante, ce qui implique généralement de faire cliquer la victime sur un lien piégé.

## Où ça apparaît typiquement
- Pages de résultats de recherche affichant le terme recherché tel qu'il a été saisi.
- Messages d'erreur reprenant une valeur de paramètre dans le texte affiché.
- Champs de formulaire pré-remplis avec la valeur précédemment soumise, insérés sans encodage dans l'attribut `value`.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Insertion d'un paramètre de requête HTTP directement dans le HTML de réponse sans passer par une fonction d'encodage contextuel du framework.
- Concaténation de chaînes pour générer du HTML incluant une valeur de requête, plutôt que l'utilisation d'un moteur de templates avec échappement automatique.
- Désactivation explicite de l'échappement automatique du moteur de templates pour la sortie concernée (`|safe`, `{!! !!}`, `dangerouslySetInnerHTML`) sur une donnée issue de la requête.

## Remédiation
- Utiliser l'échappement automatique du moteur de templates et ne le désactiver que pour du contenu strictement contrôlé par le développeur.
- Appliquer un encodage contextuel adapté (HTML, attribut, JavaScript, URL) selon l'endroit où la donnée est insérée.
- Mettre en place une Content Security Policy stricte en défense en profondeur.
- Voir `rules/remediation/reflected-xss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/reflected-xss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Cross Site Scripting Prevention
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
