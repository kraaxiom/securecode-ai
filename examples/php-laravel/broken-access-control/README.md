# Broken Access Control (catégorie générale)

Le code vulnérable ne vérifie que l'authentification avant de permettre le téléchargement d'une facture, en s'appuyant à tort sur le masquage du lien côté interface. La version corrigée introduit une `InvoicePolicy` centralisée appliquant le principe "deny by default" : l'accès n'est accordé que si l'utilisateur est propriétaire de la facture ou administrateur. Cette faille correspond à CWE-284 (Improper Access Control), tel qu'indiqué dans `knowledge/authorization/broken-access-control.md`.
