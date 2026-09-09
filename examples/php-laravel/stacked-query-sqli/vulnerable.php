<?php

// Faille : le nom fourni par l'utilisateur est concaténé dans une requête
// exécutée avec l'émulation de préparation activée par défaut par PDO,
// ce qui autorise l'exécution multi-instructions. Un attaquant peut
// ajouter un point-virgule suivi d'une instruction SQL distincte
// (ex: INSERT, UPDATE, DROP) qui s'exécutera à la suite de la requête.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class ProfileController extends Controller
{
    public function updateName(Request $request)
    {
        $name = $request->input('name');

        // Concaténation directe + multi-statement non désactivé côté driver.
        DB::statement("UPDATE users SET name = '{$name}' WHERE id = 1");

        return response()->json(['message' => 'Profil mis à jour']);
    }
}
