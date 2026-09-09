# Privilege Escalation — CWE-269

La version vulnérable permet à un utilisateur de modifier lui-même son propre champ `role` via un endpoint de profil, sans restreindre les valeurs acceptées ni vérifier son rôle actuel. Un utilisateur standard peut donc s'attribuer le rôle `ADMIN` en soumettant simplement `role=ADMIN`, obtenant une élévation de privilèges verticale complète.

La version corrigée restreint les rôles auto-attribuables à une liste blanche métier explicite qui exclut tout rôle privilégié, et déplace la possibilité de promouvoir un utilisateur vers un rôle sensible dans un endpoint distinct, protégé par `@PreAuthorize("hasRole('ADMIN')")` et donc réservé aux administrateurs existants. Aucun chemin ne permet plus à un utilisateur de s'auto-promouvoir.

**Référence** : CWE-269 (Improper Privilege Management).
