# XML Injection (CWE-91)

Le code vulnérable construit un document XML par concaténation directe du champ `name`, sans échapper les caractères spéciaux (`<`, `>`, `&`), ce qui permet à un attaquant d'altérer la structure du document généré. La correction utilise `DOMDocument` avec `createTextNode()`, qui échappe automatiquement le contenu des nœuds texte, empêchant toute injection de balise (CWE-91 : XML Injection, aka Blind XPath Injection).
