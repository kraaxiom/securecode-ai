<?php

// Correction : la requête utilise un paramètre lié via un point
// d'interrogation, ce qui élimine toute possibilité pour l'entrée
// utilisateur de modifier la structure de la requête SQL.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class UserSearchController extends Controller
{
    public function search(Request $request)
    {
        $name = $request->string('name')->toString();

        // Requête préparée avec paramètre lié : aucune concaténation.
        $users = DB::select('SELECT id, name, email FROM users WHERE name = ?', [$name]);

        return response()->json($users);
    }
}
