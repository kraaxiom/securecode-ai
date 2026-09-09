# markdown-xss (CWE-79)

La version vulnerable convertit le Markdown utilisateur en HTML avec l'option `html.WithUnsafe()` de goldmark, qui laisse passer tel quel le HTML brut (dont `<script>`) present dans la source. La version corrigee retire cette option pour beneficier de l'echappement par defaut du HTML brut, puis applique en defense en profondeur une sanitisation `bluemonday` (allowlist stricte) sur la sortie HTML.
