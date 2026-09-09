<?php

// CORRIGÉ — Universal XSS / UXSS (CWE-79)
// Un middleware émet une Content-Security-Policy stricte restreignant
// script-src et frame-src aux origines nécessaires, ainsi qu'une
// Permissions-Policy limitant les capacités des composants tiers. Le script
// CDN utilise désormais Subresource Integrity et l'iframe tierce un
// attribut sandbox minimal (voir la vue associée).

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;

class SecurityHeaders
{
    public function handle(Request $request, Closure $next)
    {
        $response = $next($request);

        $response->headers->set(
            'Content-Security-Policy',
            "default-src 'self'; script-src 'self' https://cdn.example.com; " .
            "frame-src 'self' https://widget-tiers.example.com"
        );
        $response->headers->set(
            'Permissions-Policy',
            'camera=(), microphone=(), geolocation=()'
        );

        return $response;
    }
}

// resources/views/layouts/app.blade.php
// <script src="https://cdn.example.com/widget.js"
//         integrity="sha384-<hash-du-fichier-verifie>"
//         crossorigin="anonymous"></script>
// <iframe src="https://widget-tiers.example.com/chat"
//         sandbox="allow-scripts allow-same-origin"
//         referrerpolicy="no-referrer"></iframe>
