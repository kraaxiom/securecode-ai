<?php

// VULNÉRABLE — Stored XSS (CWE-79)
// Le commentaire persisté en base est réaffiché dans la vue Blade avec
// l'échappement automatique désactivé ({!! !!}), permettant à un contenu
// malveillant soumis une seule fois de s'exécuter chez tous les visiteurs
// consultant la page, sans lien piégé ni interaction supplémentaire.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Models\Comment;
use Illuminate\Http\Request;

class CommentController extends Controller
{
    public function store(Request $request, int $articleId)
    {
        Comment::create([
            'article_id' => $articleId,
            'body' => $request->input('body'),
        ]);

        return redirect()->route('articles.show', $articleId);
    }
}

// resources/views/articles/show.blade.php
// @foreach ($comments as $comment)
//     <div class="comment">{!! $comment->body !!}</div>
// @endforeach
