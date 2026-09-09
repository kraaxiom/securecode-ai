<?php

// Correction : le chargeur d'entités externes est désactivé au niveau
// global libxml, et le document est chargé avec LIBXML_NONET (aucun accès
// réseau) sans LIBXML_NOENT (pas de substitution d'entités). Cela empêche
// la résolution d'entités externes pointant vers des fichiers locaux ou
// des URL distantes.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DOMDocument;

class ImportController extends Controller
{
    public function import(Request $request)
    {
        $userSuppliedXml = $request->getContent();

        // Désactive le chargement d'entités externes pour tout le processus libxml.
        $previous = libxml_set_external_entity_loader(null);

        $doc = new DOMDocument();
        // LIBXML_NONET : aucun accès réseau. Ne jamais combiner avec LIBXML_NOENT
        // sur une entrée non fiable (substituerait les entités déclarées).
        $doc->loadXML($userSuppliedXml, LIBXML_NONET);

        libxml_set_external_entity_loader($previous);

        $title = $doc->getElementsByTagName('title')->item(0)?->nodeValue;

        return response()->json(['title' => $title]);
    }
}
