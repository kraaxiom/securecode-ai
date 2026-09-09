# RC4 — Chiffrement par flux cassé

## CWE
CWE-327 — Use of a Broken or Risky Cryptographic Algorithm

## Description
RC4 présente des biais statistiques exploitables dans son flux de clés et a
fait l'objet d'attaques pratiques réussies (notamment contre WEP et TLS). Il
ne fournit aucune garantie d'intégrité.

## Pourquoi c'est vulnérable
`vulnerable.go` utilise `crypto/rc4` pour chiffrer des données confidentielles.
Le flux de sortie peut être partiellement distingué d'un flux aléatoire,
permettant à un attaquant de retrouver des fragments de texte clair après
observation d'un volume suffisant de trafic chiffré.

## Correction
`fixed.go` remplace RC4 par **AES-256-GCM**, avec nonce aléatoire unique par
message et clé chargée depuis l'environnement (jamais en dur). La configuration
TLS du serveur doit également exclure explicitement les suites RC4.

## Références
- rules/remediation/rc4.md
- knowledge/crypto/rc4.md
- CWE-327: https://cwe.mitre.org/data/definitions/327.html
