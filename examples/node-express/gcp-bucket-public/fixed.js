// Correction : Bucket Google Cloud Storage public (CWE-284)
// Le fichier reste privé (aucun binding IAM "allUsers"), et l'accès en
// lecture se fait via une URL signée à durée limitée, générée à la demande
// pour l'utilisateur autorisé.
const express = require('express');
const { Storage } = require('@google-cloud/storage');
const router = express.Router();

const storage = new Storage();
const BUCKET = 'my-app-bucket';

router.post('/files', async (req, res) => {
  const bucket = storage.bucket(BUCKET);
  const file = bucket.file(req.body.filename);

  await file.save(req.body.content);

  // Pas d'appel à makePublic() : le bucket reste protégé (Public Access Prevention activé).
  const [signedUrl] = await file.getSignedUrl({
    action: 'read',
    expires: Date.now() + 30 * 60 * 1000, // 30 minutes
  });

  res.json({ url: signedUrl });
});

module.exports = router;
