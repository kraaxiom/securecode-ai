<?php

// Faille : le paramètre "host" fourni par l'utilisateur est concaténé
// directement dans une commande shell exécutée via system(). Un attaquant
// peut injecter des méta-caractères shell (";", "|", "&&", backticks)
// pour exécuter des commandes arbitraires sur le serveur.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class NetworkDiagnosticController extends Controller
{
    public function ping(Request $request)
    {
        $host = $request->input('host');

        // Concaténation directe d'une entrée utilisateur dans un appel shell.
        $output = [];
        system('ping -c 4 ' . $host, $exitCode);

        return response()->json([
            'host' => $host,
            'exit_code' => $exitCode,
        ]);
    }
}
