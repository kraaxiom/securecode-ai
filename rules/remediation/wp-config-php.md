# Remédiation — Exposition wp-config.php en webroot public

## Principe
Ce fichier ne doit jamais être atteignable par une requête HTTP directe. La correction se fait au niveau du déploiement/de la configuration serveur, pas dans le code applicatif : retirer le fichier du répertoire servi publiquement, et ajouter une règle de blocage explicite en défense en profondeur.

## Nginx
```nginx
location ~* wp-config\.php\.(bak|old|orig|save|swp|~)$ {
    deny all;
    return 404;
}
location = /wp-config.php {
    include fastcgi_params;
    fastcgi_pass unix:/run/php/php-fpm.sock;
}
```

## Apache (.htaccess ou bloc VirtualHost)
```apache
<FilesMatch "wp-config\.php\.(bak|old|orig|save|swp|~)$">
    Require all denied
</FilesMatch>
<Files wp-config.php>
    Require all granted
</Files>
```

## .gitignore / .dockerignore
Empêcher que le fichier soit committé ou copié dans l'image/artefact de déploiement :
```gitignore
wp-config.php.bak
wp-config.php.old
wp-config.php~
*.php.swp
```

## Vérification CI/CD (avant déploiement)
Ajouter une étape de pipeline qui échoue le build si le motif suivant est détecté dans l'artefact de déploiement : scanner le webroot déployé à la recherche de wp-config.php.bak/.old/.orig/~ et bloquer le déploiement si trouvé ; vérifier que wp-config.php lui-même est bien interprété par PHP et jamais servi en texte brut.

Exemple générique (à adapter au CI utilisé) :
```bash
if find ./dist ./public ./build -type f \( -name "wp-config.php.bak" -o -name "wp-config.php.old" -o -name "wp-config.php~" -o -name "wp-config.php.swp" \) 2>/dev/null | grep -q .; then
  echo "Fichier sensible détecté dans l'artefact de déploiement — build bloqué." >&2
  exit 1
fi
```

## Checklist de vérification post-patch
- [ ] Le fichier (wp-config.php.bak, wp-config.php.old, wp-config.php~, wp-config.php.swp) n'est plus atteignable par une requête HTTP directe (test : `curl -I https://exemple.tld/<fichier>` doit renvoyer 403/404).
- [ ] La règle de blocage serveur (Nginx/Apache) est présente et testée sur l'environnement de préproduction.
- [ ] Le fichier est listé dans `.gitignore`/`.dockerignore` si applicable, pour éviter une réintroduction future.
- [ ] Une étape de vérification CI/CD bloque désormais tout déploiement contenant ce fichier dans le webroot.
- [ ] Si le fichier exposé contenait des secrets, ceux-ci ont été révoqués/régénérés (mots de passe, clés API, clés privées) — l'exposition seule suffit à les considérer compromis.
