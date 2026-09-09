<?php

// Faille : le fichier uploadé est stocké avec l'ACL 'public-read' codée en
// dur, ce qui rend l'objet accessible publiquement via son URL directe, sans
// aucune authentification. Combiné à une bucket policy ou un Block Public
// Access désactivé, n'importe quel attaquant connaissant ou devinant la clé
// de l'objet peut lire des documents potentiellement sensibles.

namespace App\Http\Controllers;

use Aws\S3\S3Client;
use Illuminate\Http\Request;

class DocumentUploadController extends Controller
{
    public function upload(Request $request)
    {
        $file = $request->file('document');

        $s3 = new S3Client([
            'version' => 'latest',
            'region'  => 'eu-west-1',
        ]);

        $result = $s3->putObject([
            'Bucket'     => 'my-app-bucket',
            'Key'        => 'uploads/' . $file->getClientOriginalName(),
            'Body'       => fopen($file->getRealPath(), 'rb'),
            'ACL'        => 'public-read',
        ]);

        return response()->json(['url' => $result['ObjectURL']]);
    }
}
