<?php

// Faille : Mass Assignment.
// Le corps entier de la requête est passé tel quel à update(), sans liste
// blanche de champs autorisés. Un attaquant peut injecter des champs
// supplémentaires comme 'role' ou 'is_admin' dans le JSON envoyé, modifiant
// ainsi des attributs sensibles qui ne devraient pas être modifiables via
// ce formulaire de profil.

namespace App\Http\Controllers;

use App\Models\User;
use Illuminate\Http\Request;

class ProfileController extends Controller
{
    public function update(Request $request, $id)
    {
        $user = User::findOrFail($id);
        $user->update($request->all());

        return response()->json($user);
    }
}

// routes/web.php
// Route::put('/profile/{id}', [ProfileController::class, 'update'])->middleware('auth');
