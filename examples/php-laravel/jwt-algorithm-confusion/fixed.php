<?php
// Correction : l'algorithme de vérification est explicitement figé à
// RS256, excluant structurellement toute confusion avec HS256, conformément
// au pattern décrit dans rules/remediation/jwt-algorithm-confusion.md.

namespace App\Http\Middleware;

use Closure;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Illuminate\Http\Request;

class VerifyApiToken
{
    public function handle(Request $request, Closure $next)
    {
        $token = $request->bearerToken();
        $publicKey = file_get_contents(storage_path('keys/public.pem'));

        // Uniquement RS256 accepté ; la clé publique RSA ne peut jamais
        // être réutilisée comme secret HMAC potentiel.
        $decoded = JWT::decode($token, new Key($publicKey, 'RS256'));

        $request->attributes->set('jwt_claims', $decoded);

        return $next($request);
    }
}
