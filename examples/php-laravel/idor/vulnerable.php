<?php

// Faille : Insecure Direct Object Reference (IDOR).
// La facture est récupérée uniquement via son ID transmis par le client,
// sans vérifier qu'elle appartient à l'utilisateur connecté. En modifiant
// l'ID dans l'URL, un utilisateur authentifié peut consulter les factures
// de n'importe quel autre compte.

namespace App\Http\Controllers;

use App\Models\Invoice;

class InvoiceController extends Controller
{
    public function show($id)
    {
        $invoice = Invoice::findOrFail($id);

        return view('invoices.show', compact('invoice'));
    }
}

// routes/web.php
// Route::get('/invoices/{id}', [InvoiceController::class, 'show'])->middleware('auth');
