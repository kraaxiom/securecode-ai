<?php
// Faille : l'algorithme de vérification n'est pas figé côté serveur, il est
// implicitement déduit du token présenté. Un attaquant peut alors forger un
// token signé en HS256 en réutilisant la clé publique RSA de l'application
// comme secret HMAC, trompant ainsi la vérification.

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

        // Algorithme non figé explicitement : la clé est passée sans
        // contraindre la famille d'algorithme acceptée.
        $decoded = JWT::decode($token, new Key($publicKey, null));

        $request->attributes->set('jwt_claims', $decoded);

        return $next($request);
    }
}
