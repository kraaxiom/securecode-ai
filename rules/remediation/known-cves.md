# Remédiation — Composants avec CVE connues

## Principe
Mettre à jour toute dépendance figée sur une version affectée par une CVE publique vers une version corrigée, et intégrer un scanner SCA en CI/CD pour empêcher la régression.

## PHP (composer.json)
```json
// Avant — vulnérable : version figée avec CVE publiée (ex: symfony/http-foundation < 5.4.20)
{
    "require": {
        "symfony/http-foundation": "5.4.19"
    }
}
```
```json
// Après — sécurisé : version corrigée avec borne minimale explicite
{
    "require": {
        "symfony/http-foundation": "^5.4.20"
    }
}
```
```bash
composer update symfony/http-foundation
composer audit
```

## JavaScript (package.json)
```json
// Avant — vulnérable : version figée avec CVE publiée
{
    "dependencies": {
        "lodash": "4.17.15"
    }
}
```
```json
// Après — sécurisé
{
    "dependencies": {
        "lodash": "^4.17.21"
    }
}
```
```bash
npm audit fix
npm audit --audit-level=high
```

## Checklist de vérification post-patch
- [ ] La version installée correspond à une release qui corrige la CVE identifiée (vérifiée dans l'avisory NVD/GHSA/OSV).
- [ ] `composer audit` / `npm audit` (ou équivalent SCA) ne remonte plus la CVE après mise à jour.
- [ ] Le lockfile est régénéré et commité avec la nouvelle version.
- [ ] Un scan SCA est intégré à la CI/CD pour bloquer toute réintroduction d'une version vulnérable connue.
