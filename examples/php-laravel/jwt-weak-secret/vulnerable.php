<?php
// Faille : le secret utilisé pour signer les tokens JWT est court, codé en
// dur et prévisible. Un attaquant peut le retrouver hors ligne par attaque
// par dictionnaire, puis forger des tokens valides avec n'importe quel
// contenu, y compris des rôles privilégiés.

namespace App\Services;

use Firebase\JWT\JWT;

class ApiTokenService
{
    private string $secret = 'secret123';

    public function issueToken(array $payload): string
    {
        $payload['iat'] = time();
        $payload['exp'] = time() + 3600;

        return JWT::encode($payload, $this->secret, 'HS256');
    }
}
