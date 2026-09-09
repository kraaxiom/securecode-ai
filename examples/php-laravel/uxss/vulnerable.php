<?php

// VULNÉRABLE — Universal XSS / UXSS (CWE-79)
// L'application n'émet aucun en-tête Content-Security-Policy ni
// Permissions-Policy. Elle intègre par ailleurs un widget tiers (chat
// support) via une iframe sans sandbox et un script CDN sans intégrité
// vérifiée, augmentant la surface d'exposition à une faille UXSS provenant
// d'un composant tiers.

namespace App\Http\Middleware;

use Closure;
use Illuminate\Http\Request;

class SecurityHeaders
{
    public function handle(Request $request, Closure $next)
    {
        // Aucun en-tête de sécurité n'est défini.
        return $next($request);
    }
}

// resources/views/layouts/app.blade.php
// <script src="https://cdn.example.com/widget.js"></script>
// <iframe src="https://widget-tiers.example.com/chat"></iframe>
