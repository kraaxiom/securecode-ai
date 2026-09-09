<?php

// Correction : le document XML est construit via DOMDocument, qui
// échappe automatiquement le contenu des nœuds texte. La valeur
// utilisateur ne peut plus altérer la structure du document généré.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DOMDocument;

class ExportController extends Controller
{
    public function export(Request $request)
    {
        $name = $request->string('name')->toString();

        $doc = new DOMDocument('1.0', 'UTF-8');
        $userEl = $doc->createElement('user');
        $nameEl = $doc->createElement('name');
        $nameEl->appendChild($doc->createTextNode($name)); // échappement automatique
        $userEl->appendChild($nameEl);
        $doc->appendChild($userEl);

        return response($doc->saveXML(), 200)->header('Content-Type', 'application/xml');
    }
}
