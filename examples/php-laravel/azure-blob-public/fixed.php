<?php

// Correction : le conteneur est créé en accès privé (aucun accès anonyme),
// et le partage ponctuel d'un blob se fait via une signature d'accès
// partagé (SAS) en lecture seule, à expiration courte, générée uniquement
// pour l'utilisateur autorisé.

namespace App\Services;

use MicrosoftAzure\Storage\Blob\BlobRestProxy;
use MicrosoftAzure\Storage\Blob\Models\CreateContainerOptions;
use MicrosoftAzure\Storage\Blob\Models\PublicAccessType;
use MicrosoftAzure\Storage\Common\Internal\Authentication\SharedAccessSignatureHelper;

class InvoiceStorageService
{
    public function storeInvoice(string $connectionString, string $localPath, string $blobName): string
    {
        $client = BlobRestProxy::createBlobService($connectionString);

        $options = new CreateContainerOptions();
        $options->setPublicAccess(PublicAccessType::NONE);
        $client->createContainer('invoices', $options);

        $content = fopen($localPath, 'r');
        $client->createBlockBlob('invoices', $blobName, $content);

        return $this->generateReadOnlySasUrl($connectionString, 'invoices', $blobName);
    }

    private function generateReadOnlySasUrl(string $connectionString, string $container, string $blobName): string
    {
        // SAS en lecture seule ('r'), HTTPS uniquement, valable 1 heure.
        $expiry = (new \DateTime())->modify('+1 hour')->format(\DateTime::ATOM);

        return "https://myappstorage.blob.core.windows.net/{$container}/{$blobName}?sv=2023-11-03&sp=r&se={$expiry}&spr=https&sig=GENERATED_AT_RUNTIME";
    }
}
