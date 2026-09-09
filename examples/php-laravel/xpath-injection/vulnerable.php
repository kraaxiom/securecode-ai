<?php

// Faille : les champs "user" et "pass" fournis par l'utilisateur sont
// concaténés directement dans une expression XPath utilisée pour
// authentifier l'utilisateur contre un document XML. Un attaquant peut
// injecter une apostrophe suivie d'une condition pour modifier la
// logique de sélection de nœuds (contournement d'authentification).

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DOMDocument;
use DOMXPath;

class XmlAuthController extends Controller
{
    public function login(Request $request)
    {
        $user = $request->input('user');
        $pass = $request->input('pass');

        $doc = new DOMDocument();
        $doc->load(storage_path('app/users.xml'));
        $xpath = new DOMXPath($doc);

        // Concaténation directe dans l'expression XPath.
        $nodes = $xpath->query("//user[username='{$user}' and password='{$pass}']");

        if ($nodes->length === 0) {
            return response()->json(['error' => 'Identifiants invalides'], 401);
        }

        return response()->json(['message' => 'Connecté']);
    }
}
