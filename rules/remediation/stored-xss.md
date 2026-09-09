# Remédiation — Stored Cross-Site Scripting

## Principe
Appliquer un encodage de sortie contextuel systématique lors de l'affichage de toute donnée persistée d'origine utilisateur. Pour le contenu riche où du HTML doit être autorisé, utiliser une bibliothèque de sanitisation HTML reconnue avant stockage et/ou avant affichage.

## PHP
```php
// Avant — vulnérable
$comment = $db->query("SELECT body FROM comments WHERE id = ?", [$id])->fetch();
echo $comment['body'];

// Après — sécurisé
$comment = $db->query("SELECT body FROM comments WHERE id = ?", [$id])->fetch();
echo htmlspecialchars($comment['body'], ENT_QUOTES, 'UTF-8');
```

```php
// Avant — vulnérable — Laravel Blade avec échappement désactivé sur une donnée stockée
{!! $comment->body !!}

// Après — sécurisé — échappement automatique + sanitisation en amont si HTML riche requis
{{ $comment->body }}

// Si du HTML riche doit être conservé (éditeur WYSIWYG), sanitiser avant stockage :
$clean = (new HTMLPurifier())->purify($request->input('body'));
Comment::create(['body' => $clean]);
```

## JavaScript / Node.js
```js
// Avant — vulnérable
app.get('/profile/:id', async (req, res) => {
  const user = await User.findById(req.params.id);
  res.send(`<div class="bio">${user.bio}</div>`);
});

// Après — sécurisé — moteur de templates avec échappement automatique
app.get('/profile/:id', async (req, res) => {
  const user = await User.findById(req.params.id);
  res.render('profile', { bio: user.bio }); // <div class="bio"><%= bio %></div>
});
```

```js
// Avant — vulnérable — React
function Profile({ bio }) {
  return <div dangerouslySetInnerHTML={{ __html: bio }} />;
}

// Après — sécurisé — sanitisation explicite si HTML riche nécessaire, sinon rendu texte
import DOMPurify from 'dompurify';

function Profile({ bio }) {
  return <div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(bio) }} />;
}
```

## Python (Django)
```python
# Avant — vulnérable — filtre safe appliqué à une donnée stockée
{{ comment.body|safe }}

# Après — sécurisé — échappement automatique Django conservé
{{ comment.body }}

# Si HTML riche nécessaire, sanitiser avant stockage :
import bleach

clean_body = bleach.clean(request.POST['body'], tags=['p', 'b', 'i', 'a'], strip=True)
Comment.objects.create(body=clean_body)
```

## Checklist de vérification post-patch
- [ ] Toute donnée persistée d'origine utilisateur est encodée contextuellement à l'affichage (échappement automatique du moteur de templates conservé).
- [ ] Aucune désactivation explicite de l'échappement (`|safe`, `{!! !!}`, `dangerouslySetInnerHTML`) ne reste sur une donnée provenant du stockage.
- [ ] Le contenu riche autorisé (éditeur WYSIWYG) est sanitisé via une bibliothèque HTML reconnue avant stockage.
- [ ] Un test de non-régression confirme que le contenu stocké légitime (accents, ponctuation, mise en forme autorisée) s'affiche toujours correctement.
- [ ] Une Content Security Policy stricte est en place en défense en profondeur.
