<?php

// Correction : rejet de toute entrée contenant des caractères \r ou \n et
// contrainte à un chemin relatif commençant par '/', avant utilisation dans
// l'en-tête Location. Aucune séquence CRLF ne peut plus être injectée dans
// la réponse HTTP.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class RedirectController extends Controller
{
    public function redirectTo(Request $request)
    {
        $next = $request->input('next', '/');

        if (preg_match('/[\r\n]/', $next) || !str_starts_with($next, '/')) {
            $next = '/';
        }

        return redirect($next, 302);
    }
}
