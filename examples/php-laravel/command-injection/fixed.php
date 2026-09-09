<?php

// Correction : validation stricte de l'entrée comme adresse IP (liste
// blanche de format), puis exécution via Symfony Process avec les arguments
// passés en tableau distinct, sans jamais transiter par un interpréteur
// shell.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use InvalidArgumentException;
use Symfony\Component\Process\Process;

class NetworkDiagnosticController extends Controller
{
    public function ping(Request $request)
    {
        $host = $request->input('host');

        if (!filter_var($host, FILTER_VALIDATE_IP)) {
            throw new InvalidArgumentException('Hôte invalide');
        }

        $process = new Process(['ping', '-c', '3', $host]);
        $process->run();

        return response()->json([
            'output' => $process->getOutput(),
            'exit_code' => $process->getExitCode(),
        ]);
    }
}
