<?php

// Correction : les valeurs utilisateur sont échappées via DOMXPath::quote()
// avant insertion dans l'expression, ce qui neutralise les apostrophes et
// opérateurs XPath. XPath reste déconseillé comme mécanisme
// d'authentification ; ce correctif protège le pattern existant en
// attendant une migration vers une base avec hachage de mot de passe.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use DOMDocument;
use DOMXPath;

class XmlAuthController extends Controller
{
    public function login(Request $request)
    {
        $user = $request->string('user')->toString();
        $pass = $request->string('pass')->toString();

        $doc = new DOMDocument();
        $doc->load(storage_path('app/users.xml'));
        $xpath = new DOMXPath($doc);

        // Échappement dédié des valeurs insérées dans l'expression XPath.
        $safeUser = $xpath->quote($user);
        $safePass = $xpath->quote($pass);
        $nodes = $xpath->query("//user[username={$safeUser} and password={$safePass}]");

        if ($nodes->length === 0) {
            return response()->json(['error' => 'Identifiants invalides'], 401);
        }

        return response()->json(['message' => 'Connecté']);
    }
}
