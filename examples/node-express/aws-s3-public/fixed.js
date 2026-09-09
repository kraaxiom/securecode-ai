// Correction : Bucket S3 public (CWE-284)
// L'objet est uploadé sans ACL publique (le bucket a "Block Public Access"
// activé côté infrastructure) et l'accès en lecture se fait uniquement via
// une URL pré-signée à durée de vie courte, générée à la demande.
const express = require('express');
const { S3Client, PutObjectCommand, GetObjectCommand } = require('@aws-sdk/client-s3');
const { getSignedUrl } = require('@aws-sdk/s3-request-presigner');
const router = express.Router();

const s3 = new S3Client({ region: 'eu-west-1' });
const BUCKET = 'my-app-bucket';

router.post('/uploads', async (req, res) => {
  const key = `uploads/${req.body.filename}`;

  // Pas d'ACL publique : le bucket reste privé par défaut.
  await s3.send(new PutObjectCommand({
    Bucket: BUCKET,
    Key: key,
    Body: req.body.content,
  }));

  // Accès temporaire (15 minutes) au lieu d'une URL publique permanente.
  const signedUrl = await getSignedUrl(
    s3,
    new GetObjectCommand({ Bucket: BUCKET, Key: key }),
    { expiresIn: 900 }
  );
  res.json({ url: signedUrl });
});

module.exports = router;
