<?php

// Correction : utilisation de la Mailable/Mail de Laravel (basée sur
// Symfony Mailer), qui échappe correctement les en-têtes, combinée à un
// filtrage explicite des caractères de contrôle (\r, \n) sur les champs
// libres avant leur usage, en défense en profondeur.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Mail;

class ContactController extends Controller
{
    public function send(Request $request)
    {
        $validated = $request->validate([
            'name' => ['required', 'string', 'max:150'],
            'subject' => ['required', 'string', 'max:150'],
            'message' => ['required', 'string', 'max:5000'],
        ]);

        // Rejet défensif de tout caractère de contrôle résiduel dans les champs libres.
        $stripControlChars = fn (string $value): string => str_replace(["\r", "\n"], '', $value);

        $name = $stripControlChars($validated['name']);
        $subject = $stripControlChars($validated['subject']);

        // Mail::raw() via Symfony Mailer échappe automatiquement les en-têtes.
        Mail::raw($validated['message'], function ($mail) use ($name, $subject) {
            $mail->from('contact@example.com', $name)
                ->to('dest@example.com')
                ->subject($subject);
        });

        return response()->json(['message' => 'Message envoyé']);
    }
}
