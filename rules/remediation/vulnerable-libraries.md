# Remédiation — Bibliothèques vulnérables ou obsolètes

## Principe
Remplacer ou mettre à niveau toute bibliothèque non maintenue ou structurellement à risque (même sans CVE précise) vers une alternative activement maintenue, et établir un inventaire (SBOM) pour suivre l'état de maintenance dans le temps.

## PHP (composer.json)
```json
// Avant — vulnérable : dépendance abandonnée (dernier commit > 3 ans, dépôt archivé)
{
    "require": {
        "abandoned/old-serializer": "^1.2"
    }
}
```
```json
// Après — sécurisé : remplacement par une alternative maintenue
{
    "require": {
        "symfony/serializer": "^6.4"
    }
}
```

## JavaScript (package.json)
```json
// Avant — vulnérable : dépendance majeure très en retard, dépôt archivé
{
    "dependencies": {
        "request": "^2.88.2"
    }
}
```
```json
// Après — sécurisé : remplacement par une alternative maintenue
{
    "dependencies": {
        "undici": "^6.0.0"
    }
}
```

## Checklist de vérification post-patch
- [ ] La bibliothèque de remplacement est activement maintenue (release récente, dépôt non archivé).
- [ ] Un SBOM (Software Bill of Materials) est généré et inclut cette dépendance avec son statut de maintenance.
- [ ] Les usages de l'ancienne API dans le code applicatif ont été migrés et testés (pas de régression fonctionnelle).
- [ ] Un scanner SCA signale désormais l'état de maintenance des dépendances en CI/CD, pas seulement les CVE connues.
