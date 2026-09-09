# Remédiation — Server-Side Includes (SSI) Injection

## Principe
Désactiver SSI (en particulier l'exécution de commande via SSI) sur les répertoires où du contenu utilisateur est stocké. Échapper ou supprimer les séquences `<!--#` dans tout contenu utilisateur inséré dans une page interprétée par le serveur, et préférer un moteur de templates applicatif à SSI pour du contenu dynamique.

## PHP (génération de page servie en .shtml)
```php
// Avant — vulnérable : le commentaire utilisateur est écrit tel quel dans un fichier .shtml
$comment = $_POST['comment'];
file_put_contents('comments.shtml', "<p>$comment</p>", FILE_APPEND);

// Après — sécurisé : neutraliser la séquence de directive SSI avant écriture
$comment = $_POST['comment'];
$comment = htmlspecialchars($comment, ENT_QUOTES, 'UTF-8');
$comment = str_replace('<!--#', '&lt;!--#', $comment); // neutralise toute directive résiduelle
file_put_contents('comments.shtml', "<p>$comment</p>", FILE_APPEND);
// Préférable : ne jamais écrire dans un fichier interprété par SSI, utiliser un moteur de templates
```

## JavaScript / Node.js (équivalent : rendu de fragments dans une page servie statiquement avec SSI activé côté serveur web)
```js
// Avant — vulnérable
const comment = req.body.comment;
fs.appendFileSync('public/comments.shtml', `<p>${comment}</p>`);

// Après — sécurisé
const escapeHtml = require('escape-html');

function stripSsiDirectives(input) {
  return input.replace(/<!--#/g, '&lt;!--#');
}

const comment = stripSsiDirectives(escapeHtml(req.body.comment));
fs.appendFileSync('public/comments.shtml', `<p>${comment}</p>`);
// Préférable : générer la page via un moteur de templates (EJS, Handlebars) et désactiver SSI côté serveur web
```

## Python (équivalent : génération de fichier .shtml servi par Apache mod_include)
```python
# Avant — vulnérable
comment = request.form['comment']
with open('comments.shtml', 'a') as f:
    f.write(f"<p>{comment}</p>")

# Après — sécurisé
import html

def strip_ssi_directives(value: str) -> str:
    return value.replace('<!--#', '&lt;!--#')

comment = strip_ssi_directives(html.escape(request.form['comment']))
with open('comments.shtml', 'a') as f:
    f.write(f"<p>{comment}</p>")
# Préférable : rendre le contenu via un moteur de templates (Jinja2) plutôt que via un fichier SSI
```

## Checklist de vérification post-patch
- [ ] SSI (et en particulier l'exécution de commande via SSI, directive `exec`) est désactivé sur tout répertoire contenant du contenu utilisateur.
- [ ] Toute séquence `<!--#` dans un contenu utilisateur est échappée ou supprimée avant écriture dans un fichier interprété par SSI.
- [ ] Le contenu dynamique est de préférence généré via un moteur de templates applicatif plutôt que via des fichiers `.shtml`.
- [ ] Un test confirme qu'un commentaire contenant `<!--#` ne déclenche aucune directive SSI une fois rendu.
