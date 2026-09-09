# MFA Bypass — CWE-287

La version vulnérable marque la session comme pleinement authentifiée dès la vérification du mot de passe, avant même l'étape MFA. Les routes protégées ne contrôlent que ce flag `isAuthenticated`, jamais un flag MFA dédié : un attaquant connaissant le mot de passe peut accéder directement aux ressources sensibles en ignorant simplement l'appel à `/mfa/verify`, rendant le second facteur purement décoratif.

La version corrigée introduit un état intermédiaire explicite ("en attente de MFA") entre la vérification du mot de passe et la connexion complète : la session ne passe à l'état "pleinement authentifiée" qu'après validation du code MFA, et chaque route protégée vérifie ce statut complet plutôt que la seule authentification par mot de passe. Il est recommandé de centraliser ce contrôle dans un filtre Spring Security unique plutôt que de le dupliquer par contrôleur, afin d'éviter tout oubli.

**Référence** : CWE-287 (Improper Authentication).
