<?php

// Correction : requête préparée avec paramètre lié, et l'émulation de
// préparation PDO est explicitement désactivée pour empêcher le driver
// d'accepter du multi-statement construit côté client. Un point-virgule
// dans "name" est traité comme une simple donnée, pas comme un séparateur
// d'instructions.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use PDO;

class ProfileController extends Controller
{
    public function updateName(Request $request)
    {
        $name = $request->string('name')->toString();

        // Désactive l'émulation de préparation : empêche le multi-statement côté client.
        DB::connection()->getPdo()->setAttribute(PDO::ATTR_EMULATE_PREPARES, false);

        DB::statement('UPDATE users SET name = ? WHERE id = 1', [$name]);

        return response()->json(['message' => 'Profil mis à jour']);
    }
}
