# Remédiation — Reflected Cross-Site Scripting

## Principe
Utiliser l'échappement automatique du moteur de templates pour toute donnée issue de la requête HTTP, et ne le désactiver que pour du contenu strictement contrôlé par le développeur. Appliquer un encodage contextuel adapté (HTML, attribut, JavaScript, URL) selon l'emplacement d'insertion.

## PHP
```php
// Avant — vulnérable
$term = $_GET['q'];
echo "Résultats pour : " . $term;

// Après — sécurisé
$term = $_GET['q'];
echo "Résultats pour : " . htmlspecialchars($term, ENT_QUOTES, 'UTF-8');
```

```php
// Avant — vulnérable — Blade avec échappement désactivé
{!! request('q') !!}

// Après — sécurisé — échappement automatique Blade
{{ request('q') }}
```

## JavaScript / Node.js (Express)
```js
// Avant — vulnérable
app.get('/search', (req, res) => {
  res.send(`<h1>Résultats pour : ${req.query.q}</h1>`);
});

// Après — sécurisé — moteur de templates avec échappement automatique (EJS)
app.get('/search', (req, res) => {
  res.render('results', { term: req.query.q }); // <h1>Résultats pour : <%= term %></h1>
});
```

```js
// Avant — vulnérable — React avec sink dangereux
function SearchResult({ term }) {
  return <div dangerouslySetInnerHTML={{ __html: term }} />;
}

// Après — sécurisé — rendu texte natif, pas d'interprétation HTML
function SearchResult({ term }) {
  return <div>{term}</div>;
}
```

## Python (Flask / Jinja2)
```python
# Avant — vulnérable — échappement désactivé explicitement
from flask import request, render_template_string

template = "Résultats pour : {{ q|safe }}"
return render_template_string(template, q=request.args.get('q'))

# Après — sécurisé — échappement automatique Jinja2 conservé
template = "Résultats pour : {{ q }}"
return render_template_string(template, q=request.args.get('q'))
```

## Checklist de vérification post-patch
- [ ] Toute valeur issue de la requête HTTP (query string, en-tête, corps) affichée dans la réponse passe par l'échappement automatique du moteur de templates.
- [ ] Aucune désactivation explicite de l'échappement (`|safe`, `{!! !!}`, `dangerouslySetInnerHTML`) ne reste sur une donnée dérivée de la requête.
- [ ] L'encodage contextuel utilisé correspond à l'emplacement d'insertion (HTML, attribut, URL, JavaScript).
- [ ] Un test de non-régression confirme que la recherche/le formulaire légitime s'affiche toujours correctement.
- [ ] Une Content Security Policy stricte est en place en défense en profondeur.
