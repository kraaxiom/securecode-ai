# Remédiation — Secrets intégrés dans une image Docker

## Principe
Ne jamais copier un secret dans une couche d'image, même temporairement. Utiliser les secrets de build BuildKit (`RUN --mount=type=secret`) qui ne persistent pas dans l'historique des couches, et injecter les secrets d'exécution via des variables d'environnement du conteneur ou un gestionnaire de secrets dédié.

## Dockerfile — ARG/ENV en clair
```dockerfile
# Avant — vulnérable : le secret reste visible via `docker history`
FROM php:8.3-fpm
ARG API_KEY=sk_live_EXAMPLE_NOT_A_REAL_KEY
ENV DATABASE_PASSWORD=SuperSecret123
COPY . /var/www/html

# Après — sécurisé : secret injecté à l'exécution uniquement, jamais dans l'image
FROM php:8.3-fpm
COPY . /var/www/html
# API_KEY et DATABASE_PASSWORD sont fournis via l'environnement du conteneur
# au lancement (docker run -e / secrets Kubernetes / Docker secrets), pas au build
```

## Dockerfile — copie puis suppression d'un fichier de credentials
```dockerfile
# Avant — vulnérable : le fichier reste récupérable dans la couche où il a été copié
FROM composer:2 AS build
COPY composer.json composer.json.lock ./
COPY auth.json /root/.composer/auth.json
RUN composer install --no-dev
RUN rm /root/.composer/auth.json

# Après — sécurisé : secret de build BuildKit, jamais écrit dans une couche
FROM composer:2 AS build
COPY composer.json composer.json.lock ./
RUN --mount=type=secret,id=composer_auth,target=/root/.composer/auth.json \
    composer install --no-dev
```

Build correspondant :
```
DOCKER_BUILDKIT=1 docker build --secret id=composer_auth,src=./auth.json -t my-app .
```

## docker-compose.yml — injection à l'exécution via secrets
```yaml
# Après — sécurisé : secret monté en fichier à l'exécution, absent de l'image
services:
  app:
    image: my-app:latest
    secrets:
      - db_password

secrets:
  db_password:
    file: ./secrets/db_password.txt
```

## Checklist de vérification post-patch
- [ ] Aucune instruction `ARG`/`ENV` du `Dockerfile` ne contient de valeur de secret littérale.
- [ ] `docker history --no-trunc <image>` ne révèle aucun secret dans les couches intermédiaires.
- [ ] Les secrets nécessaires au build utilisent `RUN --mount=type=secret` (BuildKit) et non `COPY`/`ADD`.
- [ ] Les secrets d'exécution sont injectés via variables d'environnement du conteneur, Docker secrets, ou un gestionnaire de secrets externe (Vault, AWS Secrets Manager, etc.), jamais au moment du build.
- [ ] Un build multi-stage garantit que les étapes manipulant des secrets ne sont pas présentes dans l'image finale publiée.
