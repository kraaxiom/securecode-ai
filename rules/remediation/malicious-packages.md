# Remédiation — Paquets malveillants

## Principe
Retirer immédiatement tout paquet identifié comme malveillant (ou dont la version installée correspond à une release compromise), auditer les scripts de cycle de vie exécutés à l'installation, et verrouiller les installations futures par hash/signature.

## PHP (composer.json)
```json
// Avant — vulnérable : paquet compromis avec script post-install non audité
{
    "require": {
        "some-vendor/utils": "^3.1"
    },
    "scripts": {
        "post-install-cmd": "some-vendor/utils/setup.php"
    }
}
```
```json
// Après — sécurisé : paquet retiré/remplacé, intégrité vérifiée, scripts désactivés par défaut
{
    "require": {
        "acme/utils-safe": "^1.0"
    },
    "config": {
        "allow-plugins": false
    }
}
```
```bash
composer remove some-vendor/utils
composer require acme/utils-safe --prefer-dist
```

## JavaScript (package.json)
```json
// Avant — vulnérable : paquet compromis exécutant un postinstall obscurci
{
    "dependencies": {
        "left-pad-utils": "^1.3.7"
    }
}
```
```json
// Après — sécurisé : paquet retiré, scripts d'installation désactivés globalement
{
    "dependencies": {
        "left-pad-utils": "REMOVED — voir alternative auditée"
    }
}
```
```ini
# .npmrc — après
ignore-scripts=true
```

## Checklist de vérification post-patch
- [ ] Le paquet malveillant identifié est retiré du fichier de dépendances et du lockfile (aucune référence transitive résiduelle).
- [ ] Les scripts `postinstall`/`preinstall` du projet et des dépendances directes ont été audités manuellement.
- [ ] Toute clé/secret potentiellement exposée pendant la période de compromission a été révoquée et régénérée.
- [ ] Un scanner de paquets malveillants (en complément du SCA classique) est intégré à la CI/CD.
