---
id: missing-coop
category: headers
cwe: CWE-1021
owasp: A05:2021-Security Misconfiguration
severity_default: low
languages: [php, js, python, java, csharp, go, rust]
---

# Cross-Origin-Opener-Policy (COOP) manquant

## Description
L'absence de l'en-tête `Cross-Origin-Opener-Policy` laisse la fenêtre de la page partager son groupe de contexte de navigation avec des fenêtres ouvertes par des sites tiers (via `window.open` ou des liens `target="_blank"`). Un site malveillant peut alors conserver une référence `window.opener` vers la page et, dans certains scénarios, manipuler sa navigation (attaques de type "tabnabbing") ou tirer parti d'attaques par canal auxiliaire exploitant l'isolation processus (type Spectre) si COOP/COEP ne sont pas activés ensemble.

## Où ça apparaît typiquement
- Applications ouvrant ou étant ouvertes par des fenêtres tierces (paiement, SSO, partage social) sans isolation d'origine.
- Sites n'ayant jamais configuré les en-têtes d'isolation cross-origin modernes (COOP/COEP/CORP).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- En-tête `Cross-Origin-Opener-Policy` absent des réponses HTTP.
- Liens `target="_blank"` vers des domaines externes sans attribut `rel="noopener"`.

## Remédiation
- Ajouter `Cross-Origin-Opener-Policy: same-origin` (ou `same-origin-allow-popups` si des popups tiers légitimes sont nécessaires).
- Ajouter systématiquement `rel="noopener noreferrer"` sur les liens `target="_blank"` vers des origines externes.
- Combiner avec `Cross-Origin-Embedder-Policy` pour bénéficier de l'isolation cross-origin complète si l'application en a besoin (ex: usage de `SharedArrayBuffer`).
- Voir `rules/remediation/missing-coop.md`.

## Exemple avant/après
Voir `examples/js/missing-coop/`.

## Références
- OWASP Secure Headers Project
- MDN Web Docs: Cross-Origin-Opener-Policy
- CWE-1021: Improper Restriction of Rendered UI Layers or Frames
