# JWT alg:none — CWE-347

La version vulnérable implémente un décodage JWT maison qui, lorsque le header déclare `"alg":"none"`, court-circuite toute vérification de signature et accepte le payload tel quel. Un attaquant peut alors forger un token entièrement dépourvu de signature avec des claims arbitraires (par exemple `role: admin`), sans jamais posséder de secret ni de clé privée, et se faire authentifier comme n'importe quel utilisateur.

La version corrigée abandonne le décodage manuel au profit d'une bibliothèque JWT éprouvée (jjwt), configurée avec une clé de signature obligatoire : `parseClaimsJws` exige une signature valide et rejette automatiquement tout token non signé (`alg: none`) via une exception, avant tout accès au contenu du payload.

**Référence** : CWE-347 (Improper Verification of Cryptographic Signature).
