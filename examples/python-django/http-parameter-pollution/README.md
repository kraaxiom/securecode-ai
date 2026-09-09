# HTTP Parameter Pollution

La version vulnérable utilise `request.GET.get("role")`, qui ne retourne silencieusement que la première valeur d'un paramètre dupliqué, ce qui correspond à CWE-235 : une divergence d'interprétation entre un proxy/WAF en amont et Django peut être exploitée pour contourner un contrôle. La version corrigée utilise `request.GET.getlist("role")` et rejette explicitement toute requête contenant plusieurs occurrences du paramètre `role`.
