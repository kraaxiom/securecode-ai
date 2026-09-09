<?php

// CORRIGÉ — DOM XSS (CWE-79)
// La donnée de requête n'est plus concaténée directement dans le script
// inline : elle est sérialisée en JSON échappé via json_encode() avec les
// flags JSON_HEX_*, ce qui empêche toute rupture hors du littéral
// JavaScript ou du contexte HTML environnant.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class SearchController extends Controller
{
    public function results(Request $request)
    {
        $term = $request->query('q');

        // Passage par JSON échappé — plus de concaténation directe dans le script.
        $safeTerm = json_encode($term, JSON_HEX_TAG | JSON_HEX_APOS | JSON_HEX_QUOT | JSON_HEX_AMP);

        return response(
            "<script>var searchTerm = " . $safeTerm . ";</script>" .
            "<div id='results'></div>"
        )->header('Content-Type', 'text/html');
    }
}
