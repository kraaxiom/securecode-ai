# Remédiation — SSRF vers le service de métadonnées Azure (IMDS)

## Principe
Ne jamais laisser une entrée utilisateur déterminer, directement ou indirectement, l'hôte/IP/port d'une requête réseau sortante émise par le serveur. Appliquer une whitelist stricte de destinations autorisées, résoudre le DNS puis valider l'adresse IP résultante (plages privées/loopback/link-local exclues), et désactiver les redirections automatiques non revalidées.

## PHP
```php
// Avant — vulnérable
$url = $_POST['callback_url'];
$response = file_get_contents($url);

// Après — sécurisé
$allowedHosts = ['api.partenaire.example.com'];
$parts = parse_url($_POST['callback_url']);
if (!$parts || !in_array($parts['host'] ?? '', $allowedHosts, true) || ($parts['scheme'] ?? '') !== 'https') {
    http_response_code(400);
    exit('URL non autorisée');
}
$ip = gethostbyname($parts['host']);
if (filter_var($ip, FILTER_VALIDATE_IP, FILTER_FLAG_NO_PRIV_RANGE | FILTER_FLAG_NO_RES_RANGE) === false) {
    http_response_code(400);
    exit('Destination interdite');
}
$ch = curl_init($_POST['callback_url']);
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, false);
curl_setopt($ch, CURLOPT_TIMEOUT, 5);
$response = curl_exec($ch);
```

## Node.js
```js
// Avant — vulnérable
const response = await fetch(req.body.url);

// Après — sécurisé
const { URL } = require('url');
const dns = require('dns').promises;

const ALLOWED_HOSTS = ['api.partenaire.example.com'];

async function safeFetch(rawUrl) {
  const parsed = new URL(rawUrl);
  if (parsed.protocol !== 'https:' || !ALLOWED_HOSTS.includes(parsed.hostname)) {
    throw new Error('URL non autorisée');
  }
  const { address } = await dns.lookup(parsed.hostname);
  if (isPrivateOrReserved(address)) {
    throw new Error('Destination interdite');
  }
  return fetch(parsed.toString(), { redirect: 'manual', signal: AbortSignal.timeout(5000) });
}
```

## Python (requests)
```python
# Avant — vulnérable
resp = requests.get(request.form['url'])

# Après — sécurisé
import ipaddress
import socket
from urllib.parse import urlparse

ALLOWED_HOSTS = {"api.partenaire.example.com"}

def safe_get(raw_url):
    parsed = urlparse(raw_url)
    if parsed.scheme != "https" or parsed.hostname not in ALLOWED_HOSTS:
        raise ValueError("URL non autorisée")
    ip = socket.gethostbyname(parsed.hostname)
    addr = ipaddress.ip_address(ip)
    if addr.is_private or addr.is_loopback or addr.is_link_local or addr.is_reserved:
        raise ValueError("Destination interdite")
    return requests.get(raw_url, allow_redirects=False, timeout=5)
```

## Note spécifique — Métadonnées Azure (IMDS)
Filtrer explicitement l'adresse `169.254.169.254` (et le nom d'hôte de métadonnées le cas échéant) dans toute logique de validation des destinations de requêtes sortantes, en plus de la whitelist générale. Endpoint sensible : `169.254.169.254 (en-tête Metadata: true requis)`. Côté infrastructure, appliquer le principe du moindre privilège sur le rôle/l'identité/le compte de service attaché à la ressource, et privilégier la variante la plus sécurisée du service de métadonnées quand elle existe (ex: IMDSv2 pour AWS).

## Checklist de vérification post-patch
- [ ] La destination de la requête sortante est validée contre une whitelist explicite de domaines/hôtes métier.
- [ ] Le DNS est résolu et l'adresse IP obtenue est vérifiée comme non privée/non loopback/non link-local avant l'émission de la requête.
- [ ] Les redirections HTTP automatiques sont désactivées ou chaque redirection est revalidée avec les mêmes contrôles.
- [ ] Un timeout raisonnable est appliqué à toute requête sortante déclenchée par une entrée utilisateur.
