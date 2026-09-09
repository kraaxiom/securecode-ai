<?php

// Faille : les champs "name" et "subject" saisis par l'utilisateur sont
// concaténés directement dans les en-têtes d'un email envoyé via mail().
// Un attaquant peut injecter des retours chariot (\r\n) pour ajouter des
// en-têtes arbitraires (Cc, Bcc) ou un second corps de message.

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class ContactController extends Controller
{
    public function send(Request $request)
    {
        $name = $request->input('name');
        $subject = $request->input('subject');
        $body = $request->input('message');

        // Concaténation directe d'entrées utilisateur dans les en-têtes.
        $headers = "From: contact@example.com\r\nReply-To: $name\r\n";

        mail('dest@example.com', $subject, $body, $headers);

        return response()->json(['message' => 'Message envoyé']);
    }
}
