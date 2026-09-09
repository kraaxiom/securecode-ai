<?php

// Faille : le document XML envoyé par l'utilisateur est chargé avec la
// configuration par défaut de DOMDocument, qui ne désactive pas le
// traitement des DTD et des entités externes. Un attaquant peut déclarer
// une entité externe pointant vers un fichier local, menant à une
// divulgation de fichiers sensibles ou une SSRF.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DOMDocument;

class ImportController extends Controller
{
    public function import(Request $request)
    {
        $userSuppliedXml = $request->getContent();

        $doc = new DOMDocument();
        // Aucune désactivation des entités externes / DTD.
        $doc->loadXML($userSuppliedXml);

        $title = $doc->getElementsByTagName('title')->item(0)?->nodeValue;

        return response()->json(['title' => $title]);
    }
}
