# SVG XSS

Le handler `serve_avatar` servait les fichiers SVG uploadés tels quels, avec `Content-Type: image/svg+xml`, depuis l'origine de l'application, permettant l'exécution de script contenu dans le SVG (CWE-79). La version corrigée sanitise le contenu XML en retirant les balises `script` avant de servir le fichier, et force le téléchargement via `Content-Disposition: attachment` plutôt qu'un affichage inline, limitant l'impact d'un contenu résiduel malveillant.
