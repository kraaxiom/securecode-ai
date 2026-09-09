<?php
// Faille : l'algorithme de vérification n'est pas explicitement restreint,
// laissant potentiellement passer un token signé avec 'alg: none'. Un
// attaquant pourrait alors forger un token dont il contrôle intégralement
// le contenu (y compris le rôle) sans posséder aucun secret.

namespace App\Http\Middleware;

use Closure;
use Firebase\JWT\JWT;
use Firebase\JWT\Key;
use Illuminate\Http\Request;

class VerifySessionToken
{
    public function handle(Request $request, Closure $next)
    {
        $token = $request->bearerToken();
        $secret = config('services.jwt.secret');

        // Algorithme vide : rien n'exclut explicitement 'none'.
        $decoded = JWT::decode($token, new Key($secret, ''));

        if ($decoded->role === 'admin') {
            $request->attributes->set('is_admin', true);
        }

        return $next($request);
    }
}
