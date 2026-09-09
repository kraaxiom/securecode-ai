# Remédiation — Server-Side Template Injection (SSTI)

## Principe
Ne jamais construire dynamiquement le texte d'un template à partir d'une entrée utilisateur. Les données utilisateur doivent toujours être passées comme variables de contexte au moteur de rendu, jamais concaténées à la structure même du template avant compilation.

## PHP (Twig)
```php
// Avant — vulnérable
$template = "Bonjour " . $_GET['name'] . ", bienvenue !";
echo $twig->createTemplate($template)->render();

// Après — sécurisé
echo $twig->render('welcome.html.twig', ['name' => $_GET['name']]);
```

## JavaScript / Node.js (Handlebars)
```js
// Avant — vulnérable
const source = `<p>Bonjour ${req.query.name}</p>`;
const template = Handlebars.compile(source);
res.send(template({}));

// Après — sécurisé
const template = Handlebars.compile('<p>Bonjour {{name}}</p>');
res.send(template({ name: req.query.name }));
```

## Python (Jinja2 / Flask)
```python
# Avant — vulnérable
from flask import render_template_string, request
return render_template_string("Bonjour " + request.args.get("name"))

# Après — sécurisé
from flask import render_template_string, request
return render_template_string("Bonjour {{ name }}", name=request.args.get("name"))
```

## Java (FreeMarker)
```java
// Avant — vulnérable
String templateText = "Bonjour " + request.getParameter("name");
Template t = new Template("dyn", new StringReader(templateText), cfg);
t.process(model, out);

// Après — sécurisé
Template t = cfg.getTemplate("welcome.ftl");
model.put("name", request.getParameter("name"));
t.process(model, out);
```

## Checklist de vérification post-patch
- [ ] Le texte du template n'est plus construit par concaténation d'une entrée utilisateur.
- [ ] Toute donnée utilisateur transite uniquement par le contexte/modèle de variables, jamais par la structure du template.
- [ ] Le moteur de template est utilisé en mode "sandbox" si un rendu dynamique reste indispensable.
- [ ] Un test de non-régression confirme que le rendu légitime (variables normales) fonctionne toujours.
- [ ] Une revue confirme l'absence d'autres appels `render_template_string`/`compile`/équivalent construits dynamiquement dans le reste du code.
