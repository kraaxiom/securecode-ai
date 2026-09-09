# Remédiation — Weak Privileges (privilèges excessifs en base de données)

## Principe
Créer un rôle applicatif dédié par service, limité strictement aux opérations CRUD nécessaires, séparer le compte de migration (droits DDL, usage ponctuel) du compte d'exécution applicatif, et réaliser une revue périodique des privilèges.

## SQL (création de rôle applicatif)
```sql
-- Avant — vulnérable : l'application se connecte avec le rôle superutilisateur
-- DSN applicatif : postgres://postgres:postgres@db/app

-- Après — sécurisé : rôle applicatif dédié, droits minimaux
CREATE ROLE app_readwrite LOGIN PASSWORD 'REDACTED';
GRANT CONNECT ON DATABASE app TO app_readwrite;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO app_readwrite;
REVOKE CREATE, DROP, ALTER ON SCHEMA public FROM app_readwrite; -- pas de droits DDL

-- Compte de migration séparé, utilisé uniquement en CI/CD, jamais par l'application en runtime
CREATE ROLE app_migrator LOGIN PASSWORD 'REDACTED';
GRANT ALL PRIVILEGES ON SCHEMA public TO app_migrator;
```

## PHP (connexion applicative)
```php
// Avant — vulnérable
$pdo = new PDO('pgsql:host=db;dbname=app', 'postgres', 'postgres');

// Après — sécurisé : compte applicatif dédié, droits CRUD uniquement
$pdo = new PDO('pgsql:host=db;dbname=app', getenv('APP_DB_USER'), getenv('APP_DB_PASSWORD'));
// Les migrations utilisent un compte distinct (APP_MIGRATOR_*), jamais celui de runtime.
```

## Node.js / Python (config de connexion)
```js
// Avant — vulnérable
const pool = new Pool({ user: 'root', password: 'root', database: 'app' });

// Après — sécurisé
const pool = new Pool({
  user: process.env.APP_DB_USER,       // rôle app_readwrite, pas de droits DDL
  password: process.env.APP_DB_PASSWORD,
  database: 'app',
});
```
```python
# Avant — vulnérable
conn = psycopg2.connect(dbname="app", user="postgres", password="postgres")

# Après — sécurisé
conn = psycopg2.connect(
    dbname="app",
    user=os.environ["APP_DB_USER"],      # rôle applicatif dédié, CRUD uniquement
    password=os.environ["APP_DB_PASSWORD"],
)
```

## Checklist de vérification post-patch
- [ ] Le compte de connexion applicatif utilise un rôle dédié limité au CRUD nécessaire, pas le rôle administrateur/superutilisateur.
- [ ] Le compte de migration (droits DDL) est distinct du compte d'exécution applicatif quotidien et n'est utilisé qu'en CI/CD.
- [ ] Chaque service dispose de son propre rôle de base de données, sans partage entre microservices ayant des besoins différents.
- [ ] Une revue périodique documentée des privilèges est planifiée, avec révocation des droits non utilisés.
