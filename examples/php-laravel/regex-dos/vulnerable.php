<?php

// Faille : la regex de validation d'email utilise des quantificateurs
// imbriqués ("+)+") appliqués à une entrée non bornée en taille. Une
// chaîne pathologique (ex: une longue suite de caractères sans "@" final)
// peut provoquer un temps de calcul exponentiel côté serveur (ReDoS).

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class NewsletterController extends Controller
{
    public function subscribe(Request $request)
    {
        $email = $request->input('email');

        // Quantificateurs imbriqués + aucune limite de longueur en amont.
        if (!preg_match('/^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+$/', $email)) {
            return response()->json(['error' => 'Email invalide'], 422);
        }

        // ... enregistrement de l'abonné
        return response()->json(['message' => 'Abonnement confirmé']);
    }
}
