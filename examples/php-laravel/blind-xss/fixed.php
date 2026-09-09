<?php

// CORRIGÉ — Blind XSS (CWE-79)
// Toute donnée provenant de l'extérieur du périmètre de confiance (ici un
// utilisateur externe non authentifié) est traitée comme non fiable, y
// compris dans une interface interne/admin. Encodage de sortie contextuel
// systématique via htmlspecialchars().

namespace App\Http\Controllers\Admin;

use App\Http\Controllers\Controller;
use App\Models\SupportTicket;

class SupportTicketController extends Controller
{
    public function show(int $id)
    {
        $ticket = SupportTicket::findOrFail($id);

        // Encodage contextuel systématique, même en back-office : la donnée
        // vient d'un canal externe (formulaire public), elle reste non fiable.
        return response(
            "<div class='ticket-message'>" .
            htmlspecialchars($ticket->message, ENT_QUOTES, 'UTF-8') .
            "</div>" .
            "<div class='ticket-agent'>" .
            htmlspecialchars($ticket->user_agent, ENT_QUOTES, 'UTF-8') .
            "</div>"
        );
    }
}
