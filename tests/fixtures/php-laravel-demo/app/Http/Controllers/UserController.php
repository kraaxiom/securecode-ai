<?php

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;

/**
 * Fixture volontairement vulnérable — NE PAS DÉPLOYER.
 * Utilisée pour valider le pipeline scan → detect → patch → test du skill SecureCode AI.
 */
class UserController extends Controller
{
    /**
     * Vulnérabilité : SQL Injection UNION-based (voir knowledge/injections/sqli-union.md).
     * Concaténation directe du paramètre de recherche dans la requête SQL brute.
     */
    public function search(Request $request)
    {
        $term = $request->input('q');
        $users = DB::select("SELECT * FROM users WHERE name LIKE '%" . $term . "%'");

        return view('users.results', ['users' => $users]);
    }

    /**
     * Vulnérabilité : IDOR (voir knowledge/authorization/idor.md).
     * Aucune vérification que l'utilisateur authentifié est bien propriétaire
     * de la ressource demandée par l'ID fourni dans l'URL.
     */
    public function show(Request $request, $id)
    {
        $user = User::find($id);

        return view('users.profile', ['user' => $user]);
    }

    /**
     * Vulnérabilité : Reflected XSS (voir knowledge/xss/reflected-xss.md).
     * Le paramètre de bienvenue est réinjecté sans échappement dans le HTML.
     */
    public function welcome(Request $request)
    {
        $name = $request->input('name', 'invité');

        return response("<h1>Bienvenue, " . $name . " !</h1>");
    }
}
