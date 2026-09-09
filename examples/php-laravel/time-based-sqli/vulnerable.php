<?php

// Faille : le paramètre "id" est concaténé directement dans la requête
// SQL, et les erreurs sont masquées (réponse générique). Un attaquant ne
// voit ni donnée ni message d'erreur, mais peut injecter une condition
// provoquant une pause (ex: SLEEP) et déduire des informations en
// observant le délai de réponse (injection SQL aveugle basée sur le temps).

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use Throwable;

class OrderStatusController extends Controller
{
    public function show(Request $request)
    {
        $id = $request->input('id');

        try {
            // Concaténation directe, aucun timeout de requête configuré.
            DB::select("SELECT status FROM orders WHERE id = " . $id);
        } catch (Throwable $e) {
            // Message générique : aucune donnée exploitée, mais la faille reste présente.
        }

        return response()->json(['message' => 'Requête traitée']);
    }
}
