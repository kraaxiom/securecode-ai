<?php
// Correction : le secret provient d'une variable d'environnement, généré
// via un CSPRNG (random_bytes(32), soit 256 bits d'entropie) et stocké hors
// du code, conformément au pattern décrit dans
// rules/remediation/weak-jwt-secret.md.

namespace App\Services;

use Firebase\JWT\JWT;

class ApiTokenService
{
    private string $secret;

    public function __construct()
    {
        // Généré une fois via random_bytes(32) puis stocké dans le gestionnaire
        // de secrets ; jamais codé en dur dans le code source.
        $this->secret = (string) getenv('JWT_SECRET');
    }

    public function issueToken(array $payload): string
    {
        $payload['iat'] = time();
        $payload['exp'] = time() + 900; // expiration courte

        return JWT::encode($payload, $this->secret, 'HS256');
    }
}
