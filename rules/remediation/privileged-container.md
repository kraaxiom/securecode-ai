# Remédiation — Conteneur en mode privilégié

## Principe
Supprimer le mode `--privileged`/`privileged: true` et n'accorder que les capabilities Linux réellement nécessaires, en partant d'une base sans aucun privilège (`cap_drop: ALL`).

## docker-compose.yml
```yaml
# Avant — vulnérable : toutes les capabilities accordées, isolation quasi désactivée
services:
  app:
    image: my-app:latest
    privileged: true

# Après — sécurisé : capabilities minimales explicites
services:
  app:
    image: my-app:latest
    cap_drop:
      - ALL
    cap_add:
      - NET_BIND_SERVICE   # uniquement si le processus doit écouter un port < 1024
    security_opt:
      - no-new-privileges:true
    read_only: true
```

## Commande `docker run`
```
# Avant — vulnérable
docker run --privileged my-app

# Après — sécurisé
docker run --cap-drop=ALL --cap-add=NET_BIND_SERVICE --security-opt=no-new-privileges:true my-app
```

## Cas nécessitant un accès matériel réel (à isoler et documenter)
```yaml
# Si un accès à un périphérique spécifique est réellement requis, le déclarer
# explicitement plutôt que d'activer le mode privilégié complet
services:
  video-encoder:
    image: video-encoder:latest
    devices:
      - "/dev/dri:/dev/dri"   # accès GPU ciblé uniquement
    cap_drop:
      - ALL
```

## Checklist de vérification post-patch
- [ ] `privileged: true` / `--privileged` n'apparaît plus dans aucune configuration de déploiement.
- [ ] Le conteneur démarre avec `cap_drop: ALL` puis n'ajoute que les capabilities strictement nécessaires via `cap_add`.
- [ ] `no-new-privileges:true` est activé pour empêcher l'élévation de privilèges via des binaires setuid.
- [ ] Les profils seccomp et AppArmor/SELinux par défaut de Docker ne sont pas désactivés (`--security-opt seccomp=unconfined` absent).
- [ ] L'application fonctionne toujours correctement après retrait du mode privilégié (test de non-régression fonctionnel).
