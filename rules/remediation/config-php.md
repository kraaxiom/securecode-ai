# Remédiation — Exposition config.php en webroot public

## Principe
Ce fichier ne doit jamais être atteignable par une requête HTTP directe. La correction se fait au niveau du déploiement/de la configuration serveur, pas dans le code applicatif : retirer le fichier du répertoire servi publiquement, et ajouter une règle de blocage explicite en défense en profondeur.

## Nginx
```nginx
location ~* /config(\.inc)?\.php\.(bak|old|orig|save|swp)$ {
    deny all;
    return 404;
}
# S'assurer que les .php sont interprétés et jamais servis en texte brut
location ~ \.php$ {
    include fastcgi_params;
    fastcgi_pass unix:/run/php/php-fpm.sock;
}
```

## Apache (.htaccess ou bloc VirtualHost)
```apache
<FilesMatch "config(\.inc)?\.php\.(bak|old|orig|save|swp)$">
    Require all denied
</FilesMatch>
# Empêcher l'exécution/lecture de fichiers de config hors du répertoire applicatif prévu
<Directory "/var/www/html/config">
    Require all denied
</Directory>
```

## .gitignore / .dockerignore
Empêcher que le fichier soit committé ou copié dans l'image/artefact de déploiement :
```gitignore
config.local.php
config.production.php
*.php.bak
*.php.orig
```

## Vérification CI/CD (avant déploiement)
Ajouter une étape de pipeline qui échoue le build si le motif suivant est détecté dans l'artefact de déploiement : aucun fichier config*.php.bak/.old/.orig ne doit être présent dans l'artefact de déploiement ; vérifier que le vhost interprète bien tous les .php (pas de fallback texte brut).

Exemple générique (à adapter au CI utilisé) :
```bash
if find ./dist ./public ./build -type f \( -name "config.php" -o -name "config.inc.php" \) 2>/dev/null | grep -q .; then
  echo "Fichier sensible détecté dans l'artefact de déploiement — build bloqué." >&2
  exit 1
fi
```

## Checklist de vérification post-patch
- [ ] Le fichier (config.php, config.inc.php, configuration.php (et copies .bak/.old/.orig)) n'est plus atteignable par une requête HTTP directe (test : `curl -I https://exemple.tld/<fichier>` doit renvoyer 403/404).
- [ ] La règle de blocage serveur (Nginx/Apache) est présente et testée sur l'environnement de préproduction.
- [ ] Le fichier est listé dans `.gitignore`/`.dockerignore` si applicable, pour éviter une réintroduction future.
- [ ] Une étape de vérification CI/CD bloque désormais tout déploiement contenant ce fichier dans le webroot.
- [ ] Si le fichier exposé contenait des secrets, ceux-ci ont été révoqués/régénérés (mots de passe, clés API, clés privées) — l'exposition seule suffit à les considérer compromis.
