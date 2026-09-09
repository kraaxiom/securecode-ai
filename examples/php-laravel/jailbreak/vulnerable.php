<?php

// Faille : le system prompt est la seule barrière de sécurité de l'agent
// conversationnel. Le message utilisateur et l'historique sont concaténés
// tels quels sans aucune couche de modération indépendante en entrée ou en
// sortie, et rien ne limite ni ne surveille le nombre de tours utilisés pour
// tenter d'éroder progressivement les restrictions du modèle.

namespace App\Services;

class ChatAssistantService
{
    private string $systemPrompt = 'Tu es un assistant utile et sûr.';

    public function chat(string $userMessage, array $history, LlmClient $llm): string
    {
        $messages = array_merge(
            [['role' => 'system', 'content' => $this->systemPrompt]],
            $history,
            [['role' => 'user', 'content' => $userMessage]]
        );

        // Aucune modération d'entrée, aucun filtre de sortie, aucune
        // surveillance du nombre de tours de la conversation.
        return $llm->chat($messages)->content;
    }
}
