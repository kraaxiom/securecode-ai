# Remédiation — Expression Language (EL) Injection

## Principe
Ne jamais construire une expression EL/OGNL/SpEL (ou équivalent : template Twig, EJS, Jinja2) à partir d'une entrée utilisateur concaténée. Utiliser des expressions statiques et transmettre les données utilisateur uniquement comme valeurs de paramètres liés, jamais comme fragment de code évalué.

## Java (Spring EL)
```java
// Avant — vulnérable
String userExpr = request.getParameter("expr");
ExpressionParser parser = new SpelExpressionParser();
Expression exp = parser.parseExpression(userExpr);
Object result = exp.getValue(context);

// Après — sécurisé
// L'expression est statique et définie par le développeur ; seule la donnée
// utilisateur est injectée comme variable du contexte d'évaluation.
ExpressionParser parser = new SpelExpressionParser();
Expression exp = parser.parseExpression("#input.length()");
StandardEvaluationContext context = new StandardEvaluationContext();
context.setVariable("input", request.getParameter("value"));
Object result = exp.getValue(context);
```

## PHP (Twig)
```php
// Avant — vulnérable
$tplSource = "Bonjour " . $_GET['nom'];
$template = $twig->createTemplate($tplSource);
echo $template->render();

// Après — sécurisé
// Le nom utilisateur est passé comme variable de rendu, jamais concaténé
// dans le code source du template.
$template = $twig->createTemplate("Bonjour {{ nom }}");
echo $template->render(['nom' => $_GET['nom']]);
```

## JavaScript / Node.js (EJS)
```js
// Avant — vulnérable
const tplSource = `<p>Bonjour ${req.query.nom}</p>`;
res.send(ejs.render(tplSource));

// Après — sécurisé
// Le template est une chaîne statique ; la donnée utilisateur est transmise
// comme variable de rendu et échappée automatiquement par le moteur.
const tplSource = "<p>Bonjour <%= nom %></p>";
res.send(ejs.render(tplSource, { nom: req.query.nom }));
```

## Python (Jinja2)
```python
# Avant — vulnérable
from jinja2 import Template
tpl_source = "Bonjour " + request.args.get("nom")
template = Template(tpl_source)
output = template.render()

# Après — sécurisé
# Le template est statique ; la donnée utilisateur est passée en variable
# de contexte, jamais insérée dans la source du template.
template = Template("Bonjour {{ nom }}")
output = template.render(nom=request.args.get("nom"))
```

## Checklist de vérification post-patch
- [ ] Aucune expression EL/OGNL/SpEL ou source de template n'est construite par concaténation d'une entrée utilisateur.
- [ ] Les données utilisateur ne sont transmises qu'en tant que variables du contexte d'évaluation/rendu, jamais comme fragment de code.
- [ ] Un test de non-régression confirme que les cas d'usage légitimes (rendu normal) fonctionnent toujours.
- [ ] Un test confirme qu'une entrée contenant des caractères spéciaux d'expression (`#`, `${...}`, `{{...}}`) n'est pas évaluée comme code.
- [ ] Le moteur d'évaluation/template utilisé est en version à jour et, si possible, configuré en mode sandbox restreint.
