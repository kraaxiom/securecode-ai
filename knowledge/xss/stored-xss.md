---
id: stored-xss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: critical
languages: [php, js, python, java, csharp, go]
---

# Stored Cross-Site Scripting (Stored XSS)

## Description
Le XSS stocké survient lorsqu'une entrée utilisateur malveillante est persistée côté serveur (base de données, fichier, cache) puis réaffichée à d'autres utilisateurs sans encodage contextuel adapté. Il est considéré comme plus critique que le XSS réfléchi car il ne nécessite aucune interaction spécifique de la victime (pas de lien piégé à faire cliquer) et peut affecter un grand nombre d'utilisateurs consultant la page concernée, y compris des administrateurs.

## Où ça apparaît typiquement
- Commentaires, avis, messages de forum ou de chat affichés à d'autres utilisateurs.
- Profils utilisateurs (nom, biographie) affichés sur des pages consultées par des tiers.
- Champs de configuration ou de personnalisation stockés puis rendus dans une interface partagée (nom d'équipe, titre de projet).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Donnée persistée en base sans validation de format, puis réaffichée via un sink qui interprète du HTML sans encodage.
- Désactivation explicite de l'échappement automatique du moteur de templates sur une donnée provenant du stockage (`|safe`, `{!! !!}`, `dangerouslySetInnerHTML`).
- Absence de sanitisation HTML lors de l'enregistrement ou de l'affichage de contenu riche fourni par l'utilisateur (éditeur WYSIWYG).

## Remédiation
- Appliquer un encodage de sortie contextuel systématique lors de l'affichage de toute donnée persistée d'origine utilisateur.
- Pour le contenu riche (HTML autorisé), utiliser une bibliothèque de sanitisation HTML reconnue avant stockage et/ou avant affichage.
- Mettre en place une Content Security Policy stricte en défense en profondeur.
- Voir `rules/remediation/stored-xss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/stored-xss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Cross Site Scripting Prevention
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
