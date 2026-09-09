# UXSS (Universal Cross-Site Scripting)

Le handler `dashboard` intégrait une iframe tierce sans attribut `sandbox` ni en-têtes de sécurité (CSP, Permissions-Policy), élargissant la surface d'exposition en cas de faille dans le composant tiers (CWE-79). La version corrigée ajoute un `sandbox` restrictif à l'iframe ainsi qu'une Content-Security-Policy stricte et une Permissions-Policy, réduisant les capacités disponibles pour un composant tiers compromis même si la cause première reste hors du code applicatif.
