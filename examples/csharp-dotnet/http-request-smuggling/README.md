# HTTP Request Smuggling (CWE-444)

Le middleware vulnérable transmet toute requête au pipeline suivant sans vérifier la cohérence des en-têtes `Content-Length` et `Transfer-Encoding`, ce qui peut être exploité si un proxy en amont interprète différemment la limite entre deux requêtes. La correction ajoute une vérification explicite qui rejette avec un code 400 toute requête portant les deux en-têtes simultanément, en défense en profondeur complémentaire à la configuration du proxy.
