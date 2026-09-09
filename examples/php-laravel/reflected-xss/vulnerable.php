<?php

// VULNÉRABLE — Reflected XSS (CWE-79)
// Le terme de recherche saisi par l'utilisateur est réinjecté dans la vue
// Blade avec l'échappement automatique explicitement désactivé ({!! !!}),
// ce qui permet à un lien piégé contenant du HTML/JS malveillant dans le
// paramètre "q" de s'exécuter dans le navigateur de la victime.

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
// <h1>Résultats pour : {!! $term !!}</h1>
