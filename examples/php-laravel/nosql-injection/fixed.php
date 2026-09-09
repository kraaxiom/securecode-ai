<?php

// Correction : chaque champ attendu est strictement validé par type avant
// d'être transmis à la requête MongoDB. Tout objet/tableau/opérateur
// ($ne, $gt, $where...) envoyé à la place d'une chaîne scalaire est rejeté
// en amont, ce qui empêche l'attaquant de modifier la logique de la requête.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Validation\ValidationException;
use MongoDB\Client;

class LoginController extends Controller
{
    public function login(Request $request)
    {
        // Validation de schéma stricte : rejette tout ce qui n'est pas une chaîne.
        $validated = $request->validate([
            'username' => ['required', 'string', 'max:190'],
            'password' => ['required', 'string', 'max:190'],
        ]);

        $client = new Client(config('database.mongo_dsn'));
        $collection = $client->selectDatabase('app')->selectCollection('users');

        $username = (string) $validated['username'];
        $password = (string) $validated['password'];

        $user = $collection->findOne([
            'username' => $username,
            'password' => $password,
        ]);

        if (!$user) {
            return response()->json(['error' => 'Identifiants invalides'], 401);
        }

        return response()->json(['message' => 'Connecté', 'user' => (string) $user['_id']]);
    }
}
