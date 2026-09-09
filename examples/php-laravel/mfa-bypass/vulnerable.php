<?php
// Faille : la session est pleinement authentifiée dès la vérification du
// mot de passe, avant toute validation du second facteur. Un attaquant en
// possession du mot de passe seul obtient un accès complet sans jamais
// avoir à franchir l'étape MFA.

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;

class LoginController extends Controller
{
    public function login(Request $request)
    {
        if (Auth::attempt($request->only('email', 'password'))) {
            return redirect()->route('dashboard'); // accès complet avant MFA !
        }

        return back()->withErrors([
            'email' => 'Identifiants invalides',
        ]);
    }
}
