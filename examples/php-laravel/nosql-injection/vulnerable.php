<?php

// Faille : le corps de la requête HTTP est transmis tel quel comme filtre
// à une collection MongoDB. Un attaquant peut envoyer un objet (ex: un
// opérateur "$ne") à la place d'une chaîne, ce qui permet de contourner
// la logique d'authentification en modifiant la structure de la requête.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use MongoDB\Client;

class LoginController extends Controller
{
    public function login(Request $request)
    {
        $client = new Client(config('database.mongo_dsn'));
        $collection = $client->selectDatabase('app')->selectCollection('users');

        $username = $request->input('username');
        $password = $request->input('password');

        // Aucune validation de type : $username / $password peuvent être
        // des tableaux/objets JSON interprétés comme des opérateurs Mongo.
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
