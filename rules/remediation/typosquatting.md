# Remédiation — Typosquatting de paquets

## Principe
Corriger le nom de tout paquet mal orthographié vers le paquet officiel exact, et mettre en place un contrôle automatisé de similarité de noms pour empêcher toute réintroduction future.

## PHP (composer.json)
```json
// Avant — vulnérable : faute de frappe sur un paquet populaire
{
    "require": {
        "guzzlehtttp/guzzle": "^7.0"
    }
}
```
```json
// Après — sécurisé : nom exact du paquet officiel
{
    "require": {
        "guzzlehttp/guzzle": "^7.0"
    }
}
```

## JavaScript (package.json)
```json
// Avant — vulnérable : faute de frappe sur un paquet populaire
{
    "dependencies": {
        "expresss": "^4.18.0"
    }
}
```
```json
// Après — sécurisé : nom exact du paquet officiel
{
    "dependencies": {
        "express": "^4.18.0"
    }
}
```

## Checklist de vérification post-patch
- [ ] Le nom exact du paquet correspond à la page officielle du registre (npm/Packagist) et au dépôt source attendu.
- [ ] La revue de code inclut désormais un contrôle automatisé de distance d'édition sur tout nouveau nom de dépendance ajouté.
- [ ] Une allowlist de paquets approuvés est en place au niveau du registre privé/proxy d'entreprise.
- [ ] Le lockfile a été régénéré après correction et ne référence plus le paquet squatté.
