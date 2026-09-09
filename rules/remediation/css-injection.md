# Remédiation — Injection CSS

## Principe
Ne jamais interpoler une entrée utilisateur brute dans du CSS. Valider strictement le format attendu (ex: regex stricte pour une couleur hexadécimale) et préférer des classes CSS prédéfinies plutôt que des styles inline générés dynamiquement.

## PHP
```php
// Avant — vulnérable
$color = $_GET['theme_color'];
echo "<style>body { background: $color; }</style>";

// Après — sécurisé
$color = $_GET['theme_color'];
if (!preg_match('/^#[0-9a-fA-F]{6}$/', $color)) {
    $color = '#ffffff'; // valeur par défaut sûre
}
echo "<style>body { background: " . htmlspecialchars($color, ENT_QUOTES) . "; }</style>";
```

## Node.js / JS (rendu côté client)
```js
// Avant — vulnérable
const userColor = new URLSearchParams(location.search).get('color');
document.getElementById('theme').innerHTML = `<style>body { background: ${userColor}; }</style>`;

// Après — sécurisé
const HEX_COLOR = /^#[0-9a-fA-F]{6}$/;
const userColor = new URLSearchParams(location.search).get('color');
const safeColor = HEX_COLOR.test(userColor) ? userColor : '#ffffff';
document.getElementById('theme').style.setProperty('--bg-color', safeColor);
```

## Python (Flask + Jinja2)
```python
# Avant — vulnérable
theme_color = request.args.get('color')
return render_template_string(f"<style>body {{ background: {theme_color}; }}</style>")

# Après — sécurisé
import re
theme_color = request.args.get('color', '')
if not re.fullmatch(r'#[0-9a-fA-F]{6}', theme_color):
    theme_color = '#ffffff'
return render_template('theme.html', color=theme_color)  # Jinja2 échappe automatiquement
```

## Checklist de vérification post-patch
- [ ] Aucune entrée utilisateur n'est interpolée directement dans un bloc `<style>` ou un attribut `style`.
- [ ] Les valeurs de personnalisation (couleur, taille) sont validées par une regex stricte avant usage.
- [ ] `innerHTML` n'est plus utilisé pour injecter des balises `<style>` dynamiques ; `style.setProperty`/classes CSS préférés.
- [ ] Une CSP avec `style-src` restrictif est en place en complément.
- [ ] Un test confirme qu'une valeur non conforme (ex: `red; } body { display:none` ou contenant `url(...)`) est rejetée ou neutralisée.
