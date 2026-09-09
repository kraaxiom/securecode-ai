<?php

// Faille : la valeur "next" fournie par l'utilisateur est utilisée telle
// quelle pour construire un en-tête Location via l'en-tête bas niveau du
// Response Laravel. Une entrée contenant des séquences CR/LF peut injecter
// des en-têtes supplémentaires ou tronquer la réponse HTTP (injection CRLF).

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\Response;

class RedirectController extends Controller
{
    public function redirectTo(Request $request)
    {
        $next = $request->input('next');

        $response = new Response('', 302);
        $response->headers->set('Location', $next);

        return $response;
    }
}
