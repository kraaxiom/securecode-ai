# Remédiation — Exposition archive de sauvegarde (.zip/.tar.gz/.bak) en webroot public

## Principe
Ce fichier ne doit jamais être atteignable par une requête HTTP directe. La correction se fait au niveau du déploiement/de la configuration serveur, pas dans le code applicatif : retirer le fichier du répertoire servi publiquement, et ajouter une règle de blocage explicite en défense en profondeur.

## Nginx
```nginx
location ~* \.(zip|tar|tar\.gz|tgz|bak|old|sql\.gz|backup)$ {
    deny all;
    return 404;
}
```

## Apache (.htaccess ou bloc VirtualHost)
```apache
<FilesMatch "\.(zip|tar|tar\.gz|tgz|bak|old|sql\.gz|backup)$">
    Require all denied
</FilesMatch>
```

## .gitignore / .dockerignore
Empêcher que le fichier soit committé ou copié dans l'image/artefact de déploiement :
```gitignore
*.zip
*.tar.gz
*.tgz
*.bak
*.old
*.sql.gz
backup/
backups/
```

## Vérification CI/CD (avant déploiement)
Ajouter une étape de pipeline qui échoue le build si le motif suivant est détecté dans l'artefact de déploiement : aucun fichier d'archive (*.zip, *.tar.gz, *.bak, *.sql.gz) ne doit se trouver sous le répertoire déployé comme webroot.

Exemple générique (à adapter au CI utilisé) :
```bash
if find ./dist ./public ./build -type f \( -name "*.zip" -o -name "*.tar.gz" -o -name "*.tgz" -o -name "*.bak" -o -name "*.old" -o -name "*.sql.gz" \) 2>/dev/null | grep -q .; then
  echo "Fichier sensible détecté dans l'artefact de déploiement — build bloqué." >&2
  exit 1
fi
```

## Checklist de vérification post-patch
- [ ] Le fichier (*.zip, *.tar.gz, *.tgz, *.bak, *.old, *.sql.gz) n'est plus atteignable par une requête HTTP directe (test : `curl -I https://exemple.tld/<fichier>` doit renvoyer 403/404).
- [ ] La règle de blocage serveur (Nginx/Apache) est présente et testée sur l'environnement de préproduction.
- [ ] Le fichier est listé dans `.gitignore`/`.dockerignore` si applicable, pour éviter une réintroduction future.
- [ ] Une étape de vérification CI/CD bloque désormais tout déploiement contenant ce fichier dans le webroot.
- [ ] Si le fichier exposé contenait des secrets, ceux-ci ont été révoqués/régénérés (mots de passe, clés API, clés privées) — l'exposition seule suffit à les considérer compromis.
