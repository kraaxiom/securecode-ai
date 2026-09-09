# Universal XSS (UXSS) — CWE-79

La version vulnérable embarque un widget tiers dans une `<iframe>` sans aucune restriction (`sandbox`, `frame-src`) et sans en-têtes CSP / `X-Frame-Options` sur ses propres pages. Une faille dans le moteur de rendu du navigateur ou dans le widget lui-même peut alors permettre à un contenu contrôlé par l'attaquant d'échapper à l'isolation attendue et de s'exécuter dans le contexte d'origine de l'application (contournement des protections normales du même-origine, d'où le terme "universal").

La version corrigée applique l'attribut `sandbox="allow-scripts"` sur l'iframe (sans `allow-same-origin`, dont la combinaison avec `allow-scripts` permettrait justement une évasion vers l'origine parente), et envoie une politique `Content-Security-Policy` restreignant `frame-src` au domaine de confiance et fixant `frame-ancestors 'none'` pour empêcher que l'application elle-même soit embarquée à son insu.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation), variante UXSS liée à l'isolation d'origine des contenus embarqués.
