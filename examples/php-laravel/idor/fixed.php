<?php

// Correction : la clause d'appartenance ('user_id') est intégrée directement
// dans la requête de récupération, en plus de l'ID demandé. Une facture
// appartenant à un autre utilisateur renvoie désormais une erreur 404 plutôt
// que d'exposer ses données.

namespace App\Http\Controllers;

use App\Models\Invoice;

class InvoiceController extends Controller
{
    public function show($id)
    {
        $invoice = Invoice::where('id', $id)
            ->where('user_id', auth()->id())
            ->firstOrFail();

        return view('invoices.show', compact('invoice'));
    }
}

// routes/web.php
// Route::get('/invoices/{id}', [InvoiceController::class, 'show'])->middleware('auth');
