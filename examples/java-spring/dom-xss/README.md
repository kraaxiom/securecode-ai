# DOM-based XSS — CWE-79

La version vulnérable sert une page dont le script client lit le paramètre d'URL `q` (source non fiable, contrôlée par l'attaquant) et l'injecte directement via `innerHTML` (sink dangereux) dans le DOM, sans jamais transiter par le serveur ni y être analysé. Une URL forgée du type `?q=<img src=x onerror=alert(1)>` exécute alors du code arbitraire dans le navigateur de la victime.

La version corrigée remplace le sink dangereux par `textContent`, qui traite systématiquement la valeur comme du texte brut et non comme du HTML interprétable — la classe de vulnérabilité est éliminée à la source. En complément, une politique CSP stricte (`script-src 'self'`, `object-src 'none'`) est appliquée côté serveur via Spring Security pour limiter l'impact de tout sink dangereux qui subsisterait ailleurs dans l'application.

**Référence** : CWE-79 (Improper Neutralization of Input During Web Page Generation), variante DOM-based XSS.
