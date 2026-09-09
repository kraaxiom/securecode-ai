<?php

// Faille : le conteneur Azure Blob Storage cible est créé/utilisé avec un
// niveau d'accès public 'container', ce qui permet à n'importe qui de lister
// et lire tous les blobs du conteneur via une simple URL HTTP, sans SAS ni
// authentification.

namespace App\Services;

use MicrosoftAzure\Storage\Blob\BlobRestProxy;
use MicrosoftAzure\Storage\Blob\Models\CreateContainerOptions;
use MicrosoftAzure\Storage\Blob\Models\PublicAccessType;

class InvoiceStorageService
{
    public function storeInvoice(string $connectionString, string $localPath, string $blobName): string
    {
        $client = BlobRestProxy::createBlobService($connectionString);

        $options = new CreateContainerOptions();
        $options->setPublicAccess(PublicAccessType::CONTAINER_AND_BLOBS);
        $client->createContainer('invoices', $options);

        $content = fopen($localPath, 'r');
        $client->createBlockBlob('invoices', $blobName, $content);

        return "https://myappstorage.blob.core.windows.net/invoices/{$blobName}";
    }
}
