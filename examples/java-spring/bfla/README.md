# BFLA — Broken Function Level Authorization — CWE-862

La version vulnérable protège l'endpoint de suppression d'utilisateur par un simple filtre d'authentification (`requireAuth`), sans jamais vérifier que l'appelant possède le rôle administrateur attendu pour cette fonction. N'importe quel utilisateur authentifié, même sans privilège particulier, peut donc appeler cette fonction sensible en connaissant simplement son URL.

La version corrigée applique un contrôle de rôle explicite et déclaratif via `@PreAuthorize("hasRole('ADMIN')")` (Spring Security), directement sur la définition de l'endpoint. Cette approche centralisée réduit le risque d'oubli comparé à une vérification manuelle dispersée dans la logique métier, et échoue de façon sûre (403) par défaut si le rôle n'est pas vérifié.

**Référence** : CWE-862 (Missing Authorization), référencé comme API5:2023 Broken Function Level Authorization dans l'OWASP API Security Top 10.
