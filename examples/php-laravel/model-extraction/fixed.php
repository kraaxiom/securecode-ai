<?php

// Correction : application d'un quota strict par clé API via le middleware
// de limitation de débit de Laravel, détection des patterns d'usage
// anormaux (volume/régularité), et suppression des logits/probabilités
// bruts de la réponse — seule la sortie textuelle nécessaire au cas
// d'usage métier est renvoyée.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Services\LlmClient;
use App\Services\UsageMonitor;
use Illuminate\Support\Facades\Log;

class InferenceController extends Controller
{
    public function __construct(
        private LlmClient $llm,
        private UsageMonitor $usageMonitor,
    ) {
    }

    // Ce endpoint doit être déclaré derrière le middleware
    // throttle:api-inference (ex: 100 requêtes/heure par clé API).
    public function infer(Request $request)
    {
        $apiKey = $request->header('X-API-Key');

        if ($this->usageMonitor->detectsExtractionPattern($apiKey)) {
            Log::warning('suspected_model_extraction', ['api_key' => $apiKey]);
            abort(429, 'Usage anormal détecté');
        }

        $input = $request->input('prompt');
        $result = $this->llm->predict($input, ['return_logits' => false]);

        return response()->json([
            'output' => $result->text,
        ]);
    }
}
