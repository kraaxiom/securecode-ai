<?php

// Correction : "id" est casté en entier puis passé comme paramètre lié
// dans la requête préparée. Une valeur non numérique ne peut plus altérer
// la structure de la requête ni y injecter une clause UNION SELECT.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class ProductController extends Controller
{
    public function show(Request $request)
    {
        $id = $request->integer('id');

        // Requête préparée avec paramètre lié, valeur strictement typée en entier.
        $product = DB::select('SELECT * FROM products WHERE id = ?', [$id]);

        return response()->json($product);
    }
}
