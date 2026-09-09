<?php

// CORRIGÉ — Mutation-based XSS / mXSS (CWE-79)
// strip_tags() est remplacé par HTMLPurifier, une bibliothèque de
// sanitisation HTML activement maintenue qui tient compte des quirks de
// reparsing du navigateur, réduisant le risque de contournement par mutation.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Models\Article;
use HTMLPurifier;
use HTMLPurifier_Config;
use Illuminate\Http\Request;

class ArticleController extends Controller
{
    public function update(Request $request, int $id)
    {
        $article = Article::findOrFail($id);

        // Bibliothèque de sanitisation HTML maintenue, robuste face aux mutations.
        $purifier = new HTMLPurifier(HTMLPurifier_Config::createDefault());
        $clean = $purifier->purify($request->input('body'));

        $article->update(['body' => $clean]);

        return redirect()->route('articles.show', $article->id);
    }
}
