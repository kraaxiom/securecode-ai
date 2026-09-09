<?php
// Correction : détection de vélocité globale par IP (tous comptes confondus)
// avec alerte de sécurité, et imposition de la MFA après authentification
// réussie, conformément au pattern décrit dans
// rules/remediation/credential-stuffing.md.

namespace App\Http\Controllers\Auth;

use App\Events\SecurityAlert;
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

        $ipKey = 'login-attempts:' . $request->ip();

        if (RateLimiter::tooManyAttempts($ipKey, 30)) {
            // Volume anormal tous comptes confondus depuis cette origine.
            SecurityAlert::dispatch('credential_stuffing_suspected', $request->ip());
            return back()->withErrors([
                'email' => 'Trop de tentatives depuis cette origine.',
            ]);
        }
        RateLimiter::hit($ipKey, 600);

        if (Auth::attempt($request->only('email', 'password'))) {
            if ($request->user()->mfa_enabled) {
                return redirect()->route('mfa.challenge');
            }
            return redirect()->intended();
        }

        return back()->withErrors([
            'email' => 'Identifiants invalides',
        ]);
    }
}
