# Mauvaise configuration IAM (permissions excessives) (CWE-269)

La version vulnérable crée programmatiquement, via le SDK AWS, une policy IAM attachée à un rôle applicatif avec `Action: "*"` et `Resource: "*"`, accordant un accès total à tous les services AWS au lieu des seules permissions nécessaires. La correction définit une policy scoped listant précisément les actions S3 requises sur le bucket applicatif, applique le principe du moindre privilège, et restreint la trust policy d'assomption de rôle à un principal précis avec condition `ExternalId`.
