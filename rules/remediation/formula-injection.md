# Remédiation — Formula Injection (CSV/XLSX)

## Principe
Neutraliser toute valeur de cellule commençant par un caractère déclencheur de formule (`=`, `+`, `-`, `@`, tabulation, retour chariot) avant écriture dans un fichier CSV/XLSX/ODS exporté, en préfixant d'une apostrophe ou en supprimant le caractère déclencheur.

## PHP
```php
// Avant — vulnérable
fputcsv($handle, [$user['nom'], $user['commentaire']]);

// Après — sécurisé
function neutraliserFormule(string $valeur): string {
    if (preg_match('/^[=+\-@\t\r]/', $valeur)) {
        return "'" . $valeur;
    }
    return $valeur;
}
fputcsv($handle, [
    neutraliserFormule($user['nom']),
    neutraliserFormule($user['commentaire']),
]);
```

## JavaScript / Node.js
```js
// Avant — vulnérable
const csvLine = [user.nom, user.commentaire].join(',');
fs.appendFileSync('export.csv', csvLine + '\n');

// Après — sécurisé
function neutraliserFormule(valeur) {
  if (/^[=+\-@\t\r]/.test(valeur)) {
    return `'${valeur}`;
  }
  return valeur;
}
const csvLine = [neutraliserFormule(user.nom), neutraliserFormule(user.commentaire)].join(',');
fs.appendFileSync('export.csv', csvLine + '\n');
```

## Python
```python
# Avant — vulnérable
writer = csv.writer(fichier)
writer.writerow([utilisateur["nom"], utilisateur["commentaire"]])

# Après — sécurisé
def neutraliser_formule(valeur: str) -> str:
    if valeur and valeur[0] in ("=", "+", "-", "@", "\t", "\r"):
        return "'" + valeur
    return valeur

writer = csv.writer(fichier)
writer.writerow([
    neutraliser_formule(utilisateur["nom"]),
    neutraliser_formule(utilisateur["commentaire"]),
])
```

## Checklist de vérification post-patch
- [ ] Toute valeur exportée vers CSV/XLSX/ODS passe par une fonction de neutralisation des caractères déclencheurs (`=`, `+`, `-`, `@`, tabulation, retour chariot).
- [ ] Un test de non-régression confirme que les exports légitimes (sans caractère déclencheur) restent inchangés.
- [ ] Un test confirme qu'une valeur commençant par `=` est correctement préfixée avant écriture dans le fichier.
- [ ] La bibliothèque d'export utilisée est vérifiée pour ne pas ré-interpréter l'apostrophe de neutralisation à l'ouverture.
- [ ] La documentation d'export mentionne ce contrôle pour éviter une régression lors d'un changement de bibliothèque.
