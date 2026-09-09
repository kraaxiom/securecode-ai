// Correction : Mauvaise configuration IAM (CWE-269)
// L'application assume un rôle IAM temporaire (STS) scoped aux seules
// actions S3 nécessaires sur le bucket applicatif, au lieu de clés
// statiques à privilèges larges.
const express = require('express');
const { S3Client, ListObjectsV2Command } = require('@aws-sdk/client-s3');
const { STSClient, AssumeRoleCommand } = require('@aws-sdk/client-sts');
const router = express.Router();

const sts = new STSClient({ region: 'eu-west-1' });
const APP_BUCKET = 'my-app-bucket';

async function getScopedS3Client() {
  // Rôle applicatif à moindre privilège, scoped à my-app-bucket uniquement.
  const { Credentials } = await sts.send(new AssumeRoleCommand({
    RoleArn: 'arn:aws:iam::123456789012:role/app-read-role',
    RoleSessionName: 'app-session',
    DurationSeconds: 900,
  }));

  return new S3Client({
    region: 'eu-west-1',
    credentials: {
      accessKeyId: Credentials.AccessKeyId,
      secretAccessKey: Credentials.SecretAccessKey,
      sessionToken: Credentials.SessionToken,
    },
  });
}

router.get('/admin/buckets', async (req, res) => {
  const s3 = await getScopedS3Client();

  // Action limitée au bucket applicatif, pas d'accès à l'ensemble du compte.
  const result = await s3.send(new ListObjectsV2Command({ Bucket: APP_BUCKET }));
  res.json(result.Contents ?? []);
});

module.exports = router;
