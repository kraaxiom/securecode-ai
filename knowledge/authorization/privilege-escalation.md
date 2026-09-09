---
id: privilege-escalation
category: authorization
cwe: CWE-269
owasp: A01:2021-Broken Access Control
severity_default: critical
languages: [php, js, python, java, csharp, go, rust]
---

# Privilege Escalation

## Description
L'escalade de privilèges désigne la capacité d'un utilisateur à obtenir des droits supérieurs à ceux qui lui ont été légitimement attribués, que ce soit horizontalement (accéder aux données d'un autre utilisateur de même niveau) ou verticalement (obtenir des droits d'administrateur depuis un compte standard). Elle résulte souvent d'une combinaison de défauts : mass assignment sur un champ de rôle, contrôle d'autorisation manquant sur une fonction d'administration, ou logique métier de changement de rôle mal protégée.

## Où ça apparaît typiquement
- Endpoints de mise à jour de profil permettant de modifier, directement ou indirectement, un champ de rôle/permission.
- Fonctions d'invitation ou de gestion d'équipe permettant d'attribuer un rôle sans vérifier que l'appelant a lui-même le droit d'attribuer ce rôle.
- Jetons/sessions dont le niveau de privilège n'est pas revalidé après un changement de rôle côté serveur.

## Indicateurs de détection (niveau pattern, pas d'exploitation)
- Champ de rôle/permission présent dans le même modèle que les champs modifiables par l'utilisateur, sans protection contre le mass assignment (voir `mass-assignment.md`).
- Fonction d'attribution de rôle ne vérifiant pas que le rôle demandé est inférieur ou égal au niveau de privilège de l'appelant.
- Session/token dont les claims de rôle ne sont pas revérifiés après une opération sensible de changement de permissions.

## Remédiation
- Protéger explicitement tout champ de rôle/permission contre la modification directe par l'utilisateur (voir remédiation mass assignment).
- Vérifier, pour toute fonction d'attribution de rôle, que l'appelant est lui-même autorisé à accorder ce niveau de privilège précis.
- Invalider/rafraîchir les sessions ou tokens après tout changement de rôle afin d'éviter la persistance d'anciens privilèges.
- Voir `rules/remediation/privilege-escalation.md` pour les diffs par langage.

## Exemple avant/après
Voir `examples/java-spring/privilege-escalation/` (à créer selon le même schéma pour chaque langage listé ci-dessus).

## Références
- OWASP Cheat Sheet: Authorization
- CWE-269: Improper Privilege Management
