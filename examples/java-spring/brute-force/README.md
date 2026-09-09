# Brute Force — CWE-307

La version vulnérable ne limite en rien le nombre de tentatives de connexion : aucun compteur d'échecs, aucun verrouillage de compte, aucun délai. Un attaquant peut donc soumettre un très grand nombre de mots de passe contre le même compte jusqu'à trouver le bon, sans être ralenti ni détecté.

La version corrigée introduit un verrouillage temporaire du compte (15 minutes) après 5 échecs consécutifs, indépendant de l'adresse IP source pour résister à la rotation d'IP, avec un message d'erreur identique que le compte soit verrouillé ou que les identifiants soient simplement invalides (pour ne pas révéler l'état du compte). En production, cette protection doit être complétée par un rate-limiter au niveau réseau/passerelle et par un CAPTCHA déclenché après quelques échecs.

**Référence** : CWE-307 (Improper Restriction of Excessive Authentication Attempts).
