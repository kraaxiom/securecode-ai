# Blind XSS — CWE-79

La version vulnérable enregistre un message de ticket support saisi par l'utilisateur puis le restitue **sans échappement** (`th:utext`) dans la page consultée par un administrateur du back-office. L'attaquant n'observe jamais directement l'exécution du script : celui-ci se déclenche plus tard, dans un contexte souvent privilégié (session admin), ce qui en fait une variante particulièrement dangereuse du XSS stocké.

La version corrigée conserve la donnée brute en base (aucune perte d'information) mais neutralise systématiquement le contenu à l'affichage : le template utilise `th:text` (échappement automatique de Thymeleaf) et, en défense en profondeur, la valeur est explicitement passée dans `Encode.forHtml()` (OWASP Java Encoder) avant d'être exposée au modèle. Une validation de longueur est également ajoutée en entrée, et une politique CSP stricte est recommandée sur les pages d'administration pour limiter l'impact d'un éventuel contournement.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation).
