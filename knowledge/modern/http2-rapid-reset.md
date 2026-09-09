---
id: http2-rapid-reset
category: modern
cwe: CWE-400
owasp: A04:2021-Insecure Design
severity_default: high
languages: [js, python, java, csharp, go, rust, php]
---

# HTTP/2 Rapid Reset

## Description
Le HTTP/2 Rapid Reset exploite le mécanisme de multiplexage de flux du protocole HTTP/2 : un client ouvre un grand nombre de flux puis les annule immédiatement (frame `RST_STREAM`) avant que le serveur n'ait fini de les traiter, en boucle continue. Le serveur consomme des ressources (CPU, allocation de contexte de requête) pour chaque flux initié, tandis que l'annulation immédiate empêche les limites classiques de connexions concurrentes de s'appliquer, permettant un déni de service avec un nombre restreint de connexions TCP.

## Où ça apparaît typiquement
- Serveurs/reverse proxies HTTP/2 sans limite sur le taux de création/annulation de flux par connexion.
- Infrastructures exposant directement un serveur HTTP/2 sans reverse proxy durci en frontal.
- Composants réseau ou frameworks HTTP/2 utilisant des versions antérieures aux correctifs publiés en 2023 (CVE-2023-44487).

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Absence de limite configurée sur le nombre de flux ouverts/réinitialisés par seconde et par connexion.
- Version de serveur HTTP/2 (Nginx, Envoy, serveurs applicatifs) antérieure aux correctifs de 2023 pour CVE-2023-44487.
- Absence de monitoring sur le ratio flux ouverts/flux annulés par connexion.

## Remédiation
- Mettre à jour les serveurs/proxies HTTP/2 vers des versions corrigeant CVE-2023-44487.
- Configurer une limite stricte sur le nombre de flux créés et annulés par connexion et par intervalle de temps.
- Déployer un reverse proxy/WAF en frontal capable de détecter et limiter ce pattern de trafic.
- Voir `rules/remediation/http2-rapid-reset.md` pour les diffs de configuration.

## Exemple avant/après
Voir `examples/config/http2-rapid-reset/` (configuration serveur avant/après limitation du taux de reset).

## Références
- CISA/OWASP Advisory: HTTP/2 Rapid Reset (CVE-2023-44487)
- CWE-400: Uncontrolled Resource Consumption
