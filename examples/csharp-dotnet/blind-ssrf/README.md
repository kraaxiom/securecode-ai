# Blind SSRF (SSRF aveugle) — ASP.NET Core / C#

## CWE

CWE-918: Server-Side Request Forgery (SSRF)
OWASP: A10:2021-Server-Side Request Forgery (SSRF)

## Description de la vulnérabilité

`Vulnerable.cs` illustre un cas de SSRF aveugle : l'endpoint
`POST /api/webhooks/register` planifie une requête HTTP sortante
"fire and forget" vers une URL fournie par le client, dans un job asynchrone
dont le résultat n'est jamais renvoyé à l'appelant. L'absence de retour
visible donne une fausse impression de sécurité, mais la requête sortante
reste tout aussi dangereuse : un attaquant peut l'utiliser pour scanner le
réseau interne ou atteindre des services sensibles, en observant des canaux
indirects (délai de réponse, callback DNS/HTTP hors bande) plutôt que le
contenu de la réponse.

## Explication de la correction

`Fixed.cs` applique strictement les mêmes contrôles qu'une SSRF classique
avant de planifier le job asynchrone : whitelist d'hôtes, validation du
schéma, résolution DNS et rejet des adresses privées/loopback/link-local,
timeout. De plus, chaque tentative (succès ou échec) est journalisée via
`ILogger`, ce qui permet une détection a posteriori d'un éventuel scan de
réseau interne déclenché par des workers ou traitements asynchrones — un
point spécifiquement souligné pour la variante "aveugle" de la SSRF.

## Références

- OWASP Cheat Sheet: Server Side Request Forgery Prevention
- CWE-918: Server-Side Request Forgery (SSRF)
- `rules/remediation/blind-ssrf.md`
- `knowledge/ssrf/blind-ssrf.md`
