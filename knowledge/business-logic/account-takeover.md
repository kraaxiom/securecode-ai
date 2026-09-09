---
id: account-takeover
category: business-logic
cwe: CWE-640
owasp: A07:2021-Identification-and-Authentication-Failures
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Account Takeover (prise de contrôle de compte)

## Description
La prise de contrôle de compte regroupe les failles de logique métier permettant à un attaquant de s'approprier le compte d'un autre utilisateur sans connaître son mot de passe, en exploitant des faiblesses dans les flux de réinitialisation de mot de passe, de changement d'e-mail, de récupération de compte ou de fusion de comptes sociaux. Contrairement au brute force classique, ces attaques exploitent des défauts de logique (jeton prévisible, absence de vérification de propriété, confusion entre identifiants).

## Où ça apparaît typiquement
- Flux de réinitialisation de mot de passe avec jeton prévisible, non expirant, ou non lié à une seule utilisation.
- Changement d'adresse e-mail ou de numéro de téléphone sans revérification du nouveau contact avant qu'il ne devienne actif.
- Connexion via fournisseur tiers (OAuth) fusionnant automatiquement un compte existant sur simple correspondance d'e-mail non vérifié.
- Questions de sécurité ou processus de récupération de compte reposant sur des informations facilement devinables ou publiques.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Jeton de réinitialisation de mot de passe généré de façon prévisible (compteur, timestamp) plutôt que via un générateur cryptographiquement sûr.
- Absence d'expiration ou de limitation d'usage unique sur les jetons de réinitialisation/vérification.
- Flux de changement d'e-mail activant immédiatement la nouvelle adresse sans étape de confirmation sur l'ancienne et la nouvelle adresse.
- Fusion automatique de comptes OAuth sur correspondance d'e-mail sans vérifier que l'e-mail est confirmé par le fournisseur tiers.
- Absence de notification à l'utilisateur lors d'un changement de compte sensible (mot de passe, e-mail, méthode de connexion).

## Remédiation
- Générer les jetons de réinitialisation avec un générateur cryptographiquement sûr, à usage unique et à expiration courte.
- Exiger une confirmation sur l'ancienne ET la nouvelle adresse lors d'un changement d'e-mail ou de méthode de contact.
- Ne fusionner des comptes via OAuth que si l'e-mail est explicitement vérifié par le fournisseur tiers, avec confirmation utilisateur.
- Notifier systématiquement l'utilisateur (par un canal indépendant) lors de tout changement sensible sur son compte.
- Voir `rules/remediation/account-takeover.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/account-takeover/`.

## Références
- OWASP Top 10: A07:2021 – Identification and Authentication Failures
- CWE-640: Weak Password Recovery Mechanism for Forgotten Password
