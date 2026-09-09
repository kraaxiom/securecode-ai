# OAuth Misconfiguration

Les handlers `oauth_login`/`oauth_callback` ne généraient ni ne vérifiaient de paramètre `state`, exposant le callback OAuth à une CSRF permettant de forcer la liaison du compte d'un attaquant à la session d'une victime (CWE-287). La version corrigée génère un `state` aléatoire stocké côté serveur (cookie `HttpOnly`), le vérifie strictement au retour, et fixe le `redirect_uri` par correspondance exacte plutôt que par motif large.
