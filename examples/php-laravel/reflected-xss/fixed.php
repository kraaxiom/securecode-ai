<?php

// CORRIGÉ — Reflected XSS (CWE-79)
// L'échappement automatique de Blade est conservé ({{ }} au lieu de {!! !!}),
// ce qui encode contextuellement le terme de recherche avant affichage et
// neutralise toute tentative d'injection HTML/JS via le paramètre "q".

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class SearchController extends Controller
{
    public function index(Request $request)
    {
        $term = $request->query('q', '');

        return view('search.results', ['term' => $term]);
    }
}

// resources/views/search/results.blade.php
// <h1>Résultats pour : {{ $term }}</h1>
