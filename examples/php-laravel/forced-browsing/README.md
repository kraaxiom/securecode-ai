# Forced Browsing

Le code vulnérable expose une page de rapports internes sans aucun middleware d'authentification ni d'autorisation, en s'appuyant uniquement sur le fait que la route n'est pas liée dans l'interface (sécurité par obscurité). La version corrigée ajoute un contrôle explicite (`auth` + `can:viewInternalReports`) appliqué à la route elle-même, qu'elle soit ou non découvrable depuis l'UI. Cette faille correspond à CWE-425 (Direct Request / 'Forced Browsing'), tel qu'indiqué dans `knowledge/authorization/forced-browsing.md`.
