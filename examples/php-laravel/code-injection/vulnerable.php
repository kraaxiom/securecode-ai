<?php

// Faille : l'expression fournie par l'utilisateur est directement évaluée
// par eval(). Cela donne à l'attaquant un contrôle total sur le code PHP
// exécuté par le serveur (injection de code), bien au-delà d'un simple
// calcul de formule.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class FormulaController extends Controller
{
    public function compute(Request $request)
    {
        $expr = $request->input('formula');

        $result = eval("return $expr;");

        return response()->json(['result' => $result]);
    }
}
