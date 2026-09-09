<?php

// Faille : le champ "name" fourni par l'utilisateur est concaténé
// directement dans une chaîne représentant un document XML, sans
// échappement. Un attaquant peut y insérer des caractères spéciaux
// ("<", ">", "&") pour altérer la structure du document généré.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class ExportController extends Controller
{
    public function export(Request $request)
    {
        $name = $request->input('name');

        // Construction du XML par concaténation de chaînes, sans échappement.
        $xml = "<user><name>{$name}</name></user>";

        return response($xml, 200)->header('Content-Type', 'application/xml');
    }
}
