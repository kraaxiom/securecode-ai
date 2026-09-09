<?php
// Correction : algorithme fort explicitement imposé (HS256), rendant
// structurellement impossible l'acceptation d'un token 'alg: none',
// conformément au pattern décrit dans rules/remediation/jwt-none.md.

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

        // 'none' impossible à passer : l'algorithme attendu est explicite.
        $decoded = JWT::decode($token, new Key($secret, 'HS256'));

        if ($decoded->role === 'admin') {
            $request->attributes->set('is_admin', true);
        }

        return $next($request);
    }
}
