<?php

// Correction : la regex est réécrite sans quantificateurs imbriqués
// (motif non ambigu), une limite de longueur est imposée sur l'entrée
// avant l'application de la regex, et une limite de backtracking PCRE
// est fixée pour protéger l'exécution sur une entrée non fiable.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class NewsletterController extends Controller
{
    public function subscribe(Request $request)
    {
        // Limite de longueur imposée avant tout traitement par regex (RFC 5321: 254 max).
        $email = substr((string) $request->input('email', ''), 0, 254);

        ini_set('pcre.backtrack_limit', '100000');

        // Regex simplifiée, sans quantificateurs imbriqués ni ambiguïté de correspondance.
        if (!preg_match('/^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/', $email)) {
            return response()->json(['error' => 'Email invalide'], 422);
        }

        // ... enregistrement de l'abonné
        return response()->json(['message' => 'Abonnement confirmé']);
    }
}
