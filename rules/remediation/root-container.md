# Remédiation — Conteneur exécuté en tant que root

## Principe
Créer un utilisateur non privilégié dans l'image et l'activer via l'instruction `USER`, afin que le processus applicatif ne s'exécute jamais avec l'UID 0.

## Dockerfile
```dockerfile
# Avant — vulnérable : le conteneur démarre en root par défaut
FROM node:20-slim
WORKDIR /app
COPY . .
RUN npm ci --production
CMD ["node", "server.js"]

# Après — sécurisé : utilisateur dédié non privilégié
FROM node:20-slim
WORKDIR /app
COPY . .
RUN npm ci --production \
    && groupadd --gid 1001 appgroup \
    && useradd --uid 1001 --gid appgroup --shell /usr/sbin/nologin --no-create-home appuser \
    && chown -R appuser:appgroup /app
USER appuser
CMD ["node", "server.js"]
```

## docker-compose.yml
```yaml
# Avant — vulnérable : aucun utilisateur explicite, hérite du root de l'image
services:
  app:
    image: my-app:latest

# Après — sécurisé : utilisateur non root forcé même si l'image ne le définit pas
services:
  app:
    image: my-app:latest
    user: "1001:1001"
```

## Checklist de vérification post-patch
- [ ] L'instruction `USER <non-root>` est présente en fin de `Dockerfile`, après toutes les opérations nécessitant des droits root (installation de paquets, `chown`).
- [ ] Aucune instruction `USER root` ne réapparaît après le `USER <non-root>`.
- [ ] Les fichiers et répertoires accédés en écriture par l'application (logs, cache, uploads) appartiennent à l'utilisateur non root ou sont montés avec les bonnes permissions.
- [ ] `docker exec <container> whoami` retourne l'utilisateur non privilégié, pas `root`.
- [ ] Le conteneur démarre et fonctionne correctement sans erreur de permission après le changement d'utilisateur.
