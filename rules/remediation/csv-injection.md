# Remédiation — CSV Injection (Formula Injection)

## Principe
Préfixer toute cellule dont le premier caractère est `=`, `+`, `-`, `@`, une tabulation ou un retour chariot par un caractère neutralisant (apostrophe) avant écriture dans un fichier CSV, ou utiliser une bibliothèque d'export qui applique cet échappement par défaut.

## PHP
```php
// Avant — vulnérable
$fp = fopen('export.csv', 'w');
foreach ($rows as $row) {
    fputcsv($fp, [$row['name'], $row['comment']]);
}

// Après — sécurisé
function sanitizeCsvCell(string $value): string {
    if (preg_match('/^[=+\-@\t\r]/', $value)) {
        return "'" . $value;
    }
    return $value;
}

$fp = fopen('export.csv', 'w');
foreach ($rows as $row) {
    fputcsv($fp, [sanitizeCsvCell($row['name']), sanitizeCsvCell($row['comment'])]);
}
```

## Node.js
```js
// Avant — vulnérable
const csv = rows.map(r => `${r.name},${r.comment}`).join('\n');
res.set('Content-Type', 'text/csv');
res.send(csv);

// Après — sécurisé
function sanitizeCsvCell(value) {
  return /^[=+\-@\t\r]/.test(value) ? `'${value}` : value;
}
const csv = rows
  .map(r => [sanitizeCsvCell(r.name), sanitizeCsvCell(r.comment)].join(','))
  .join('\n');
res.set('Content-Type', 'text/csv');
res.send(csv);
```

## Python
```python
# Avant — vulnérable
import csv
with open("export.csv", "w", newline="") as f:
    writer = csv.writer(f)
    for row in rows:
        writer.writerow([row["name"], row["comment"]])

# Après — sécurisé
import csv
import re

def sanitize_csv_cell(value: str) -> str:
    if re.match(r"^[=+\-@\t\r]", value):
        return "'" + value
    return value

with open("export.csv", "w", newline="") as f:
    writer = csv.writer(f)
    for row in rows:
        writer.writerow([sanitize_csv_cell(row["name"]), sanitize_csv_cell(row["comment"])])
```

## Checklist de vérification post-patch
- [ ] Toute cellule commençant par `=`, `+`, `-`, `@`, tabulation ou retour chariot est préfixée par un caractère neutralisant avant écriture.
- [ ] Un test confirme qu'une valeur utilisateur commençant par `=` s'ouvre comme texte brut dans un tableur, pas comme formule.
- [ ] L'export utilise une bibliothèque CSV maintenue plutôt qu'une génération manuelle de chaîne.
- [ ] L'en-tête `Content-Disposition` force le téléchargement plutôt que l'ouverture directe non maîtrisée.
