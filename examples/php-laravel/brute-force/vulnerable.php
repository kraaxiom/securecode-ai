<?php
// Faille : aucune limitation du nombre de tentatives de connexion.
// Un attaquant peut essayer un grand nombre de mots de passe de façon
// automatisée sur ce endpoint sans jamais être bloqué ni ralenti.

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class LoginController extends Controller
{
    public function login(Request $request)
    {
        $request->validate([
            'email' => 'required|email',
            'password' => 'required',
        ]);

        if (Auth::attempt($request->only('email', 'password'))) {
            return redirect()->intended();
        }

        return back()->withErrors([
            'email' => 'Identifiants invalides',
        ]);
    }
}
