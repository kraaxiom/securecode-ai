# LDAP Injection

La version vulnérable construit le filtre LDAP `(uid=...)` par concaténation directe du paramètre `uid`, ce qui correspond à CWE-90 : des caractères spéciaux LDAP (`*`, `(`, `)`, `\`) permettent de modifier la logique du filtre, potentiellement pour contourner une authentification. La version corrigée applique `escape_filter_chars()` du module `python-ldap`, qui échappe correctement ces caractères avant construction du filtre.
