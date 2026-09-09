<?php

// Faille : Privilege Escalation (escalade de privilèges verticale).
// La fonction d'attribution de rôle accepte n'importe quelle valeur envoyée
// par le client sans vérifier que l'appelant est lui-même autorisé à
// accorder ce niveau de privilège. Un utilisateur standard pourrait ainsi
// s'auto-attribuer (ou attribuer à un tiers) le rôle 'admin'.

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class RoleController extends Controller
{
    public function assignRole(Request $request, $userId)
    {
        $user = User::findOrFail($userId);
        $user->role = $request->input('role');
        $user->save();

        return response()->json($user);
    }
}

// routes/web.php
// Route::put('/users/{userId}/role', [RoleController::class, 'assignRole'])->middleware('auth');
