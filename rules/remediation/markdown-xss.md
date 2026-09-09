# Remédiation — Markdown-based Cross-Site Scripting

## Principe
Désactiver le support du HTML brut dans le moteur de rendu Markdown pour tout contenu utilisateur non fiable, et faire systématiquement passer le HTML généré par une bibliothèque de sanitisation reconnue avant insertion dans le DOM ou la réponse HTTP. Restreindre les schémas d'URL autorisés dans les liens et images.

## JavaScript / Node.js (markdown-it, marked)
```js
// Avant — vulnérable — HTML brut autorisé sur du contenu utilisateur, inséré sans sanitisation
const md = require('markdown-it')({ html: true });
element.innerHTML = md.render(userMarkdown);

// Après — sécurisé — HTML brut désactivé + sanitisation du HTML généré
const md = require('markdown-it')({ html: false, linkify: true });
import DOMPurify from 'dompurify';

const rendered = md.render(userMarkdown);
element.innerHTML = DOMPurify.sanitize(rendered, {
  ALLOWED_URI_REGEXP: /^(?:https?|mailto):/i,
});
```

## PHP (Parsedown / league/commonmark)
```php
// Avant — vulnérable
$parsedown = new Parsedown();
$parsedown->setSafeMode(false);
echo $parsedown->text($userMarkdown);

// Après — sécurisé — mode sûr activé + sanitisation additionnelle du HTML de sortie
$parsedown = new Parsedown();
$parsedown->setSafeMode(true);
$html = $parsedown->text($userMarkdown);
echo HTMLPurifier_instance()->purify($html);
```

## Python (Django/Flask — python-markdown + bleach)
```python
# Avant — vulnérable
import markdown

html = markdown.markdown(user_markdown, extensions=['markdown.extensions.extra'])
return HttpResponse(html)

# Après — sécurisé — sanitisation du HTML généré avec une liste blanche stricte
import markdown
import bleach

html = markdown.markdown(user_markdown)
allowed_tags = ['p', 'em', 'strong', 'a', 'ul', 'ol', 'li', 'code', 'pre', 'blockquote']
allowed_attrs = {'a': ['href', 'title', 'rel']}
allowed_protocols = ['http', 'https', 'mailto']

clean_html = bleach.clean(html, tags=allowed_tags, attributes=allowed_attrs, protocols=allowed_protocols)
return HttpResponse(clean_html)
```

## Checklist de vérification post-patch
- [ ] Le support du HTML brut est désactivé dans le moteur Markdown pour tout contenu utilisateur non fiable.
- [ ] Le HTML produit par le rendu Markdown est systématiquement passé dans une bibliothèque de sanitisation dédiée avant insertion dans le DOM ou la réponse.
- [ ] Les schémas d'URL autorisés dans les liens et images Markdown sont restreints à une liste blanche (`http`, `https`, `mailto`).
- [ ] Un test de non-régression confirme que le Markdown légitime (gras, liens, listes) s'affiche toujours correctement après sanitisation.
- [ ] La sanitisation est appliquée côté serveur avant stockage ET revalidée côté client avant tout rendu dynamique ultérieur.
