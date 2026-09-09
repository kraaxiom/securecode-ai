<?php

// Faille : Forced Browsing.
// La route affichant les rapports internes n'est liée depuis aucun menu de
// l'interface utilisateur, mais reste accessible sans aucun contrôle
// d'authentification ni d'autorisation si l'on connaît (ou devine) son URL.

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
// Route::get('/internal-reports', [ReportController::class, 'index']);
