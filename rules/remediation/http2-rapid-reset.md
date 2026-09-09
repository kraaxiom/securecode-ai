# Remédiation — HTTP/2 Rapid Reset (CVE-2023-44487)

## Principe
Mettre à jour les serveurs/proxies HTTP/2 vers les versions corrigeant CVE-2023-44487, et configurer une limite stricte sur le nombre de flux créés/annulés par connexion et par intervalle de temps. Compléter par un reverse proxy/WAF frontal capable de détecter ce pattern de trafic.

## Nginx
```nginx
# Avant — vulnérable : aucune limite sur le taux de reset de flux HTTP/2
http {
    server {
        listen 443 ssl http2;
        ...
    }
}

# Après — sécurisé : mise à jour + limitation explicite
# 1. S'assurer que nginx >= 1.25.3 (ou version backportée avec le correctif CVE-2023-44487)
http {
    http2_max_concurrent_streams 100;
    limit_conn_zone $binary_remote_addr zone=conn_limit:10m;
    server {
        listen 443 ssl http2;
        limit_conn conn_limit 20;
        ...
    }
}
```

## Envoy
```yaml
# Avant — vulnérable
http2_protocol_options: {}

# Après — sécurisé : limite explicite de flux concurrents + reset rapide détecté
http2_protocol_options:
  max_concurrent_streams: 100
overload_manager:
  refresh_interval: 0.25s
  resource_monitors:
    - name: "envoy.resource_monitors.global_downstream_max_connections"
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.resource_monitors.downstream_connections.v3.DownstreamConnectionsConfig
        max_active_downstream_connections: 5000
```

## Node.js (serveur HTTP/2 natif)
```js
// Avant — vulnérable : aucune limite sur les flux par session
const server = http2.createSecureServer(options);

// Après — sécurisé : mise à jour Node.js (patchée) + limite explicite
const server = http2.createSecureServer({
  ...options,
  settings: { maxConcurrentStreams: 100 },
});
server.on('session', (session) => {
  let resets = 0;
  session.on('stream', (stream) => {
    stream.on('close', () => {
      if (stream.rstCode !== 0) {
        resets++;
        if (resets > 50) session.destroy();
      }
    });
  });
});
```

## Checklist de vérification post-patch
- [ ] Le serveur/proxy HTTP/2 est mis à jour vers une version corrigeant CVE-2023-44487.
- [ ] Une limite de flux concurrents (`max_concurrent_streams` / équivalent) est configurée par connexion.
- [ ] Un mécanisme détecte et coupe les connexions présentant un ratio anormal d'ouverture/annulation de flux.
- [ ] Un reverse proxy/WAF frontal est en place pour absorber/filtrer ce type de trafic avant le backend applicatif.
