# Password Spray — CWE-307

La version vulnérable ne détecte les échecs de connexion que par compte individuel, sans aucune corrélation entre tentatives. Le password spraying exploite précisément cette faiblesse : il teste un seul mot de passe très répandu (ex: `Ete2026!`) contre de nombreux comptes différents, restant ainsi sous n'importe quel seuil de verrouillage "par compte" tout en maximisant les chances de succès statistique.

La version corrigée agrège l'activité par **source** (adresse IP) plutôt que par compte : si une même source teste un nombre anormal de comptes distincts sur une fenêtre de temps donnée, elle est temporairement bloquée. En complément, un contrôle est ajouté au moment du changement de mot de passe pour rejeter les valeurs figurant dans une liste de mots de passe les plus répandus (conforme à l'esprit NIST SP 800-63B), réduisant l'efficacité de l'attaque même si elle contourne la détection de vélocité.

**Référence** : CWE-307 (Improper Restriction of Excessive Authentication Attempts).
