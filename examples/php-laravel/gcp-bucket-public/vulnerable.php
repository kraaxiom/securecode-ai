<?php

// Faille : après l'upload, le code accorde explicitement le rôle
// 'roles/storage.objectViewer' à 'allUsers' sur l'objet, rendant le fichier
// lisible publiquement par n'importe qui sur Internet, sans authentification
// ni expiration.

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

        // Rend l'objet public à tous, de façon permanente.
        $object->acl()->add('allUsers', 'READER');

        return "https://storage.googleapis.com/my-app-bucket/{$objectName}";
    }
}
