<?php

// Faille : le paramètre "name" fourni par l'utilisateur est concaténé
// directement dans le TEXTE du template avant sa compilation par Twig.
// Le moteur interprète alors le contenu injecté comme de la syntaxe de
// template légitime, ce qui peut mener à une exécution de code arbitraire.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Twig\Environment;
use Twig\Loader\ArrayLoader;

class WelcomeController extends Controller
{
    public function greet(Request $request)
    {
        $name = $request->input('name');

        $twig = new Environment(new ArrayLoader());

        // Le template lui-même est construit dynamiquement à partir de l'entrée utilisateur.
        $templateSource = 'Bonjour ' . $name . ', bienvenue !';
        $output = $twig->createTemplate($templateSource)->render([]);

        return response($output);
    }
}
