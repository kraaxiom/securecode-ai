# Remédiation — HTTP Parameter Pollution (HPP)

## Principe
Définir explicitement le comportement attendu quand un paramètre censé être unique est reçu plusieurs fois : rejeter la requête, ou choisir de façon déterministe et documentée la valeur retenue. Harmoniser ce comportement entre toutes les couches (proxy, WAF, backend).

## PHP
```php
// Avant — vulnérable
// $_GET['role'] ne renvoie qu'une seule valeur silencieusement, sans
// signaler qu'un paramètre dupliqué a été reçu (ex: ?role=user&role=admin
// peut être interprété différemment par un WAF en amont).
$role = $_GET['role'];
assignerRole($role);

// Après — sécurisé
if (isset($_SERVER['QUERY_STRING'])) {
    parse_str($_SERVER['QUERY_STRING'], $rawParams);
}
// Rejeter explicitement toute requête contenant une pollution de paramètre.
$occurrences = substr_count($_SERVER['QUERY_STRING'] ?? '', 'role=');
if ($occurrences > 1) {
    http_response_code(400);
    exit('Paramètre dupliqué non autorisé');
}
$role = filter_input(INPUT_GET, 'role', FILTER_SANITIZE_STRING);
assignerRole($role);
```

## JavaScript / Node.js (Express)
```js
// Avant — vulnérable
// Express, selon le query parser configuré, peut transformer un paramètre
// dupliqué en tableau — utiliser req.query.role tel quel sans vérification
// peut faire passer un tableau là où une chaîne est attendue.
const role = req.query.role;
assignRole(role);

// Après — sécurisé
const roleRaw = req.query.role;
if (Array.isArray(roleRaw)) {
  return res.status(400).send('Paramètre dupliqué non autorisé');
}
const role = String(roleRaw);
assignRole(role);
```

## Python (Flask)
```python
# Avant — vulnérable
# request.args.get() ne retourne que la première valeur silencieusement,
# masquant une éventuelle tentative de pollution de paramètre.
role = request.args.get("role")
assign_role(role)

# Après — sécurisé
values = request.args.getlist("role")
if len(values) > 1:
    abort(400, "Paramètre dupliqué non autorisé")
role = values[0] if values else None
assign_role(role)
```

## Checklist de vérification post-patch
- [ ] Chaque paramètre censé être unique est explicitement vérifié pour rejeter (ou traiter de façon déterministe) les occurrences multiples.
- [ ] Le comportement de parsing des paramètres est identique entre le proxy/WAF en amont et le backend applicatif.
- [ ] Un test confirme qu'une requête avec un paramètre dupliqué critique (ex: `role`, `amount`) est rejetée ou traitée selon la règle documentée.
- [ ] Un test de non-régression confirme que les requêtes légitimes à paramètre unique fonctionnent toujours.
- [ ] La logique métier sensible (autorisation, montant, identifiant) n'utilise jamais un paramètre HTTP sans validation de cardinalité.
