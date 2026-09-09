// Faille : Conteneur Azure Blob Storage public (CWE-284)
// Le conteneur est créé avec un niveau d'accès public "container", ce qui
// autorise la lecture anonyme de tous les blobs qu'il contient. L'URL du
// blob est renvoyée telle quelle, sans jeton d'accès limité dans le temps.
const express = require('express');
const { BlobServiceClient } = require('@azure/storage-blob');
const router = express.Router();

const blobService = BlobServiceClient.fromConnectionString(process.env.AZURE_STORAGE_CONNECTION_STRING);
const CONTAINER = 'myfiles';

router.post('/documents', async (req, res) => {
  const containerClient = blobService.getContainerClient(CONTAINER);

  // Accès public au niveau conteneur : n'importe qui peut lister/lire les blobs.
  await containerClient.createIfNotExists({ access: 'container' });

  const blockBlobClient = containerClient.getBlockBlobClient(req.body.filename);
  await blockBlobClient.upload(req.body.content, req.body.content.length);

  // URL publique permanente, sans SAS ni expiration.
  res.json({ url: blockBlobClient.url });
});

module.exports = router;
