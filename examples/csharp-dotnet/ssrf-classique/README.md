# SSRF classique — ASP.NET Core / C#

## CWE

CWE-918: Server-Side Request Forgery (SSRF)
OWASP: A10:2021-Server-Side Request Forgery (SSRF)

## Description de la vulnérabilité

`Vulnerable.cs` expose un endpoint `GET /api/preview?url=...` qui transmet
directement l'URL fournie par le client à `HttpClient.GetAsync`, sans aucune
validation de schéma, d'hôte ou d'adresse IP. Le serveur agit alors comme un
proxy universel : un attaquant peut forcer le back-end à émettre des requêtes
vers des ressources internes normalement inaccessibles depuis Internet
(services d'administration, bases de données, API de métadonnées cloud,
ports internes comme Redis). De plus, `HttpClient` suit les redirections HTTP
par défaut, ce qui permet de contourner un éventuel filtrage effectué
uniquement sur l'URL initiale.

## Explication de la correction

`Fixed.cs` applique une défense en profondeur :

1. **Whitelist stricte d'hôtes** (`AllowedHosts`) — seules les destinations
   métier explicitement autorisées peuvent être contactées.
2. **Validation du schéma** — seul `https` est accepté.
3. **Résolution DNS explicite puis validation de l'IP résultante** — rejette
   les plages privées, loopback, link-local (dont `169.254.169.254`) et
   réservées, même pour un hôte whitelisté (protection contre une
   reconfiguration DNS ultérieure).
4. **Désactivation des redirections automatiques** — chaque redirection
   potentielle est détectée et rejetée plutôt que d'être suivie
   silencieusement, ce qui empêche un contournement de la whitelist via une
   réponse 3xx d'un hôte légitime pointant vers une ressource interne.
5. **Timeout explicite** sur la requête sortante.

## Références

- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
- `rules/remediation/ssrf-classique.md`
- `knowledge/ssrf/ssrf-classique.md`
