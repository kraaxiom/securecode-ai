<?php

// Faille : le nom de produit est concaténé directement dans la clause WHERE
// d'une requête SQL brute. L'attaquant peut modifier la valeur de vérité de
// la condition (ex: injecter un opérateur logique) et observer une
// différence binaire dans le résultat retourné (liste vide vs non vide),
// permettant une extraction de données caractère par caractère.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

class ProductSearchController extends Controller
{
    public function search(Request $request)
    {
        $name = $request->input('name');

        $products = DB::select("SELECT * FROM products WHERE name = '$name'");

        return response()->json($products);
    }
}
