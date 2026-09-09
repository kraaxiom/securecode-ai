# DNS Rebinding (contournement de filtre SSRF) — ASP.NET Core / C#

## CWE

CWE-918: Server-Side Request Forgery (SSRF)
OWASP: A10:2021-Server-Side Request Forgery (SSRF)

## Description de la vulnérabilité

`Vulnerable.cs` valide correctement en apparence l'hôte fourni par
l'utilisateur : résolution DNS explicite puis rejet des adresses privées.
Le problème est que cette résolution est totalement indépendante de celle
que `HttpClient` effectuera lui-même au moment de la connexion réelle
(`_httpClient.GetAsync`). Un attaquant contrôlant le domaine cible peut
configurer un TTL DNS très court et faire répondre le serveur DNS avec une
adresse IP publique légitime lors de la validation, puis avec une adresse
interne (loopback, RFC1918, `169.254.169.254`) lors de la résolution
effective — exploitant le décalage temporel entre les deux résolutions
(Time-Of-Check / Time-Of-Use).

## Explication de la correction

`Fixed.cs` applique le **DNS pinning** via un `SocketsHttpHandler` doté d'un
`ConnectCallback` personnalisé : la résolution DNS n'est effectuée qu'une
seule fois, l'adresse IP obtenue est validée immédiatement, puis la
connexion TCP est établie directement sur cette adresse — sans laisser le
client HTTP effectuer une seconde résolution DNS qui pourrait renvoyer une
IP différente. Le pinning est appliqué au niveau de la connexion TCP
elle-même, pas seulement au niveau applicatif, ce qui élimine la fenêtre
TOCTOU exploitée par le rebinding.

## Références

- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
- `knowledge/ssrf/dns-rebinding.md`
