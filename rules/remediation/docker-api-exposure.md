# Remédiation — Exposition de l'API Docker Engine (daemon TCP non authentifié)

## Principe
Ne jamais exposer le daemon Docker sur un socket TCP non authentifié. Utiliser le socket Unix local par défaut, et si un accès distant est réellement requis, l'activer uniquement avec TLS mutuel obligatoire.

## Configuration systemd (`docker.service`)
```
# Avant — vulnérable
ExecStart=/usr/bin/dockerd -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock

# Après — sécurisé
ExecStart=/usr/bin/dockerd -H unix:///var/run/docker.sock
```

## Configuration `daemon.json`
```json
// Avant — vulnérable
{
  "hosts": ["tcp://0.0.0.0:2375", "unix:///var/run/docker.sock"]
}

// Après — sécurisé (socket local uniquement)
{
  "hosts": ["unix:///var/run/docker.sock"]
}
```

## Si un accès distant est indispensable (TLS mutuel obligatoire)
```json
// Après — accès distant sécurisé
{
  "hosts": ["tcp://0.0.0.0:2376"],
  "tls": true,
  "tlsverify": true,
  "tlscacert": "/etc/docker/certs/ca.pem",
  "tlscert": "/etc/docker/certs/server-cert.pem",
  "tlskey": "/etc/docker/certs/server-key.pem"
}
```

## Restriction réseau complémentaire (pare-feu / security group)
```
# Avant — port 2375 ouvert à 0.0.0.0/0
Ingress: tcp/2375 from 0.0.0.0/0

# Après — accès restreint au réseau privé de gestion uniquement
Ingress: tcp/2376 from 10.0.0.0/24 (TLS mutuel obligatoire)
```

## Checklist de vérification post-patch
- [ ] Le daemon Docker n'écoute plus sur un port TCP sans authentification (`docker info` / `ss -tlnp | grep 2375` ne retourne rien).
- [ ] Si un accès distant est conservé, `tlsverify=true` est actif et des certificats clients valides sont exigés.
- [ ] Le port de l'API Docker (2375/2376) n'est pas accessible depuis Internet dans les règles de pare-feu/security group.
- [ ] Un test de connexion anonyme (`docker -H tcp://<host>:2375 ps`) échoue depuis un poste externe non autorisé.
