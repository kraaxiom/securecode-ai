<?php

// Faille : l'API d'inférence d'un modèle fine-tuné sur des données clients
// sensibles renvoie des scores de confiance détaillés par classe et n'impose
// aucune limite au nombre de requêtes. Un attaquant peut ainsi interroger le
// modèle de façon répétée et méthodique pour reconstruire progressivement
// des informations mémorisées issues du jeu d'entraînement (inversion de
// modèle).

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Services\CustomerRiskModel;

class RiskScoreController extends Controller
{
    public function __construct(private CustomerRiskModel $model)
    {
    }

    public function score(Request $request)
    {
        $features = $request->input('features');

        $prediction = $this->model->predict($features);

        return response()->json([
            'class' => $prediction->class,
            'confidence_per_class' => $prediction->confidenceScores,
            'embedding' => $prediction->embedding,
        ]);
    }
}
