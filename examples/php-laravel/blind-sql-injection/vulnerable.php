<?php

// Faille : l'entrée utilisateur est concaténée directement dans une requête SQL
// brute. Même si le résultat n'est jamais affiché à l'utilisateur (seul un
// booléen "actif/inactif" est renvoyé), un attaquant peut injecter des
// conditions SQL pour faire varier la réponse et en déduire des données bit
// par bit (injection SQL aveugle).

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class AccountStatusController extends Controller
{
    public function checkActive(Request $request)
    {
        $username = $request->input('username');

        $result = DB::select(
            "SELECT 1 FROM users WHERE username = '$username' AND active = 1"
        );

        $exists = count($result) > 0;

        return response()->json(['active' => $exists]);
    }
}
