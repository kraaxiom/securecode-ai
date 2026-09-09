# Broken Access Control — CWE-284

La version vulnérable ne protège l'accès au rapport financier que par le masquage du lien correspondant dans l'interface utilisateur (affiché seulement si `isAdmin` côté client). L'API serveur elle-même n'applique aucune vérification : n'importe quel client HTTP capable de deviner ou d'observer l'URL de l'endpoint peut y accéder directement, sans jamais passer par l'interface.

La version corrigée déplace le contrôle d'accès là où il doit se trouver : côté serveur, via `@PreAuthorize("hasRole('FINANCE_ADMIN')")` (Spring Security), qui échoue de façon sûre (HTTP 403) par défaut si le rôle requis n'est pas présent. Le masquage côté client peut rester en place pour l'ergonomie, mais ne constitue jamais, à lui seul, une mesure de sécurité.

**Référence** : CWE-284 (Improper Access Control), catégorie générale couvrant A01:2021 Broken Access Control de l'OWASP Top 10.
