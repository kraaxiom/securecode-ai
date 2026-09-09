## HTTP Parameter Pollution (CWE-235)

Le code vulnérable utilise `req.query.role` sans vérifier si le paramètre a été envoyé plusieurs fois, ce qui produit un tableau côté Express et peut créer une divergence d'interprétation entre les couches (proxy/WAF vs backend) exploitable pour contourner un contrôle métier. La correction détecte explicitement le cas d'un paramètre dupliqué (`Array.isArray`) et rejette la requête avec un code 400 plutôt que de traiter silencieusement une valeur ambiguë.
