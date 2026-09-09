## HTTP Request Smuggling (CWE-444)

Le code vulnérable traite chaque requête reçue sans vérifier la cohérence entre les en-têtes `Content-Length` et `Transfer-Encoding`, ce qui expose l'application à une divergence d'interprétation avec un proxy en amont et permet la contrebande de requêtes. La correction rejette explicitement (code 400) toute requête présentant les deux en-têtes simultanément, en défense en profondeur complémentaire à la configuration du proxy.
