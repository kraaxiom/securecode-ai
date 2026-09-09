---
id: typosquatting
category: supply-chain
cwe: CWE-1357
owasp: A08:2021-Software-and-Data-Integrity-Failures
severity_default: high
languages: [js, python, php, java, csharp, go, rust]
---

# Typosquatting de paquets

## Description
Le typosquatting consiste, pour un attaquant, à publier un paquet dont le nom ressemble fortement à celui d'un paquet légitime et populaire (faute de frappe proche, séparateur différent, ordre de mots inversé). Un développeur qui commet une erreur de saisie lors de l'installation, ou qui copie une commande erronée trouvée en ligne, installe alors le paquet malveillant à la place du paquet attendu. C'est une technique d'ingénierie sociale ciblant directement la chaîne d'approvisionnement logicielle.

## Où ça apparaît typiquement
- Commandes d'installation manuelles (`npm install`, `pip install`, `composer require`) copiées depuis une source non vérifiée (forum, chat, documentation non officielle).
- Fichiers de dépendances contenant un nom de paquet légèrement différent d'un paquet populaire connu.
- Scripts d'onboarding ou de documentation interne reproduisant une faute de frappe non détectée.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Nom de paquet dans les fichiers de dépendances très proche (distance d'édition faible) d'un paquet populaire du même écosystème sans en être l'orthographe exacte.
- Paquet peu téléchargé/récent portant un nom quasi identique à un paquet à forte popularité.
- Absence de vérification automatisée de la liste des dépendances ajoutées lors des revues de code.

## Remédiation
- Vérifier systématiquement le nom exact et la page officielle du paquet avant ajout aux dépendances.
- Utiliser une allowlist de paquets approuvés au niveau du registre privé/proxy d'entreprise.
- Intégrer un contrôle automatisé de similarité de noms de paquets dans la revue de code ou la CI/CD.
- Former les équipes à copier les commandes d'installation uniquement depuis la documentation officielle.
- Voir `rules/remediation/typosquatting.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/<lang>/typosquatting/`.

## Références
- OWASP Top 10: A08:2021 – Software and Data Integrity Failures
- CWE-1357: Reliance on Insufficiently Trustworthy Component
