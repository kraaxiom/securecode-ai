# Remédiation — XPath Injection

## Principe
Utiliser des expressions XPath paramétrées (variables liées) plutôt que la concaténation de chaînes. Éviter d'utiliser XPath comme mécanisme d'authentification.

## PHP (DOMXPath)
```php
// Avant — vulnérable
$user = $_POST['user'];
$pass = $_POST['pass'];
$nodes = $xpath->query("//user[username='$user' and password='$pass']");

// Après — sécurisé
// DOMXPath ne supporte pas nativement les variables liées : valider/échapper strictement
// et, idéalement, ne pas utiliser XPath pour l'authentification (préférer une base avec hash de mot de passe).
$user = $xpath->quote($_POST['user']);
$pass = $xpath->quote($_POST['pass']);
$nodes = $xpath->query("//user[username=$user and password=$pass]");
```

## JavaScript / Node.js (xpath)
```js
// Avant — vulnérable
const expr = `//user[username='${req.body.user}' and password='${req.body.pass}']`;
const nodes = xpath.select(expr, doc);

// Après — sécurisé
// Échapper strictement les apostrophes ou, mieux, remplacer XPath par une source de données
// paramétrable ; à défaut, valider un format strict (ex: alphanumérique) avant insertion.
function xpathLiteral(value) {
  if (!value.includes("'")) return `'${value}'`;
  return "concat('" + value.split("'").join("', \"'\", '") + "')";
}
const expr = `//user[username=${xpathLiteral(req.body.user)} and password=${xpathLiteral(req.body.pass)}]`;
```

## Python (lxml.etree avec XPath paramétré)
```python
# Avant — vulnérable
expr = f"//user[username='{user}' and password='{password}']"
nodes = tree.xpath(expr)

# Après — sécurisé
# lxml supporte les variables XPath liées via des paramètres nommés
nodes = tree.xpath(
    "//user[username=$u and password=$p]",
    u=user,
    p=password,
)
```

## Java (XPathExpression avec variables liées)
```java
// Avant — vulnérable
String expr = "//user[username='" + user + "' and password='" + password + "']";
XPathExpression xpe = xpath.compile(expr);

// Après — sécurisé
xpath.setXPathVariableResolver(variableName -> {
    if ("user".equals(variableName.getLocalPart())) return user;
    if ("pass".equals(variableName.getLocalPart())) return password;
    return null;
});
XPathExpression xpe = xpath.compile("//user[username=$user and password=$pass]");
```

## Checklist de vérification post-patch
- [ ] L'expression XPath n'est plus construite par concaténation directe d'une entrée utilisateur.
- [ ] Des variables liées ou une fonction d'échappement dédiée sont utilisées pour toute valeur insérée.
- [ ] XPath n'est plus utilisé comme mécanisme d'authentification, ou son usage est justifié et documenté.
- [ ] Un test confirme qu'une entrée contenant une apostrophe ou un opérateur XPath ne modifie plus la sélection de nœuds.
