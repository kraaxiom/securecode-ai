<?php

// Faille : le commentaire soumis par l'utilisateur est écrit tel quel
// dans un fichier .shtml interprété par le serveur web (SSI/mod_include).
// Un attaquant peut injecter une directive SSI (ex: "<!--#exec ... -->")
// qui sera exécutée côté serveur lors du prochain affichage de la page.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class CommentController extends Controller
{
    public function store(Request $request)
    {
        $comment = $request->input('comment');

        // Écriture brute dans un fichier .shtml servi avec SSI activé, sans neutralisation.
        file_put_contents(
            public_path('comments.shtml'),
            "<p>{$comment}</p>\n",
            FILE_APPEND
        );

        return response()->json(['message' => 'Commentaire ajouté']);
    }
}
