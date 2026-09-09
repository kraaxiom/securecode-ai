# Mauvaise configuration IAM (CWE-269)

Le code vulnérable crée un rôle IAM avec une trust policy ouverte à n'importe quel compte AWS (`Principal.AWS: "*"`) et une policy de permissions wildcard (`Action: "*"`, `Resource: "*"`) pour un rôle CI qui ne devrait que pousser des images. La correction restreint la trust policy à un provider OIDC précis avec condition `sub` et `ExternalId`, et limite la policy de permissions aux trois actions ECR strictement nécessaires. Élimine la classe de vulnérabilité CWE-269 (Improper Privilege Management).
