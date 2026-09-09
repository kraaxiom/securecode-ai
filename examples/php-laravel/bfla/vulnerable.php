<?php

// Faille : Broken Function Level Authorization (BFLA).
// Le contrôleur expose une fonction d'administration (suppression d'un utilisateur)
// en la protégeant uniquement par le middleware d'authentification ('auth').
// Aucune vérification de rôle n'est effectuée : n'importe quel utilisateur connecté,
// même sans privilège admin, peut appeler cette action s'il en connaît l'URL,
// alors que le bouton correspondant n'est visible que dans l'interface admin.

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class UserController extends Controller
{
    public function destroy($id)
    {
        $user = User::findOrFail($id);
        $user->delete();

        return response()->json(['message' => 'Utilisateur supprimé']);
    }
}

// routes/web.php (ou api.php)
// Route::delete('/users/{id}', [UserController::class, 'destroy'])->middleware('auth');
