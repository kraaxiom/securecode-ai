---
id: imagetragick
category: upload
cwe: CWE-78
owasp: A03:2021-Injection
severity_default: critical
languages: [php, python, java, csharp, go]
---

# ImageTragick (injection de commande via bibliothèque de traitement d'image)

## Description
ImageTragick désigne une classe de vulnérabilités (illustrée par CVE-2016-3714) affectant des bibliothèques de traitement d'image comme ImageMagick, où le traitement d'un fichier image spécialement construit provoque l'exécution de commandes système arbitraires. La cause profonde est l'utilisation par la bibliothèque de "delegates" (programmes externes invoqués via une construction de commande shell) pour gérer certains formats, combinée à une neutralisation insuffisante de champs contrôlés par l'attaquant (comme le nom de fichier ou des directives internes au format MVG/SVG) dans cette construction de commande. Toute application qui délègue le traitement d'image téléversée à une bibliothèque native sans isolation est potentiellement exposée à cette classe de risque.

## Où ça apparaît typiquement
- Redimensionnement, conversion ou génération de miniatures d'images uploadées via ImageMagick, GraphicsMagick ou des wrappers (`Imagick` en PHP, `Pillow` avec délégation externe, `ImageMagick.NET`).
- Traitement de formats délégués à des programmes externes (ex: conversion PDF, PS, certains formats vectoriels) plutôt que décodés nativement.
- Pipelines acceptant un nom de fichier ou des métadonnées utilisateur réinjectés dans une commande construite par la bibliothèque de traitement d'image.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Version de bibliothèque de traitement d'image obsolète ou non patchée pour les CVE connues de cette classe (à vérifier via l'inventaire de dépendances).
- Absence de configuration de "policy" restrictive (ex: `policy.xml` d'ImageMagick) désactivant les coders/delegates non nécessaires.
- Traitement d'image effectué dans le même processus/privilèges que l'application principale, sans sandboxing ni isolation (conteneur dédié, utilisateur restreint).
- Nom de fichier ou métadonnées utilisateur transmis sans neutralisation à un appel de bibliothèque susceptible de les propager vers un delegate externe.

## Remédiation
- Maintenir la bibliothèque de traitement d'image à jour et appliquer les correctifs de sécurité dès leur publication.
- Configurer une policy stricte désactivant les coders et delegates non indispensables (ex: désactiver MVG, MSL, EPHEMERAL, URL, HTTPS dans ImageMagick).
- Exécuter le traitement d'image dans un environnement isolé (conteneur sans accès réseau/fichier sensible, utilisateur à privilèges minimaux).
- Ne jamais transmettre de valeurs contrôlées par l'utilisateur dans un nom de fichier ou une commande utilisée par la bibliothèque.
- Voir `rules/remediation/imagetragick.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/imagetragick/` (et équivalents pour les autres langages listés).

## Références
- OWASP Cheat Sheet: File Upload Cheat Sheet
- CWE-78: Improper Neutralization of Special Elements used in an OS Command ('OS Command Injection')
- CVE-2016-3714 (ImageTragick)
