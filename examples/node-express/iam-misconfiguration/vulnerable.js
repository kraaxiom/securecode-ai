// Faille : Mauvaise configuration IAM (CWE-269)
// L'application utilise des identifiants AWS statiques dont la policy IAM
// associée accorde "Action: *" / "Resource: *", au lieu d'un rôle
// applicatif restreint. Toute fuite de ces clés compromet tout le compte.
const express = require('express');
const { S3Client, ListBucketsCommand } = require('@aws-sdk/client-s3');
const router = express.Router();

// Identifiants statiques à privilèges larges, codés en dur dans la config.
const s3 = new S3Client({
  region: 'eu-west-1',
  credentials: {
    accessKeyId: process.env.AWS_ACCESS_KEY_ID,   // clé rattachée à une policy "*"/"*"
    secretAccessKey: process.env.AWS_SECRET_ACCESS_KEY,
  },
});

router.get('/admin/buckets', async (req, res) => {
  // Cette route n'a besoin que de lister UN bucket applicatif, mais la
  // policy IAM sous-jacente permet d'agir sur toutes les ressources AWS.
  const result = await s3.send(new ListBucketsCommand({}));
  res.json(result.Buckets);
});

module.exports = router;
