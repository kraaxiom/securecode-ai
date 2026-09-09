<?php

// Correction : requête préparée avec paramètre lié, y compris pour cette
// requête conditionnelle dont le résultat n'est pas affiché directement.
// Cela empêche toute injection de condition SQL, éliminant le canal
// d'inférence booléen exploité par l'injection SQL aveugle.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class AccountStatusController extends Controller
{
    public function checkActive(Request $request)
    {
        $username = $request->input('username');

        $result = DB::select(
            'SELECT 1 FROM users WHERE username = :username AND active = 1',
            ['username' => $username]
        );

        $exists = count($result) > 0;

        return response()->json(['active' => $exists]);
    }
}
