<?php

// Correction : le texte du template est fixe et versionné (aucune
// concaténation d'entrée utilisateur). La donnée utilisateur est passée
// uniquement comme variable de contexte, que Twig échappe automatiquement,
// ce qui empêche son interprétation comme syntaxe de template.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Twig\Environment;
use Twig\Loader\ArrayLoader;

class WelcomeController extends Controller
{
    public function greet(Request $request)
    {
        $name = $request->string('name')->toString();

        // Template statique, contrôlé par le développeur, jamais construit dynamiquement.
        $loader = new ArrayLoader([
            'welcome' => 'Bonjour {{ name }}, bienvenue !',
        ]);
        $twig = new Environment($loader);

        // "name" transite uniquement comme variable de contexte.
        $output = $twig->render('welcome', ['name' => $name]);

        return response($output);
    }
}
