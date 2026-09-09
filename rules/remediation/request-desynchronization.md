# Remédiation — Request Smuggling / Desynchronization HTTP

## Principe
Configurer chaque composant de la chaîne HTTP pour rejeter strictement les requêtes ambiguës (présence simultanée de `Content-Length` et `Transfer-Encoding`), normaliser les requêtes au niveau du frontal, et privilégier HTTP/2 de bout en bout entre composants internes lorsque possible.

## Nginx (reverse proxy frontal)
```nginx
# Avant — vulnérable : aucune vérification, transmission telle quelle au backend
location / {
    proxy_pass http://backend;
}

# Après — sécurisé : rejet strict des requêtes ambiguës
map $http_transfer_encoding $bad_te {
    default 0;
    "~*chunked" 1;
}
server {
    location / {
        if ($bad_te = 1) {
            return 400;
        }
        proxy_http_version 1.1;
        proxy_set_header Connection "";
        proxy_pass http://backend;
    }
}
```

## Apache (httpd.conf)
```apache
# Avant — vulnérable
ProxyPass / http://backend/
ProxyPassReverse / http://backend/

# Après — sécurisé : rejet des requêtes avec Content-Length ET Transfer-Encoding
<If "%{HTTP:Transfer-Encoding} != '' && %{HTTP:Content-Length} != ''">
    RewriteEngine On
    RewriteRule .* - [F,L]
</If>
ProxyPass / http://backend/ disablereuse=On
ProxyPassReverse / http://backend/
```

## Node.js (serveur HTTP en frontal ou middleware)
```js
// Avant — vulnérable : aucune validation des en-têtes ambigus
const server = http.createServer((req, res) => proxyRequest(req, res));

// Après — sécurisé : rejet explicite avant transmission au backend
const server = http.createServer((req, res) => {
  const hasTE = 'transfer-encoding' in req.headers;
  const hasCL = 'content-length' in req.headers;
  if (hasTE && hasCL) {
    res.writeHead(400);
    return res.end('Ambiguous request: Content-Length and Transfer-Encoding both present');
  }
  proxyRequest(req, res);
});
```

## Checklist de vérification post-patch
- [ ] Chaque composant de la chaîne (frontal + backend) rejette explicitement les requêtes portant simultanément `Content-Length` et `Transfer-Encoding`.
- [ ] Le frontal normalise/réécrit les requêtes avant transmission plutôt que de les relayer telles quelles.
- [ ] Les composants internes utilisent, quand possible, HTTP/2 de bout en bout ou une implémentation HTTP unique et à jour.
- [ ] Un test de non-régression confirme qu'une requête légitime chunked ou avec Content-Length simple fonctionne toujours normalement.
