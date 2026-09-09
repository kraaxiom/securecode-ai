<?php

// Faille : le prompt système et l'entrée utilisateur sont concaténés dans
// une seule chaîne de texte brute avant d'être envoyés au modèle, sans
// séparation structurelle des rôles (system/user). Le modèle ne peut alors
// plus distinguer structurellement "instruction du développeur" et
// "contenu fourni par l'utilisateur", et la sortie du modèle est utilisée
// directement pour déclencher des appels d'outils sans validation
// applicative indépendante. Aucun payload d'exemple n'est fourni ici.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Services\LlmClient;

class AssistantController extends Controller
{
    private const SYSTEM_PROMPT = 'Tu es un assistant support client. Réponds poliment.';

    public function __construct(private LlmClient $llm)
    {
    }

    public function chat(Request $request)
    {
        $userInput = $request->input('message');

        // Concaténation brute : aucune frontière de confiance entre
        // instructions système et entrée utilisateur.
        $prompt = self::SYSTEM_PROMPT . "\n" . $userInput;

        $response = $this->llm->complete($prompt, ['tools' => ['delete_ticket', 'send_email']]);

        foreach ($response->toolCalls as $call) {
            $call->tool->execute($call->args); // exécution directe, sans validation
        }

        return response()->json(['reply' => $response->text]);
    }
}
