# XSS stocké

`vulnerable.js` affiche la biographie d'un utilisateur, lue depuis la base de données, par concaténation directe dans la réponse HTML sans encodage. `fixed.js` fait passer ce même affichage par un moteur de templates EJS à échappement automatique, garantissant que toute balise HTML précédemment stockée est rendue comme texte inerte. Cette faille correspond à CWE-79 (Improper Neutralization of Input During Web Page Generation), classée critique dans sa forme stockée car elle touche potentiellement tous les visiteurs du profil sans interaction spécifique requise.
