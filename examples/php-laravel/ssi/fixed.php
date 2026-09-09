<?php

// Correction : le commentaire est échappé en HTML puis toute séquence
// de directive SSI résiduelle ("<!--#") est neutralisée avant écriture.
// Idéalement, le contenu dynamique devrait être rendu via un moteur de
// templates Blade plutôt que via un fichier .shtml interprété par SSI.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class CommentController extends Controller
{
    public function store(Request $request)
    {
        $validated = $request->validate([
            'comment' => ['required', 'string', 'max:1000'],
        ]);

        $comment = htmlspecialchars($validated['comment'], ENT_QUOTES, 'UTF-8');
        // Neutralise toute séquence de directive SSI qui aurait survécu à l'échappement HTML.
        $comment = str_replace('<!--#', '&lt;!--#', $comment);

        file_put_contents(
            public_path('comments.shtml'),
            "<p>{$comment}</p>\n",
            FILE_APPEND
        );

        return response()->json(['message' => 'Commentaire ajouté']);
    }
}
