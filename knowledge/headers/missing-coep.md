---
id: missing-coep
category: headers
cwe: CWE-1021
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: [php, js, python, java, csharp, go, rust]
---

# Cross-Origin-Embedder-Policy (COEP) manquant

## Description
L'absence de l'en-tête `Cross-Origin-Embedder-Policy` empêche l'application de bénéficier de l'isolation cross-origin complète du navigateur. Sans COEP (associé à COOP), la page ne peut pas garantir que toutes les ressources qu'elle charge (images, scripts, iframes) proviennent d'origines ayant explicitement consenti à être embarquées, ce qui limite les protections contre certaines classes d'attaques par canal auxiliaire et empêche l'usage de fonctionnalités nécessitant l'isolation (comme `SharedArrayBuffer`).

## Où ça apparaît typiquement
- Applications web sensibles (finance, santé) traitant des données confidentielles en mémoire côté client sans isolation cross-origin.
- Applications utilisant des fonctionnalités nécessitant `SharedArrayBuffer` (traitement multimédia, WebAssembly threadé) sans avoir configuré COEP/COOP requis pour les activer.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `Cross-Origin-Embedder-Policy` absent des réponses HTTP.
- Ressources tierces chargées sans en-tête `Cross-Origin-Resource-Policy` compatible, ce qui bloquerait leur chargement si COEP était activé (signe qu'une migration n'a pas été faite).

## Remédiation
- Ajouter `Cross-Origin-Embedder-Policy: require-corp` (ou `credentialless` selon compatibilité) en complément de `Cross-Origin-Opener-Policy: same-origin`.
- S'assurer que toutes les ressources tierces embarquées définissent `Cross-Origin-Resource-Policy` ou sont servies avec CORS approprié avant d'activer COEP en production.
- Tester en environnement de staging car COEP peut bloquer le chargement de ressources tierces non conformes.
- Voir `rules/remediation/missing-coep.md`.

## Exemple avant/après
Voir `examples/js/missing-coep/`.

## Références
- OWASP Secure Headers Project
- MDN Web Docs: Cross-Origin-Embedder-Policy
- CWE-1021: Improper Restriction of Rendered UI Layers or Frames
