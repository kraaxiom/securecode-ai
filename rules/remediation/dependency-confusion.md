# Remédiation — Dependency Confusion

## Principe
Empêcher un gestionnaire de paquets de résoudre un nom de dépendance interne vers un registre public. Réserver le nom sur le registre public, scoper les paquets internes, et forcer explicitement la priorité/portée du registre privé dans la configuration.

## PHP (composer.json / Packagist)
```json
// Avant — vulnérable : dépendance interne résolue via Packagist par défaut
{
    "require": {
        "acme/internal-billing": "^2.0"
    }
}
```
```json
// Après — sécurisé : registre privé explicite + portée verrouillée
{
    "require": {
        "acme/internal-billing": "^2.0"
    },
    "repositories": [
        {
            "type": "composer",
            "url": "https://packages.acme.internal",
            "canonical": true
        },
        {
            "packagist.org": false
        }
    ]
}
```

## JavaScript (package.json / npm)
```json
// Avant — vulnérable : paquet interne sans scope, résolu sur le registre public
{
    "dependencies": {
        "internal-billing": "^2.0.0"
    }
}
```
```json
// Après — sécurisé : scope réservé + registre privé prioritaire
{
    "dependencies": {
        "@acme/internal-billing": "^2.0.0"
    }
}
```
```ini
# .npmrc — après
@acme:registry=https://npm.acme.internal/
always-auth=true
```

## Checklist de vérification post-patch
- [ ] Tout paquet interne est publié sous un scope/namespace réservé (`@acme/...` en npm, vendor dédié en Composer).
- [ ] Le fichier `.npmrc`/`composer.json` fixe explicitement le registre prioritaire pour le scope interne (pas de fallback silencieux vers le registre public).
- [ ] Le nom du paquet interne est également réservé (placeholder) sur le registre public correspondant pour empêcher un typosquat de confusion.
- [ ] Le lockfile (`package-lock.json`, `composer.lock`) est présent et vérifié en CI (intégrité/hash).
