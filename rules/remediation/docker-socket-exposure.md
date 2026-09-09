# Remédiation — Montage du socket Docker (`/var/run/docker.sock`) dans un conteneur

## Principe
Éviter de monter le socket Docker de l'hôte dans un conteneur applicatif. Si l'accès au daemon est indispensable (ex: outil CI buildant des images), passer par un proxy d'API filtrant les endpoints autorisés, en lecture seule, et jamais combiné avec du code non fiable.

## docker-compose.yml
```yaml
# Avant — vulnérable : accès total au daemon hôte, équivalent root
services:
  ci-agent:
    image: my-ci-agent:latest
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock

# Après — sécurisé : passage par un proxy filtrant, accès restreint et en lecture seule
services:
  docker-socket-proxy:
    image: tecnativa/docker-socket-proxy:latest
    environment:
      CONTAINERS: 1
      IMAGES: 1
      POST: 0        # bloque les opérations de création/écriture
      BUILD: 0
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
    networks:
      - docker-proxy-net

  ci-agent:
    image: my-ci-agent:latest
    environment:
      DOCKER_HOST: tcp://docker-socket-proxy:2375
    networks:
      - docker-proxy-net
    # plus aucun accès direct au socket de l'hôte

networks:
  docker-proxy-net:
    internal: true
```

## Commande `docker run`
```
# Avant — vulnérable
docker run -v /var/run/docker.sock:/var/run/docker.sock my-tool

# Après — sécurisé : socket monté en lecture seule via le proxy dédié uniquement,
# jamais directement dans le conteneur applicatif
docker run --network docker-proxy-net -e DOCKER_HOST=tcp://docker-socket-proxy:2375 my-tool
```

## Checklist de vérification post-patch
- [ ] Aucun conteneur applicatif ne monte directement `/var/run/docker.sock`.
- [ ] Si un accès au daemon reste nécessaire, il passe par un proxy d'API restreignant explicitement les endpoints autorisés.
- [ ] Le proxy d'API n'autorise pas les opérations d'écriture (`POST`, `BUILD`, `EXEC`) sauf besoin strictement justifié.
- [ ] Le conteneur ayant accès au proxy n'exécute pas de code non fiable (dépendances tierces non auditées, entrées utilisateur).
