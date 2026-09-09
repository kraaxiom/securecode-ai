<?php

// VULNÉRABLE — Blind XSS (CWE-79)
// Le champ "message" d'un ticket de support, soumis par un utilisateur externe
// non authentifié, est affiché tel quel dans le back-office admin. L'attaquant
// n'observe pas directement l'exécution (elle a lieu dans l'écran interne d'un
// agent support), d'où le terme "blind" : le payload reste latent jusqu'à
// consultation par un compte à privilèges élevés.

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use App\Models\SupportTicket;

class SupportTicketController extends Controller
{
    public function show(int $id)
    {
        $ticket = SupportTicket::findOrFail($id);

        // Affichage direct du message soumis par un utilisateur externe,
        // sans échappement, dans l'interface d'administration interne.
        return response(
            "<div class='ticket-message'>" . $ticket->message . "</div>" .
            "<div class='ticket-agent'>" . $ticket->user_agent . "</div>"
        );
    }
}
