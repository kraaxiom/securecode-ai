<?php

// Faille : le paramètre "id" issu de la query string est concaténé
// directement dans la requête SQL. Un attaquant peut ajouter une clause
// UNION SELECT pour extraire des données d'autres tables de la base.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class ProductController extends Controller
{
    public function show(Request $request)
    {
        $id = $request->input('id');

        // Concaténation directe : "id" n'est jamais casté ni validé en amont.
        $product = DB::select("SELECT * FROM products WHERE id = " . $id);

        return response()->json($product);
    }
}
