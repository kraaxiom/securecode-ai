<?php

// Correction : minimisation de la sortie (seule la classe prédite est
// renvoyée, sans scores de confiance détaillés ni embedding brut), ajout
// d'une limitation de débit stricte par clé API, et détection des patterns
// de requêtes évoquant une reconstruction progressive de données mémorisées.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Services\CustomerRiskModel;
use App\Services\UsageMonitor;
use Illuminate\Support\Facades\Log;

class RiskScoreController extends Controller
{
    public function __construct(
        private CustomerRiskModel $model,
        private UsageMonitor $usageMonitor,
    ) {
    }

    // Déclaré derrière le middleware throttle:api-risk-score
    // (ex: 50 requêtes/heure par clé API).
    public function score(Request $request)
    {
        $apiKey = $request->header('X-API-Key');

        if ($this->usageMonitor->detectsReconstructionPattern($apiKey)) {
            Log::warning('suspected_model_inversion', ['api_key' => $apiKey]);
            abort(429, 'Usage anormal détecté');
        }

        $features = $request->input('features');
        $prediction = $this->model->predict($features);

        return response()->json([
            'class' => $prediction->class,
        ]);
    }
}
