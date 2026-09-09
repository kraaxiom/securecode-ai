# forced-browsing (CWE-425)

La version vulnerable expose un repertoire de rapports internes via un serveur de fichiers statique sans aucune authentification ni autorisation, en s'appuyant uniquement sur le fait que l'URL n'est pas publiee (securite par l'obscurite). La version corrigee ajoute un middleware qui verifie explicitement l'authentification et le role requis avant de servir tout fichier du repertoire sensible. C'est la remediation standard recommandee contre CWE-425.
