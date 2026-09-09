# HTTP Request Smuggling (CWE-444)

Le code vulnérable détecte la présence des en-têtes `Content-Length` et `Transfer-Encoding` mais ne réagit pas à leur coexistence, laissant persister l'ambiguïté qui permet la contrebande de requêtes entre proxy et backend. La correction rejette explicitement toute requête présentant les deux en-têtes simultanément avec un code 400, en défense en profondeur en complément de la configuration du proxy, conformément à CWE-444.
