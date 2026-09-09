<?php
// Correction : limitation des tentatives par compte ET par IP, avec fenêtre
// glissante et journalisation implicite via RateLimiter, conformément au
// pattern décrit dans rules/remediation/brute-force.md.

namespace App\Http\Controllers\Auth;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\RateLimiter;

class LoginController extends Controller
{
    public function login(Request $request)
    {
        $request->validate([
            'email' => 'required|email',
            'password' => 'required',
        ]);

        $key = 'login:' . $request->ip() . ':' . strtolower($request->input('email'));

        if (RateLimiter::tooManyAttempts($key, 5)) {
            $seconds = RateLimiter::availableIn($key);
            return back()->withErrors([
                'email' => "Trop de tentatives. Réessayez dans {$seconds}s.",
            ]);
        }

        if (Auth::attempt($request->only('email', 'password'))) {
            RateLimiter::clear($key);
            return redirect()->intended();
        }

        RateLimiter::hit($key, 900); // fenêtre de 15 minutes

        return back()->withErrors([
            'email' => 'Identifiants invalides',
        ]);
    }
}
