<?php

// Correction : une liste blanche explicite des champs autorisés est appliquée
// via la validation de la requête. Les champs sensibles comme 'role' ou
// 'is_admin' ne peuvent jamais transiter par cette route, quel que soit le
// contenu du JSON envoyé par le client. En complément, le modèle User
// devrait définir $fillable = ['name', 'email'] en défense en profondeur.

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class ProfileController extends Controller
{
    public function update(Request $request, $id)
    {
        $user = User::findOrFail($id);

        $validated = $request->validate([
            'name' => 'sometimes|string|max:255',
            'email' => 'sometimes|email',
        ]);

        $user->update($validated);

        return response()->json($user);
    }
}

// routes/web.php
// Route::put('/profile/{id}', [ProfileController::class, 'update'])->middleware('auth');
