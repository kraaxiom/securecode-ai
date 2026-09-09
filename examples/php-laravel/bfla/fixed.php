<?php

// Correction : une policy Laravel vérifie explicitement le rôle de l'appelant
// avant d'exécuter l'action sensible, en plus du middleware d'authentification.
// La vérification est centralisée dans UserPolicy et appliquée via le middleware
// 'can', qu'importe que la route soit visible ou non dans l'interface utilisateur.

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class UserController extends Controller
{
    public function destroy($id)
    {
        $user = User::findOrFail($id);

        // La policy a déjà été vérifiée par le middleware 'can:delete,App\Models\User'
        // mais on peut aussi l'appeler explicitement en défense en profondeur.
        $this->authorize('delete', $user);

        $user->delete();

        return response()->json(['message' => 'Utilisateur supprimé']);
    }
}

// app/Policies/UserPolicy.php
// class UserPolicy
// {
//     public function delete(User $actor, User $target): bool
//     {
//         return $actor->role === 'admin';
//     }
// }

// routes/web.php (ou api.php)
// Route::delete('/users/{id}', [UserController::class, 'destroy'])
//     ->middleware(['auth', 'can:delete,App\Models\User']);
