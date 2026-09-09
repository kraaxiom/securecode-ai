<?php

// Correction : l'entrée est validée par liste blanche (format d'hôte
// strict) avant tout usage, puis échappée avec escapeshellarg(). Cela
// empêche l'injection de méta-caractères shell dans la commande exécutée.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use InvalidArgumentException;

class NetworkDiagnosticController extends Controller
{
    public function ping(Request $request)
    {
        $host = (string) $request->input('host');

        // Liste blanche stricte : adresse IP valide ou nom d'hôte simple uniquement.
        if (!filter_var($host, FILTER_VALIDATE_IP) && !preg_match('/^[a-zA-Z0-9.-]+$/', $host)) {
            throw new InvalidArgumentException('Hôte invalide');
        }

        $escapedHost = escapeshellarg($host);
        system('ping -c 4 ' . $escapedHost, $exitCode);

        return response()->json([
            'host' => $host,
            'exit_code' => $exitCode,
        ]);
    }
}
