<?php

// VULNÉRABLE — XSS via SVG (CWE-79)
// Un avatar uploadé au format SVG est enregistré et servi tel quel, sans
// sanitisation du contenu XML. Un SVG contenant <script> ou un gestionnaire
// d'événement (onload) s'exécute dans le contexte d'origine de l'application
// lorsqu'il est ouvert directement par le navigateur.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;

class AvatarController extends Controller
{
    public function upload(Request $request)
    {
        $file = $request->file('avatar');

        // Stockage direct du SVG uploadé, sans sanitisation du contenu XML.
        $path = $file->storeAs('avatars', $file->hashName(), 'public');

        return response()->json(['path' => $path]);
    }

    public function show(string $filename)
    {
        // Service du fichier SVG tel quel, affichage inline dans le navigateur.
        return response(Storage::disk('public')->get('avatars/' . $filename))
            ->header('Content-Type', 'image/svg+xml');
    }
}
