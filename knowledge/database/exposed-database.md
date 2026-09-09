---
id: exposed-database
category: database
cwe: CWE-284
owasp: A05:2021-Security-Misconfiguration
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Exposed Database (base de données exposée)

## Description
Ce pattern couvre les cas où une instance de base de données est directement accessible depuis un réseau non fiable (Internet public) sans pare-feu réseau restrictif, souvent combinée à des identifiants par défaut ou faibles. C'est l'une des causes les plus fréquentes de fuites massives de données, l'attaquant n'ayant même pas besoin d'exploiter une faille applicative : un simple scan de ports suffit à découvrir l'instance exposée.

## Où ça apparaît typiquement
- Instances de base de données déployées sur le cloud avec un groupe de sécurité/firewall autorisant `0.0.0.0/0` sur le port de la base.
- Environnements de développement ou de test exposant temporairement une base sur Internet pour "faciliter l'accès", jamais refermés.
- Conteneurs Docker publiant le port de la base de données sur toutes les interfaces réseau de l'hôte par erreur de configuration.
- Bases de données NoSQL configurées sans authentification activée par défaut lors de l'installation.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Règle de pare-feu/groupe de sécurité cloud autorisant l'accès au port de la base de données depuis n'importe quelle adresse IP.
- Fichier de configuration Docker/Compose publiant le port de la base de données sur l'interface réseau publique de l'hôte (`0.0.0.0`).
- Base de données sans authentification activée ou utilisant les identifiants par défaut du produit.
- Absence de segmentation réseau entre la couche base de données et les réseaux publics ou moins fiables.

## Remédiation
- Restreindre l'accès réseau à la base de données aux seules adresses/sous-réseaux applicatifs nécessaires (liste blanche stricte).
- Placer la base de données dans un sous-réseau privé sans exposition directe à Internet, accessible uniquement via le réseau applicatif interne.
- Activer systématiquement l'authentification forte et changer tout identifiant par défaut dès l'installation.
- Auditer régulièrement l'exposition réseau des instances via des scans externes automatisés.
- Voir `rules/remediation/exposed-database.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/exposed-database/`.

## Références
- OWASP Top 10: A05:2021 – Security Misconfiguration
- CWE-284: Improper Access Control
