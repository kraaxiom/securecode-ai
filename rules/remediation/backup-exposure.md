# Remédiation — Backup Exposure (exposition de sauvegardes)

## Principe
Stocker les sauvegardes en dehors du webroot, appliquer une politique d'accès restrictive explicite sur tout stockage cloud, chiffrer au repos et en transit, et restreindre les permissions du système de fichiers aux comptes techniques nécessaires.

## PHP
```php
// Avant — vulnérable : dump accessible depuis le webroot
// /var/www/html/backups/dump_2026-08-26.sql  <-- accessible via https://site/backups/dump_...sql

// Après — sécurisé : hors webroot + chiffrement + upload contrôlé
$backupPath = '/var/backups/app/'; // en dehors de /var/www/html
$dumpFile = $backupPath . 'dump_' . date('Y-m-d') . '.sql.enc';
exec(sprintf(
    'mysqldump --defaults-extra-file=%s app_db | openssl enc -aes-256-gcm -pbkdf2 -pass file:%s -out %s',
    escapeshellarg('/etc/app/db.cnf'),
    escapeshellarg('/etc/app/backup.key'),
    escapeshellarg($dumpFile)
));
chmod($dumpFile, 0600); // lisible uniquement par le compte de sauvegarde
```

## JS / Node
```js
// Avant — vulnérable : bucket S3 public par erreur
// aws s3api put-bucket-acl --bucket app-backups --acl public-read

// Après — sécurisé : accès restreint via IAM + chiffrement côté serveur
const s3 = new S3Client({ region: 'eu-west-1' });
await s3.send(new PutObjectCommand({
  Bucket: 'app-backups-private',
  Key: `dumps/${new Date().toISOString()}.sql.gz`,
  Body: dumpStream,
  ServerSideEncryption: 'aws:kms',
  SSEKMSKeyId: process.env.BACKUP_KMS_KEY_ID,
}));
// Politique de bucket : accès refusé par défaut, liste blanche des rôles IAM de sauvegarde uniquement.
```

## Python
```python
# Avant — vulnérable : backup non chiffré dans un répertoire servi par le serveur web
subprocess.run(["pg_dump", "app_db", "-f", "/var/www/html/static/backup.sql"])

# Après — sécurisé
import subprocess, os

backup_dir = "/var/backups/app"  # hors du webroot
os.makedirs(backup_dir, mode=0o700, exist_ok=True)
dump_path = f"{backup_dir}/backup_{datetime.utcnow():%Y%m%d}.sql.gpg"
dump = subprocess.run(["pg_dump", "app_db"], capture_output=True, check=True).stdout
subprocess.run(
    ["gpg", "--symmetric", "--cipher-algo", "AES256", "--output", dump_path],
    input=dump, check=True,
)
os.chmod(dump_path, 0o600)
```

## Checklist de vérification post-patch
- [ ] Aucun fichier de sauvegarde (`.sql`, `.dump`, `.bak`) n'est accessible sous le webroot ou via une route HTTP non authentifiée.
- [ ] Le bucket/stockage cloud de sauvegarde applique une politique d'accès restrictive explicite (deny par défaut).
- [ ] Les sauvegardes sont chiffrées au repos et en transit, avec des clés gérées séparément des données.
- [ ] Les permissions fichiers des sauvegardes sont restreintes (`0600`) au seul compte technique de sauvegarde/restauration.
