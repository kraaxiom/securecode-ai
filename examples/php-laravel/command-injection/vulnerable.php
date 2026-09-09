<?php

// Faille : l'hôte fourni par l'utilisateur est concaténé dans une commande
// shell exécutée via system(). Un attaquant peut insérer des métacaractères
// shell (';', '&&', '|', etc.) pour exécuter des commandes arbitraires avec
// les privilèges du processus applicatif.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class NetworkDiagnosticController extends Controller
{
    public function ping(Request $request)
    {
        $host = $request->input('host');

        $output = null;
        system('ping -c 3 ' . $host, $exitCode);

        return response()->json(['output' => $output, 'exit_code' => $exitCode]);
    }
}
