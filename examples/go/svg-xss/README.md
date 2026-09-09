# svg-xss (CWE-79)

La version vulnerable sert un fichier SVG uploade tel quel, en affichage inline avec `image/svg+xml` : un SVG peut contenir `<script>`, des gestionnaires d'evenements ou des schemas `javascript:` executes par le navigateur. La version corrigee sanitise le SVG avant stockage avec une politique `bluemonday` dediee (allowlist stricte d'elements/attributs) et sert le fichier avec `Content-Disposition: attachment` pour forcer le telechargement plutot que l'affichage inline.
