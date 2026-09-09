<?php

// VULNÉRABLE — Markdown-based XSS (CWE-79)
// Le contenu Markdown d'un article de wiki, fourni par l'utilisateur, est
// converti en HTML via Parsedown avec le mode sûr désactivé (setSafeMode(false)),
// ce qui autorise le HTML brut (balises <script>, gestionnaires d'événements,
// liens javascript:) à traverser le rendu et à s'exécuter chez le lecteur.

namespace App\Http\Controllers;

use App\Http\Controllers\Controller;
use App\Models\WikiPage;
use Parsedown;

class WikiPageController extends Controller
{
    public function show(int $id)
    {
        $page = WikiPage::findOrFail($id);

        $parsedown = new Parsedown();
        $parsedown->setSafeMode(false);

        // HTML brut autorisé dans le Markdown utilisateur, aucune sanitisation.
        $html = $parsedown->text($page->content);

        return view('wiki.show', ['renderedHtml' => $html]);
    }
}
