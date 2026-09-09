<?php

// Correction : le filtrage d'appartenance est appliqué directement dans la
// requête de base de données (where('user_id', ...)) au lieu d'être vérifié
// après coup. Un ID appartenant à un autre utilisateur renvoie 404, sans
// jamais exposer les données d'autrui.

namespace App\Http\Controllers\Api;

use App\Models\Order;
use App\Http\Controllers\Controller;

class OrderController extends Controller
{
    public function show($id)
    {
        $order = Order::where('id', $id)
            ->where('user_id', auth()->id())
            ->firstOrFail();

        return response()->json($order);
    }
}

// routes/api.php
// Route::get('/api/orders/{id}', [OrderController::class, 'show'])->middleware('auth:sanctum');
