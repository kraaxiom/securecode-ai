# XML Injection (CWE-91)

La version vulnérable construit le document `<user><name>{name}</name><role>{role}</role></user>` par interpolation de chaîne, permettant à un attaquant d'injecter des caractères `<`, `>`, `&` pour ajouter ou modifier des nœuds (ex: s'attribuer le rôle `admin`). La correction utilise l'API DOM `System.Xml.XmlDocument` : chaque valeur est affectée via `InnerText`, qui échappe automatiquement les caractères spéciaux XML, garantissant que le contenu utilisateur reste toujours dans les limites du nœud texte prévu.
