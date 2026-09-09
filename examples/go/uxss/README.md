# uxss (CWE-79)

La version vulnerable sert une page chargeant un script tiers sans verification d'integrite (Subresource Integrity) et embarquant une iframe tierce sans attribut `sandbox` : si le composant tiers est compromis, le code injecte s'execute avec les pleins privileges de l'origine de l'application. La version corrigee ajoute une Content Security Policy stricte, l'attribut `integrity`/`crossorigin` sur le script tiers et un `sandbox` minimal sur l'iframe, en defense en profondeur.
