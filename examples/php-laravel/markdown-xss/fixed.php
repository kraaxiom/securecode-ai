<?php

// CORRIGÉ — Markdown-based XSS (CWE-79)
// Le mode sûr de Parsedown est réactivé (désactive le HTML brut) et le HTML
// produit est en plus passé dans HTMLPurifier avant d'être transmis à la vue,
// en double protection contre tout contenu actif résiduel.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Models\WikiPage;
use HTMLPurifier;
use HTMLPurifier_Config;
use Parsedown;

class WikiPageController extends Controller
{
    public function show(int $id)
    {
        $page = WikiPage::findOrFail($id);

        $parsedown = new Parsedown();
        $parsedown->setSafeMode(true);

        $html = $parsedown->text($page->content);

        // Sanitisation additionnelle du HTML généré avant affichage.
        $purifier = new HTMLPurifier(HTMLPurifier_Config::createDefault());
        $cleanHtml = $purifier->purify($html);

        return view('wiki.show', ['renderedHtml' => $cleanHtml]);
    }
}
