# Mass Assignment — CWE-915

La version vulnérable désérialise directement le corps de la requête PUT dans l'entité `User` complète, y compris ses champs sensibles `role` et `isAdmin`, avant de la persister telle quelle. Un attaquant peut ainsi ajouter `"role":"ADMIN"` au JSON envoyé — alors que le formulaire prévu ne propose que le nom et l'email — et s'auto-promouvoir administrateur.

La version corrigée introduit un DTO d'entrée dédié (`UpdateUserRequest`) qui n'expose structurellement que les champs réellement modifiables par l'utilisateur (liste blanche explicite) : les champs sensibles n'existent tout simplement pas dans ce DTO et ne peuvent donc jamais être désérialisés depuis la requête. L'affectation vers l'entité persistée se fait ensuite champ par champ, de façon explicite.

**Référence** : CWE-915 (Improperly Controlled Modification of Dynamically-Determined Object Attributes), référencé comme API3:2023 Broken Object Property Level Authorization dans l'OWASP API Security Top 10.
