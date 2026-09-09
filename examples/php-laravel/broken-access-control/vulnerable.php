<?php

// Faille : Broken Access Control (catégorie générale).
// Le téléchargement d'une facture n'applique aucune vérification d'autorisation
// propre à la ressource : seule l'authentification est requise. L'application
// s'appuie à tort sur le fait que le lien de téléchargement n'est affiché que
// pour le propriétaire dans l'interface, sans contrôle équivalent côté serveur.

namespace App\Http\Controllers;

use App\Models\Invoice;

class InvoiceController extends Controller
{
    public function download($id)
    {
        $invoice = Invoice::findOrFail($id);

        return response()->download($invoice->file_path);
    }
}

// routes/web.php
// Route::get('/invoices/{id}/download', [InvoiceController::class, 'download'])->middleware('auth');
