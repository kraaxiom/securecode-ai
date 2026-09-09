<?php
// Faille : aucune détection de trafic distribué (nombreux comptes différents
// testés depuis une même origine) et aucune proposition de MFA. Un attaquant
// utilisant des identifiants issus de fuites d'autres services peut réussir
// dès qu'une fraction des couples testés est valide, sans être repéré.

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
