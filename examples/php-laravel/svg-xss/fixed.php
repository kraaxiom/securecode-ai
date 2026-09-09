<?php

// CORRIGÉ — XSS via SVG (CWE-79)
// Le contenu SVG est sanitisé avant stockage via enshrined/svg-sanitize
// (suppression de <script>, des gestionnaires on*, des schémas javascript:),
// et le fichier est servi en téléchargement forcé (Content-Disposition:
// attachment) plutôt qu'en affichage inline, limitant l'impact résiduel.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use enshrined\svgSanitize\Sanitizer;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Storage;

class AvatarController extends Controller
{
    public function upload(Request $request)
    {
        $file = $request->file('avatar');

        $sanitizer = new Sanitizer();
        $clean = $sanitizer->sanitize(file_get_contents($file->getRealPath()));

        $filename = $file->hashName();
        Storage::disk('public')->put('avatars/' . $filename, $clean);

        return response()->json(['path' => 'avatars/' . $filename]);
    }

    public function show(string $filename)
    {
        // Téléchargement forcé plutôt qu'affichage inline, en défense supplémentaire.
        return response(Storage::disk('public')->get('avatars/' . $filename))
            ->header('Content-Type', 'image/svg+xml')
            ->header('Content-Disposition', 'attachment; filename="' . basename($filename) . '"');
    }
}
