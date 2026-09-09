<?php

// CORRIGÉ — Stored XSS (CWE-79)
// Le commentaire est désormais sanitisé via HTMLPurifier avant stockage (au
// cas où du HTML minimal serait toléré), et la vue conserve l'échappement
// automatique de Blade ({{ }}) à l'affichage, en double protection.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Models\Comment;
use HTMLPurifier;
use HTMLPurifier_Config;
use Illuminate\Http\Request;

class CommentController extends Controller
{
    public function store(Request $request, int $articleId)
    {
        $purifier = new HTMLPurifier(HTMLPurifier_Config::createDefault());
        $clean = $purifier->purify($request->input('body'));

        Comment::create([
            'article_id' => $articleId,
            'body' => $clean,
        ]);

        return redirect()->route('articles.show', $articleId);
    }
}

// resources/views/articles/show.blade.php
// @foreach ($comments as $comment)
//     <div class="comment">{{ $comment->body }}</div>
// @endforeach
