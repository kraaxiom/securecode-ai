<?php

// Correction : requête préparée avec paramètre lié (élimine la classe de
// vulnérabilité, pas seulement la variante time-based), "id" est casté en
// entier, et un timeout d'exécution est appliqué au niveau de la connexion
// PDO pour limiter l'impact d'une éventuelle injection résiduelle ailleurs.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use PDO;
use Throwable;

class OrderStatusController extends Controller
{
    public function show(Request $request)
    {
        $id = $request->integer('id');

        try {
            DB::connection()->getPdo()->setAttribute(PDO::ATTR_TIMEOUT, 5);
            DB::select('SELECT status FROM orders WHERE id = ?', [$id]);
        } catch (Throwable $e) {
            // Message générique conservé, mais la requête n'est plus injectable.
        }

        return response()->json(['message' => 'Requête traitée']);
    }
}
