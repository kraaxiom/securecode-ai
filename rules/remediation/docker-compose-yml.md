# Remédiation — Exposition docker-compose.yml en webroot public

## Principe
Ce fichier ne doit jamais être atteignable par une requête HTTP directe. La correction se fait au niveau du déploiement/de la configuration serveur, pas dans le code applicatif : retirer le fichier du répertoire servi publiquement, et ajouter une règle de blocage explicite en défense en profondeur.

## Nginx
```nginx
location ~* /docker-compose(\.\w+)?\.ya?ml$ {
    deny all;
    return 404;
}
```

## Apache (.htaccess ou bloc VirtualHost)
```apache
<FilesMatch "^docker-compose(\.\w+)?\.ya?ml$">
    Require all denied
</FilesMatch>
```

## .gitignore / .dockerignore
Empêcher que le fichier soit committé ou copié dans l'image/artefact de déploiement :
```gitignore
docker-compose.override.yml
docker-compose.prod.yml
.env.docker
```

## Vérification CI/CD (avant déploiement)
Ajouter une étape de pipeline qui échoue le build si le motif suivant est détecté dans l'artefact de déploiement : docker-compose*.yml ne doit jamais être copié dans l'image/artefact final servi en production ; utiliser un .dockerignore dédié pour le build.

Exemple générique (à adapter au CI utilisé) :
```bash
if find ./dist ./public ./build -type f \( -name "docker-compose.yml" -o -name "docker-compose.*.yml" -o -name "docker-compose.override.yml" \) 2>/dev/null | grep -q .; then
  echo "Fichier sensible détecté dans l'artefact de déploiement — build bloqué." >&2
  exit 1
fi
```

## Checklist de vérification post-patch
- [ ] Le fichier (docker-compose.yml, docker-compose.*.yml, docker-compose.override.yml) n'est plus atteignable par une requête HTTP directe (test : `curl -I https://exemple.tld/<fichier>` doit renvoyer 403/404).
- [ ] La règle de blocage serveur (Nginx/Apache) est présente et testée sur l'environnement de préproduction.
- [ ] Le fichier est listé dans `.gitignore`/`.dockerignore` si applicable, pour éviter une réintroduction future.
- [ ] Une étape de vérification CI/CD bloque désormais tout déploiement contenant ce fichier dans le webroot.
- [ ] Si le fichier exposé contenait des secrets, ceux-ci ont été révoqués/régénérés (mots de passe, clés API, clés privées) — l'exposition seule suffit à les considérer compromis.
