<?php

// Faille : le service applicatif assume un rôle IAM AWS via des
// identifiants statiques disposant d'une policy 'Action: *, Resource: *'
// (droits d'administrateur complets). En cas de fuite de ces identifiants
// ou de compromission du service, l'attaquant hérite d'un contrôle total
// sur l'ensemble du compte cloud, bien au-delà du besoin réel de la
// fonctionnalité (lecture de fichiers S3).

namespace App\Services;

use Aws\S3\S3Client;

class ReportExportService
{
    public function fetchExport(string $key): string
    {
        // Identifiants statiques d'un utilisateur IAM disposant d'une
        // policy '*'/'*' (admin complet) au lieu d'un rôle scoped.
        $s3 = new S3Client([
            'version'     => 'latest',
            'region'      => 'eu-west-1',
            'credentials' => [
                'key'    => config('services.aws.admin_key'),
                'secret' => config('services.aws.admin_secret'),
            ],
        ]);

        $result = $s3->getObject(['Bucket' => 'my-app-bucket', 'Key' => $key]);

        return (string) $result['Body'];
    }
}
