<?php

// VULNÉRABLE — Mutation-based XSS / mXSS (CWE-79)
// Le contenu HTML riche fourni par un utilisateur (éditeur WYSIWYG) est
// "nettoyé" avec strip_tags() en autorisant une liste de balises, une
// technique obsolète et fragile face aux quirks de reparsing du navigateur :
// un fragment jugé inerte peut être réinterprété ("muté") différemment lors
// de son insertion réelle dans le DOM.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Models\Article;
use Illuminate\Http\Request;

class ArticleController extends Controller
{
    public function update(Request $request, int $id)
    {
        $article = Article::findOrFail($id);

        // Sanitisation naïve et non maintenue à jour face aux vecteurs de mutation.
        $clean = strip_tags($request->input('body'), '<b><i><a><p>');

        $article->update(['body' => $clean]);

        return redirect()->route('articles.show', $article->id);
    }
}
