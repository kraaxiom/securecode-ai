<?php

// Correction : le prompt est construit via les rôles structurés de l'API
// (system/user), jamais par concaténation de texte brut, ce qui maintient
// une frontière de confiance explicite entre instructions développeur et
// entrée utilisateur. Chaque appel d'outil déclenché par le modèle est
// validé contre un schéma strict avant exécution, le principe du moindre
// privilège est appliqué (seuls des outils à faible impact sont exposés
// par défaut) et les actions sensibles requièrent une confirmation
// explicite.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Services\LlmClient;
use App\Services\ToolSchemaValidator;
use Illuminate\Support\Facades\Log;

class AssistantController extends Controller
{
    private const SYSTEM_PROMPT = 'Tu es un assistant support client. Réponds poliment.';
    private const HIGH_IMPACT_TOOLS = ['delete_ticket'];

    public function __construct(
        private LlmClient $llm,
        private ToolSchemaValidator $validator,
    ) {
    }

    public function chat(Request $request)
    {
        $userInput = $request->input('message');

        $response = $this->llm->chat([
            ['role' => 'system', 'content' => self::SYSTEM_PROMPT],
            ['role' => 'user', 'content' => $userInput],
        ], ['tools' => ['send_email']]); // outils à faible impact uniquement, par défaut

        foreach ($response->toolCalls as $call) {
            $this->validator->validate($call->tool->name, $call->args);

            if (in_array($call->tool->name, self::HIGH_IMPACT_TOOLS, true)) {
                abort(403, 'Confirmation humaine requise pour cette action');
            }

            $call->tool->execute($call->args);
        }

        Log::info('chat_turn', ['input' => $userInput, 'response_id' => $response->id]);

        return response()->json(['reply' => $response->text]);
    }
}
