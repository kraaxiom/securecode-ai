<?php

// Correction : utilisation du query builder Eloquent, qui lie automatiquement
// le paramètre `name` au lieu de le concaténer dans la requête SQL. La valeur
// de vérité de la clause WHERE ne peut plus être altérée par l'entrée
// utilisateur, ce qui élimine le canal d'inférence booléen.

namespace App\Http\Controllers;

use App\Models\Product;
use Illuminate\Http\Request;

class ProductSearchController extends Controller
{
    public function search(Request $request)
    {
        $name = $request->input('name');

        $products = Product::where('name', $name)->get();

        return response()->json($products);
    }
}
