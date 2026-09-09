# Remédiation — DNS Rebinding (contournement SSRF)

## Principe
Ne jamais valider un hôte via une résolution DNS distincte de celle réellement utilisée par la connexion sortante. Résoudre le DNS une seule fois, valider l'adresse IP obtenue, puis établir la connexion directement sur cette IP validée (DNS pinning) au lieu de laisser le client HTTP re-résoudre le nom séparément.

## PHP (cURL avec résolution épinglée)
```php
// Avant — vulnérable (résolution DNS séparée de la requête réelle)
$host = parse_url($url, PHP_URL_HOST);
$ip = gethostbyname($host);
if (!is_private_ip($ip)) {
    $response = file_get_contents($url); // cURL/DNS résout à nouveau ici -> rebinding possible
}

// Après — sécurisé (IP épinglée via CURLOPT_RESOLVE)
$host = parse_url($url, PHP_URL_HOST);
$ip = gethostbyname($host);
if (is_private_ip($ip)) {
    throw new RuntimeException('Destination interdite');
}
$ch = curl_init($url);
curl_setopt($ch, CURLOPT_RESOLVE, ["$host:443:$ip"]); // force l'IP validée
curl_setopt($ch, CURLOPT_FOLLOWLOCATION, false);
$response = curl_exec($ch);
```

## Node.js (résolution épinglée avec agent HTTP personnalisé)
```js
// Avant — vulnérable
const response = await fetch(url); // Node re-résout le DNS à la connexion

// Après — sécurisé
const dns = require('dns').promises;
const https = require('https');

const { address } = await dns.lookup(parsedUrl.hostname);
if (isPrivateOrReserved(address)) throw new Error('Destination interdite');

const agent = new https.Agent({
  lookup: (hostname, options, callback) => callback(null, address, 4),
});
const response = await fetch(url, { agent, redirect: 'manual' });
```

## Python (résolution épinglée)
```python
# Avant — vulnérable
ip = socket.gethostbyname(hostname)
if not is_private(ip):
    resp = requests.get(url)  # requests/urllib3 re-résout le DNS séparément

# Après — sécurisé (épinglage via HTTPAdapter personnalisé)
import requests
from requests.adapters import HTTPAdapter

class PinnedDNSAdapter(HTTPAdapter):
    def __init__(self, pinned_ip, *args, **kwargs):
        self.pinned_ip = pinned_ip
        super().__init__(*args, **kwargs)

ip = socket.gethostbyname(hostname)
if is_private(ip):
    raise ValueError("Destination interdite")
session = requests.Session()
session.mount("https://", PinnedDNSAdapter(ip))
resp = session.get(url, allow_redirects=False, timeout=5)
```

## Checklist de vérification post-patch
- [ ] L'adresse IP validée est la même que celle réellement utilisée pour établir la connexion sortante (DNS pinning effectif).
- [ ] Aucune revalidation d'URL n'est différée dans le temps sans re-résolution/re-vérification complète.
- [ ] Le client/proxy sortant applique un filtrage réseau au niveau de la connexion TCP, pas uniquement au niveau applicatif.
- [ ] Le service effectuant les requêtes sortantes est isolé réseau des ressources internes sensibles, en défense en profondeur.
