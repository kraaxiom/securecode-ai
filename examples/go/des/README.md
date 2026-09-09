# DES / 3DES — Algorithme de chiffrement cassé

## CWE
CWE-327 — Use of a Broken or Risky Cryptographic Algorithm

## Description
DES utilise une clé de 56 bits effectifs, cassable par force brute en quelques
heures avec du matériel moderne. 3DES, bien que plus robuste, reste vulnérable
à l'attaque Sweet32 (collisions de bloc 64 bits) lorsqu'il chiffre de gros
volumes de données en mode CBC.

## Pourquoi c'est vulnérable
Le fichier `vulnerable.go` utilise `crypto/des` avec une clé codée en dur de 8
octets et un IV nul en mode CBC. Cela expose les données à :
- une attaque par force brute exhaustive sur l'espace de clé (2^56) ;
- une réutilisation d'IV qui annule les propriétés de sécurité du mode CBC ;
- des attaques de collision de bloc sur de longs flux chiffrés.

## Correction
Le fichier `fixed.go` remplace DES par **AES-256 en mode GCM** :
- clé de 256 bits chargée depuis une variable d'environnement (jamais en dur) ;
- chiffrement authentifié (GCM) qui garantit confidentialité et intégrité ;
- nonce aléatoire de 12 octets généré via `crypto/rand` à chaque opération, jamais réutilisé.

## Références
- rules/remediation/des.md
- knowledge/crypto/des.md
- OWASP Cryptographic Storage Cheat Sheet
- CWE-327: https://cwe.mitre.org/data/definitions/327.html
