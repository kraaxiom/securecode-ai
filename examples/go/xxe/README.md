# xxe (CWE-611)

La version vulnerable normalise le XML utilisateur via l'outil externe `xmllint --noent`, qui resout les entites externes declarees dans une DTD (`<!ENTITY xxe SYSTEM "file:///etc/passwd">`), permettant la lecture de fichiers arbitraires ou du SSRF. La version corrigee utilise `encoding/xml` de la bibliotheque standard Go, qui ne resout jamais les entites externes ni les DTD par defaut, avec une carte d'entites vide en defense en profondeur explicite.
