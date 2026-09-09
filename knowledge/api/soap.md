---
id: soap
category: api
cwe: CWE-611
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: [php, java, csharp, python]
---

# Sécurité des API SOAP

## Description
Les API SOAP reposent sur du XML, ce qui les expose classiquement aux injections d'entités externes XML (XXE) lorsque le parseur XML sous-jacent résout les entités externes déclarées dans une DTD sans restriction. Un parseur mal configuré permet à un attaquant de faire lire des fichiers locaux au serveur ou d'initier des requêtes vers des systèmes internes via l'entité externe. À cela s'ajoute souvent une implémentation faible ou absente de WS-Security (signature, chiffrement des messages), laissant transiter des données sensibles en clair ou sans garantie d'intégrité.

## Où ça apparaît typiquement
- Endpoints SOAP utilisant un parseur XML par défaut sans désactivation explicite du traitement des entités externes et de la résolution DTD.
- Bibliothèques clientes/serveurs SOAP anciennes (versions historiques de frameworks Java, .NET, PHP) dont la configuration sécurisée n'est pas appliquée par défaut.
- Échanges SOAP inter-organisations sans WS-Security ni signature de message, reposant uniquement sur le transport TLS.
- Traitement de fichiers WSDL ou de messages SOAP provenant de sources externes non fiables.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Initialisation d'un parseur XML (`DocumentBuilderFactory`, `XmlDocument`, `libxml`) sans désactivation explicite des entités externes (`setFeature` DOCTYPE/external-general-entities, `XmlResolver = null`, `LIBXML_NOENT` absent des protections).
- Absence de configuration WS-Security (signature, chiffrement des éléments sensibles) dans les échanges contenant des données confidentielles.
- Traitement de documents SOAP/WSDL entrants sans validation de schéma stricte au préalable.
- Version de bibliothèque XML/SOAP connue pour un comportement XXE par défaut, non mise à jour.

## Remédiation
- Désactiver systématiquement la résolution des entités externes et le traitement de DTD dans tout parseur XML utilisé pour traiter des messages SOAP entrants.
- Mettre à jour les bibliothèques XML/SOAP vers des versions dont la configuration sécurisée par défaut désactive XXE.
- Implémenter WS-Security (signature et/ou chiffrement des éléments sensibles) pour les échanges contenant des données confidentielles, en complément de TLS au niveau transport.
- Valider strictement les messages SOAP entrants contre leur schéma XSD avant traitement métier.
- Voir `rules/remediation/soap.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/php/soap/` (et les répertoires équivalents pour java, csharp, python).

## Références
- OWASP Top 10 2021: A05-Security Misconfiguration
- OWASP Cheat Sheet: XML External Entity (XXE) Prevention
- CWE-611: Improper Restriction of XML External Entity Reference
