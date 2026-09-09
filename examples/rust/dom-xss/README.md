# DOM XSS

Le handler `search` injectait le paramètre de requête `q` par concaténation de chaînes directement dans un bloc `<script>` inline, permettant à un attaquant de sortir du contexte JavaScript prévu (CWE-79). La version corrigée sérialise `q` en JSON via `serde_json::json!`, ce qui échappe correctement guillemets et caractères spéciaux avant insertion dans le script. Cette approche évite tout sink dangereux et neutralise la donnée non fiable dès sa mise en forme.
