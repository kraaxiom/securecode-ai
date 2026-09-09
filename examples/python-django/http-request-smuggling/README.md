# HTTP Request Smuggling

La version vulnérable traite le corps de la requête sans vérifier la cohérence des en-têtes de délimitation, ce qui correspond à CWE-444 : une divergence d'interprétation entre le proxy en amont et Django sur la présence conjointe de `Content-Length` et `Transfer-Encoding` peut permettre de faire traiter une partie de la requête comme une requête distincte. La version corrigée ajoute un middleware `reject_ambiguous_requests` qui rejette explicitement ce cas ambigu, en défense en profondeur de la configuration proxy.
