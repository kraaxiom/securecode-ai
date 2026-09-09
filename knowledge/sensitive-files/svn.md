---
id: svn
category: sensitive-files
cwe: CWE-538
owasp: A05:2021-Security Misconfiguration
severity_default: high
languages: []
---

# Exposition du répertoire .svn

## Description
Le répertoire `.svn` (Subversion) présent dans chaque dossier d'une working copy contient des métadonnées et parfois des copies en texte brut des fichiers versionnés, y compris des versions antérieures potentiellement sensibles. Comme il est répété dans chaque sous-dossier, sa présence sur un serveur web mal configuré peut exposer des fragments de code source ou de configuration dans plusieurs répertoires distincts, pas seulement à la racine.

## Où ça apparaît typiquement
- Déploiement par checkout SVN direct dans le webroot au lieu d'un export propre.
- Absence de règle serveur bloquant les répertoires `.svn` récursivement.
- Anciens projets migrés depuis SVN où les métadonnées n'ont pas été nettoyées après export.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Présence de `.svn/` dans un ou plusieurs répertoires du webroot déployé.
- Réponse HTTP 200 sur une requête vers `/.svn/entries` ou `/.svn/wc.db`.
- Absence de directive de blocage récursive des dotfiles/dossiers cachés dans la configuration serveur.

## Remédiation
- Utiliser `svn export` (sans métadonnées) plutôt qu'un `checkout` pour générer les artefacts de déploiement.
- Bloquer explicitement l'accès à tout répertoire `.svn` au niveau du serveur web, de façon récursive.
- Intégrer une étape de nettoyage dans le pipeline de déploiement qui vérifie l'absence de dossiers de VCS.
- Voir `rules/remediation/svn.md` pour les diffs par configuration serveur.

## Exemple avant/après
Voir `examples/config/svn/` (configuration serveur avant/après blocage de l'accès).

## Références
- OWASP Top 10: A05:2021-Security Misconfiguration
- CWE-538: Insertion of Sensitive Information into Externally-Accessible File or Directory
