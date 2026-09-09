# Remédiation — Code Injection

## Principe
Éliminer toute évaluation ou exécution dynamique de code construite à partir d'une entrée utilisateur (`eval`, `Function`, `create_function`, désérialisation non sûre). Remplacer par des structures de données déclaratives ou des listes blanches de fonctions autorisées.

## PHP
```php
// Avant — vulnérable
$expr = $_GET['formula'];
$result = eval("return $expr;");

// Après — sécurisé
$allowedOperations = ['add', 'sub', 'mul', 'div'];
$op = $_GET['op'];
if (!in_array($op, $allowedOperations, true)) {
    throw new InvalidArgumentException('Opération non autorisée');
}
$result = match ($op) {
    'add' => $a + $b,
    'sub' => $a - $b,
    'mul' => $a * $b,
    'div' => $a / $b,
};
```

## JavaScript (Node.js)
```js
// Avant — vulnérable
const expr = req.body.formula;
const result = eval(expr);

// Après — sécurisé
const allowedOperations = { add: (a, b) => a + b, sub: (a, b) => a - b };
const op = req.body.op;
if (!Object.prototype.hasOwnProperty.call(allowedOperations, op)) {
  throw new Error('Opération non autorisée');
}
const result = allowedOperations[op](req.body.a, req.body.b);
```

## Python
```python
# Avant — vulnérable
expr = request.form["formula"]
result = eval(expr)

# Après — sécurisé
ALLOWED_OPERATIONS = {"add": lambda a, b: a + b, "sub": lambda a, b: a - b}
op = request.form["op"]
if op not in ALLOWED_OPERATIONS:
    raise ValueError("Opération non autorisée")
result = ALLOWED_OPERATIONS[op](a, b)
```

## Checklist de vérification post-patch
- [ ] Aucun appel `eval`/`exec`/`Function`/`create_function` restant avec une entrée non constante dans le fichier corrigé.
- [ ] Toute désérialisation de données non fiables passe par une liste blanche de types/classes autorisés.
- [ ] Un test de non-régression confirme que les opérations métier légitimes fonctionnent toujours via le mécanisme déclaratif.
- [ ] Un test confirme qu'une valeur d'opération non prévue est rejetée explicitement (exception, code 400).
- [ ] Aucune trace de code dynamique reste dans les chemins accessibles par des entrées externes (webhooks, uploads, paramètres d'URL).
