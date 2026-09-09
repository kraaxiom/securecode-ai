// Faille : Bucket S3 public (CWE-284)
// Le bucket est créé avec une ACL "public-read" et l'application renvoie
// directement l'URL publique de l'objet au client. N'importe qui disposant
// de cette URL (ou la devinant) peut lire les fichiers sans authentification.
const express = require('express');
const { S3Client, PutObjectCommand } = require('@aws-sdk/client-s3');
const router = express.Router();

const s3 = new S3Client({ region: 'eu-west-1' });
const BUCKET = 'my-app-bucket';

router.post('/uploads', async (req, res) => {
  const key = `uploads/${req.body.filename}`;

  // ACL publique : l'objet devient accessible à tout le monde sur Internet.
  await s3.send(new PutObjectCommand({
    Bucket: BUCKET,
    Key: key,
    Body: req.body.content,
    ACL: 'public-read',
  }));

  // URL publique permanente renvoyée telle quelle, sans contrôle d'accès.
  const publicUrl = `https://${BUCKET}.s3.amazonaws.com/${key}`;
  res.json({ url: publicUrl });
});

module.exports = router;
