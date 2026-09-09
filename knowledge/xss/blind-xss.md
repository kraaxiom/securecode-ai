---
id: blind-xss
category: xss
cwe: CWE-79
owasp: A03:2021-Injection
severity_default: high
languages: [php, js, python, java, csharp, go]
---

# Blind Cross-Site Scripting (Blind XSS)

## Description
Le XSS aveugle est une variante du XSS stocké où la charge injectée s'exécute dans un contexte que l'attaquant ne peut pas observer directement, typiquement un back-office, un panneau d'administration ou un outil interne consultant des données soumises par des utilisateurs externes (formulaires de contact, tickets de support, logs). L'attaquant doit s'appuyer sur des canaux de rappel (callback) externes pour savoir si son injection s'est exécutée, ce qui la rend plus difficile à détecter lors de tests classiques mais tout aussi dangereuse en production, car elle vise souvent des comptes à privilèges élevés.

## Où ça apparaît typiquement
- Champs soumis par des utilisateurs non authentifiés puis affichés uniquement côté administration (formulaires de contact, avis, tickets support, en-têtes User-Agent journalisés puis affichés dans un dashboard).
- Systèmes de journalisation dont les logs sont ensuite visualisés dans une interface web sans échappement.
- Champs de métadonnées (nom de fichier uploadé, référent HTTP) affichés dans des outils internes.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Données provenant d'utilisateurs non authentifiés ou externes affichées dans une interface d'administration sans passage par un encodage contextuel.
- Absence de traitement différencié entre les données destinées à un affichage public et celles destinées à un affichage interne (les deux doivent être échappées).
- Journalisation de champs HTTP bruts (headers, paramètres) réinjectés tels quels dans une page de consultation de logs.

## Remédiation
- Appliquer un encodage de sortie contextuel systématique sur toute donnée affichée, y compris dans les interfaces internes/admin.
- Mettre en place une Content Security Policy stricte sur les interfaces d'administration.
- Traiter toute donnée provenant de l'extérieur du périmètre de confiance comme non fiable, quel que soit l'écran où elle est finalement affichée.
- Voir `rules/remediation/blind-xss.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php-laravel/blind-xss/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Cross Site Scripting Prevention
- CWE-79: Improper Neutralization of Input During Web Page Generation ('Cross-site Scripting')
