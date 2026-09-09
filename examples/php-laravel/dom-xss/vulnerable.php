<?php

// VULNÉRABLE — DOM XSS (CWE-79)
// Bien que le DOM XSS soit typiquement une faille purement côté client, le
// serveur peut y contribuer en injectant une valeur de requête non fiable
// dans un bloc <script> inline par simple concaténation de chaînes. Le
// navigateur exécute alors ce JavaScript, offrant un vecteur d'injection
// équivalent à un DOM XSS côté client.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class SearchController extends Controller
{
    public function results(Request $request)
    {
        $term = $request->query('q');

        // Concaténation directe d'une donnée de requête dans un script inline.
        return response(
            "<script>var searchTerm = '" . $term . "';</script>" .
            "<div id='results'></div>"
        )->header('Content-Type', 'text/html');
    }
}
