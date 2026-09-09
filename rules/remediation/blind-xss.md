# Remédiation — Blind Cross-Site Scripting (Blind XSS)

## Principe
Appliquer un encodage de sortie contextuel systématique sur toute donnée affichée, y compris dans les interfaces internes/admin — traiter toute donnée provenant de l'extérieur du périmètre de confiance comme non fiable, quel que soit l'écran où elle est finalement affichée.

## PHP (back-office affichant des tickets support)
```php
// Avant — vulnérable
echo "<div class='ticket-message'>" . $ticket['message'] . "</div>";

// Après — sécurisé
echo "<div class='ticket-message'>" . htmlspecialchars($ticket['message'], ENT_QUOTES, 'UTF-8') . "</div>";
```

## JavaScript / Node.js (dashboard admin, template EJS)
```js
// Avant — vulnérable (sortie non échappée)
res.send(`<div class="log-entry">${logEntry.userAgent}</div>`);

// Après — sécurisé (échappement automatique via <%= %> en EJS, jamais <%- %> sur donnée externe)
// template.ejs : <div class="log-entry"><%= logEntry.userAgent %></div>
res.render('logs', { logEntry });
```

## Python (Django admin / vue interne)
```python
# Avant — vulnérable
from django.utils.safestring import mark_safe
return mark_safe(f"<div class='comment'>{comment.body}</div>")

# Après — sécurisé
from django.utils.html import escape
return f"<div class='comment'>{escape(comment.body)}</div>"
# Ou, mieux : laisser le moteur de template Django échapper automatiquement (ne pas utiliser mark_safe/|safe)
```

## Checklist de vérification post-patch
- [ ] Toute donnée affichée dans une interface d'administration ou un outil interne passe par un encodage de sortie contextuel.
- [ ] Aucun mécanisme de contournement d'échappement (`mark_safe`, `|safe`, `dangerouslySetInnerHTML`, `<%- %>`) n'est utilisé sur une donnée externe non fiable.
- [ ] Une Content Security Policy stricte est en place sur les interfaces d'administration.
- [ ] Un test confirme qu'un contenu contenant des balises HTML soumis via un canal externe (formulaire de contact, ticket, User-Agent) s'affiche comme texte inerte dans le back-office.
- [ ] Un mécanisme de callback/canary est utilisé lors des tests de sécurité pour détecter une éventuelle exécution différée dans un contexte non observable directement.
