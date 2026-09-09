# XSS via SVG — CWE-79

La version vulnérable stocke et sert les fichiers SVG uploadés par les utilisateurs tels quels, avec `Content-Type: image/svg+xml`, en affichage inline sur le même domaine que l'application. Or le format SVG est un dialecte XML capable d'embarquer `<script>`, des gestionnaires d'événements (`onload`, `onerror`) ou des liens `javascript:`, exécutés par le navigateur dans le contexte de l'application — y compris avec accès aux cookies de session.

La version corrigée sanitise le contenu SVG avant stockage (suppression de `<script>`, des attributs `on*`, des schémas `javascript:` et des balises `<foreignObject>`), et sert systématiquement le fichier en `Content-Disposition: attachment` pour forcer le téléchargement plutôt que le rendu inline, avec un en-tête `Content-Security-Policy: script-src 'none'` en complément. Servir ce type de contenu depuis un sous-domaine dédié sans cookies reste la meilleure isolation possible.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation).
