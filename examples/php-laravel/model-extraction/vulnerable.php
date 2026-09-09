<?php

// Faille : l'endpoint d'inférence est accessible sans aucune limitation de
// débit ni quota par clé API, et renvoie les probabilités brutes du modèle
// en plus du texte généré. Un attaquant peut ainsi interroger massivement
// et systématiquement le modèle pour en reconstruire un équivalent
// fonctionnel (extraction/distillation), les scores bruts facilitant la
// reconstitution des frontières de décision internes.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Services\LlmClient;

class InferenceController extends Controller
{
    public function __construct(private LlmClient $llm)
    {
    }

    public function infer(Request $request)
    {
        $input = $request->input('prompt');

        $result = $this->llm->predict($input, ['return_logits' => true]);

        return response()->json([
            'output' => $result->text,
            'logits' => $result->logits,
            'probabilities' => $result->probabilities,
        ]);
    }
}
