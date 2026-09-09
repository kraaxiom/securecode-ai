# Remédiation — Exposition credentials.json en webroot public

## Principe
Ce fichier ne doit jamais être atteignable par une requête HTTP directe. La correction se fait au niveau du déploiement/de la configuration serveur, pas dans le code applicatif : retirer le fichier du répertoire servi publiquement, et ajouter une règle de blocage explicite en défense en profondeur.

## Nginx
```nginx
location ~* /credentials\.json$ {
    deny all;
    return 404;
}
location ~* \.(json)$ {
    # Si ce dossier ne sert que des assets publics légitimes, restreindre par whitelist plutôt que par blacklist
    deny all;
    return 404;
}
```

## Apache (.htaccess ou bloc VirtualHost)
```apache
<FilesMatch "^credentials\.json$">
    Require all denied
</FilesMatch>
```

## .gitignore / .dockerignore
Empêcher que le fichier soit committé ou copié dans l'image/artefact de déploiement :
```gitignore
credentials.json
*.credentials.json
service-account*.json
```

## Vérification CI/CD (avant déploiement)
Ajouter une étape de pipeline qui échoue le build si le motif suivant est détecté dans l'artefact de déploiement : bloquer le build/déploiement si un fichier nommé credentials.json (ou *service-account*.json) est présent hors d'un coffre-fort de secrets.

Exemple générique (à adapter au CI utilisé) :
```bash
if find ./dist ./public ./build -type f \( -name "credentials.json" -o -name "*-credentials.json" -o -name "service-account*.json" \) 2>/dev/null | grep -q .; then
  echo "Fichier sensible détecté dans l'artefact de déploiement — build bloqué." >&2
  exit 1
fi
```

## Checklist de vérification post-patch
- [ ] Le fichier (credentials.json, *-credentials.json, service-account*.json) n'est plus atteignable par une requête HTTP directe (test : `curl -I https://exemple.tld/<fichier>` doit renvoyer 403/404).
- [ ] La règle de blocage serveur (Nginx/Apache) est présente et testée sur l'environnement de préproduction.
- [ ] Le fichier est listé dans `.gitignore`/`.dockerignore` si applicable, pour éviter une réintroduction future.
- [ ] Une étape de vérification CI/CD bloque désormais tout déploiement contenant ce fichier dans le webroot.
- [ ] Si le fichier exposé contenait des secrets, ceux-ci ont été révoqués/régénérés (mots de passe, clés API, clés privées) — l'exposition seule suffit à les considérer compromis.
