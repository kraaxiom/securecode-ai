// Faille : Bucket Google Cloud Storage public (CWE-284)
// L'application accorde le rôle "storage.objectViewer" à "allUsers" sur
// chaque fichier uploadé, rendant l'objet lisible par n'importe qui sur
// Internet, puis expose l'URL publique permanente au client.
const express = require('express');
const { Storage } = require('@google-cloud/storage');
const router = express.Router();

const storage = new Storage();
const BUCKET = 'my-app-bucket';

router.post('/files', async (req, res) => {
  const bucket = storage.bucket(BUCKET);
  const file = bucket.file(req.body.filename);

  await file.save(req.body.content);

  // Rend l'objet public à tous ("allUsers") : aucune restriction d'accès.
  await file.makePublic();

  res.json({ url: `https://storage.googleapis.com/${BUCKET}/${req.body.filename}` });
});

module.exports = router;
