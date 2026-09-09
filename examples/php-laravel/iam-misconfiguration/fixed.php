<?php

// Correction : le service utilise des identifiants temporaires obtenus via
// AssumeRole sur un rôle IAM scoped, limité aux seules actions
// 's3:GetObject' sur le préfixe strictement nécessaire, avec une durée de
// vie courte, au lieu de clés statiques à privilèges larges.

namespace App\Services;

use Aws\S3\S3Client;
use Aws\Sts\StsClient;

class ReportExportService
{
    public function fetchExport(string $key): string
    {
        $sts = new StsClient(['version' => 'latest', 'region' => 'eu-west-1']);

        // Rôle applicatif scoped ("app-read-role"), pas un rôle admin,
        // avec des identifiants temporaires valables 15 minutes.
        $assumed = $sts->assumeRole([
            'RoleArn'         => 'arn:aws:iam::123456789012:role/app-read-role',
            'RoleSessionName' => 'report-export',
            'DurationSeconds' => 900,
        ]);

        $creds = $assumed['Credentials'];

        $s3 = new S3Client([
            'version'     => 'latest',
            'region'      => 'eu-west-1',
            'credentials' => [
                'key'    => $creds['AccessKeyId'],
                'secret' => $creds['SecretAccessKey'],
                'token'  => $creds['SessionToken'],
            ],
        ]);

        $result = $s3->getObject(['Bucket' => 'my-app-bucket', 'Key' => $key]);

        return (string) $result['Body'];
    }
}
