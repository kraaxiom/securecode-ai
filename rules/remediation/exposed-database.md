# Remédiation — Exposed Database (base de données exposée)

## Principe
Restreindre l'accès réseau à la base de données aux seules adresses/sous-réseaux applicatifs nécessaires, placer la base dans un sous-réseau privé, activer systématiquement l'authentification forte et changer tout identifiant par défaut.

## Infrastructure (exemple Terraform / groupe de sécurité cloud)
```hcl
# Avant — vulnérable : port ouvert à Internet
resource "aws_security_group_rule" "db_ingress" {
  type              = "ingress"
  from_port         = 5432
  to_port           = 5432
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]   # exposé à Internet
  security_group_id = aws_security_group.db.id
}

# Après — sécurisé : accès restreint au sous-réseau applicatif privé uniquement
resource "aws_security_group_rule" "db_ingress" {
  type                     = "ingress"
  from_port                = 5432
  to_port                  = 5432
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.app.id  # liste blanche stricte
  security_group_id        = aws_security_group.db.id
}
```

## Docker / Compose
```yaml
# Avant — vulnérable : port publié sur toutes les interfaces de l'hôte
services:
  db:
    image: postgres:16
    ports:
      - "5432:5432"   # 0.0.0.0:5432 accessible depuis l'extérieur

# Après — sécurisé : pas de publication externe, accès via réseau interne uniquement
services:
  db:
    image: postgres:16
    environment:
      POSTGRES_PASSWORD_FILE: /run/secrets/db_password
    expose:
      - "5432"        # accessible uniquement aux autres conteneurs du réseau Docker
    networks:
      - internal
networks:
  internal:
    internal: true
```

## Application (connexion applicative)
```php
// Avant — vulnérable : connexion sans mot de passe / identifiant par défaut
$pdo = new PDO('mysql:host=db.example.com;dbname=app', 'root', '');

// Après — sécurisé : identifiants forts, hôte non exposé publiquement, TLS
$pdo = new PDO(
    'mysql:host=db.internal;dbname=app;charset=utf8mb4',
    getenv('DB_USER'),
    getenv('DB_PASSWORD'),
    [PDO::MYSQL_ATTR_SSL_CA => '/etc/ssl/certs/rds-ca.pem']
);
```

## Checklist de vérification post-patch
- [ ] Le port de la base de données n'est accessible depuis aucune adresse IP publique (`0.0.0.0/0` retiré des règles de pare-feu).
- [ ] La base est placée dans un sous-réseau privé, accessible uniquement depuis le réseau applicatif interne.
- [ ] L'authentification est activée et aucun identifiant par défaut n'est utilisé.
- [ ] Un scan externe automatisé confirme l'absence d'exposition du port de base de données sur Internet.
