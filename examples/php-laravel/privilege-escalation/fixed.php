<?php

// Correction : l'appelant doit être explicitement autorisé à accorder le rôle
// demandé (canGrantRole), et les sessions/tokens existants de l'utilisateur
// modifié sont invalidés après le changement de rôle afin d'éviter que
// d'anciens privilèges ne persistent dans un token déjà émis.

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class RoleController extends Controller
{
    public function assignRole(Request $request, $userId)
    {
        $requestedRole = $request->input('role');

        if (! auth()->user()->canGrantRole($requestedRole)) {
            abort(403, 'Vous ne pouvez pas attribuer ce niveau de privilège.');
        }

        $user = User::findOrFail($userId);
        $user->role = $requestedRole;
        $user->save();
        $user->tokens()->delete(); // invalide les sessions/tokens existants

        return response()->json($user);
    }
}

// routes/web.php
// Route::put('/users/{userId}/role', [RoleController::class, 'assignRole'])->middleware('auth');
