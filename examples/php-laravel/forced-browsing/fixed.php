<?php

// Correction : un contrôle d'authentification et d'autorisation explicite est
// appliqué à la route, indépendamment de sa découvrabilité dans l'interface.
// L'accès nécessite désormais d'être connecté ET d'avoir la permission
// 'viewInternalReports', peu importe que le lien soit affiché ou non.

namespace App\Http\Controllers;

class ReportController extends Controller
{
    public function index()
    {
        $reports = \App\Models\Report::all();

        return view('reports.internal', compact('reports'));
    }
}

// routes/web.php
// Route::get('/internal-reports', [ReportController::class, 'index'])
//     ->middleware(['auth', 'can:viewInternalReports,App\Models\User']);
