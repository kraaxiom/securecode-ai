// Correction : Conteneur Azure Blob Storage public (CWE-284)
// Le conteneur est créé sans accès public ("access" omis = private) et
// l'accès en lecture se fait via une signature SAS en lecture seule,
// avec une expiration courte générée à la demande.
const express = require('express');
const { BlobServiceClient, generateBlobSASQueryParameters, BlobSASPermissions, StorageSharedKeyCredential } = require('@azure/storage-blob');
const router = express.Router();

const blobService = BlobServiceClient.fromConnectionString(process.env.AZURE_STORAGE_CONNECTION_STRING);
const CONTAINER = 'myfiles';

router.post('/documents', async (req, res) => {
  const containerClient = blobService.getContainerClient(CONTAINER);

  // Conteneur privé par défaut : aucun accès anonyme.
  await containerClient.createIfNotExists();

  const blockBlobClient = containerClient.getBlockBlobClient(req.body.filename);
  await blockBlobClient.upload(req.body.content, req.body.content.length);

  // SAS en lecture seule, valable 1 heure, au lieu d'un accès public permanent.
  const sasToken = generateBlobSASQueryParameters({
    containerName: CONTAINER,
    blobName: req.body.filename,
    permissions: BlobSASPermissions.parse('r'),
    expiresOn: new Date(Date.now() + 60 * 60 * 1000),
  }, blockBlobClient.credential).toString();

  res.json({ url: `${blockBlobClient.url}?${sasToken}` });
});

module.exports = router;
