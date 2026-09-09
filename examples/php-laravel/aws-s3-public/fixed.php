<?php

// Correction : le fichier est stocké avec l'ACL 'private' (accès refusé par
// défaut, en cohérence avec un Block Public Access activé côté compte et
// bucket), et le partage ponctuel se fait via une URL pré-signée à
// expiration courte, générée uniquement pour l'utilisateur autorisé.

namespace App\Http\Controllers;

use Aws\S3\S3Client;
use Illuminate\Http\Request;

class DocumentUploadController extends Controller
{
    public function upload(Request $request)
    {
        $file = $request->file('document');
        $key = 'uploads/' . $file->hashName();

        $s3 = new S3Client([
            'version' => 'latest',
            'region'  => 'eu-west-1',
        ]);

        $s3->putObject([
            'Bucket' => 'my-app-bucket',
            'Key'    => $key,
            'Body'   => fopen($file->getRealPath(), 'rb'),
            'ACL'    => 'private',
        ]);

        $command = $s3->getCommand('GetObject', [
            'Bucket' => 'my-app-bucket',
            'Key'    => $key,
        ]);
        $presignedRequest = $s3->createPresignedRequest($command, '+15 minutes');

        return response()->json(['url' => (string) $presignedRequest->getUri()]);
    }
}
