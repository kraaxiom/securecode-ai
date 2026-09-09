# Weak Password Policy — CWE-521

La version vulnérable n'impose aucune contrainte à l'inscription : un mot de passe d'un seul caractère est accepté, sans vérification contre les listes de mots de passe les plus communs et sans hachage avant stockage. Les comptes ainsi créés sont triviaux à compromettre par attaque en dictionnaire ou brute force, et une fuite de la base expose directement les mots de passe en clair.

La version corrigée applique les recommandations NIST SP 800-63B : une longueur minimale substantielle (12 caractères, jugée plus efficace que des règles de complexité artificielles type "1 majuscule + 1 chiffre + 1 symbole"), un rejet des mots de passe figurant dans une liste de fuites connues, et un hachage systématique avec BCrypt avant tout stockage.

**Référence** : CWE-521 (Weak Password Requirements).
