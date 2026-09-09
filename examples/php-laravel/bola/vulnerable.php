<?php

// Faille : Broken Object Level Authorization (BOLA).
// L'endpoint d'API récupère une commande directement par son ID transmis par
// le client, sans vérifier qu'elle appartient bien à l'utilisateur authentifié.
// En changeant l'ID dans l'URL, n'importe quel utilisateur connecté peut
// consulter la commande d'un autre compte.

namespace App\Http\Controllers\Api;

use App\Models\Order;
use App\Http\Controllers\Controller;

class OrderController extends Controller
{
    public function show($id)
    {
        $order = Order::findOrFail($id);

        return response()->json($order);
    }
}

// routes/api.php
// Route::get('/api/orders/{id}', [OrderController::class, 'show'])->middleware('auth:sanctum');
