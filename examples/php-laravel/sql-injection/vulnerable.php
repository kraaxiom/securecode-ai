<?php

// Faille : le paramètre "name" fourni par l'utilisateur est concaténé
// directement dans la requête SQL exécutée via DB::select(). Un attaquant
// peut injecter un fragment SQL (ex: quote suivie d'une condition) pour
// altérer la logique de la requête.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class UserSearchController extends Controller
{
    public function search(Request $request)
    {
        $name = $request->input('name');

        // Concaténation directe d'une entrée utilisateur dans la requête SQL.
        $users = DB::select("SELECT id, name, email FROM users WHERE name = '" . $name . "'");

        return response()->json($users);
    }
}
