<?php

// Correction : une policy centralisée (InvoicePolicy) applique le principe
// "deny by default" : l'accès n'est autorisé que si l'utilisateur est
// propriétaire de la facture ou administrateur. La règle est appliquée de
// façon systématique via le middleware 'can', pas seulement dans l'UI.

namespace App\Http\Controllers;

use App\Models\Invoice;

class InvoiceController extends Controller
{
    public function download($id)
    {
        $invoice = Invoice::findOrFail($id);

        $this->authorize('view', $invoice);

        return response()->download($invoice->file_path);
    }
}

// app/Policies/InvoicePolicy.php
// class InvoicePolicy
// {
//     public function view(User $user, Invoice $invoice): bool
//     {
//         return $invoice->user_id === $user->id || $user->role === 'admin';
//     }
// }

// routes/web.php
// Route::get('/invoices/{id}/download', [InvoiceController::class, 'download'])
//     ->middleware(['auth', 'can:view,invoice']);
