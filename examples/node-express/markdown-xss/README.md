# XSS via Markdown

`vulnerable.js` active le support du HTML brut (`html: true`) dans le moteur `markdown-it` pour du texte saisi par l'utilisateur, puis renvoie le HTML produit sans aucune sanitisation. `fixed.js` désactive ce support (`html: false`) et fait passer systématiquement le rendu final dans DOMPurify avec une liste blanche de schémas d'URL, avant de le renvoyer au client. Cette classe de vulnérabilité correspond à CWE-79 (Improper Neutralization of Input During Web Page Generation), les moteurs Markdown transformant du texte en HTML actif sans neutralisation par défaut.
