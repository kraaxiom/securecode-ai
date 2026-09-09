# OAuth Misconfiguration — CWE-287

La version vulnérable accepte n'importe quelle `redirect_uri` fournie par le client sans la valider contre une liste blanche, et ne génère ni ne vérifie le paramètre `state` censé protéger le flux OAuth contre le CSRF. Un attaquant peut ainsi rediriger le code d'autorisation vers un domaine qu'il contrôle (vol de code), ou forcer une victime à lier son compte au compte OAuth de l'attaquant (login CSRF), en initiant lui-même le flux et en faisant suivre le lien de callback à la victime.

La version corrigée valide strictement la `redirect_uri` contre une liste blanche d'URIs enregistrées, génère un `state` cryptographiquement aléatoire lié à la session courante à l'étape `/authorize`, et exige sa vérification stricte (correspondance exacte, usage unique) au retour du callback avant tout échange du code contre un jeton d'accès.

**Référence** : CWE-287 (Improper Authentication).
