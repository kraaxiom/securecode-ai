<?php

// Correction : l'objet reste privé (aucun binding IAM 'allUsers') et
// l'accès ponctuel se fait via une URL signée à durée de vie courte,
// générée uniquement pour l'utilisateur autorisé, cohérente avec un bucket
// configuré en 'uniform_bucket_level_access' et 'Public Access Prevention'
// activée.

namespace App\Services;

use Google\Cloud\Storage\StorageClient;

class ReportPublisherService
{
    public function publish(string $localPath, string $objectName): string
    {
        $storage = new StorageClient(['projectId' => 'my-project']);
        $bucket = $storage->bucket('my-app-bucket');

        $object = $bucket->upload(fopen($localPath, 'r'), [
            'name' => $objectName,
        ]);

        // Pas de binding public : l'accès est accordé au niveau IAM du
        // bucket à des service accounts précis (roles/storage.objectViewer
        // scoped, hors code applicatif).
        return $object->signedUrl(new \DateTime('+30 minutes'));
    }
}
