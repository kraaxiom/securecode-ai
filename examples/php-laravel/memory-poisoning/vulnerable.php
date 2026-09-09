<?php

// Faille : les "faits" extraits automatiquement de la réponse de l'agent
// sont écrits directement dans la mémoire persistante de l'utilisateur sans
// confirmation ni cloisonnement vérifié. Lors des sessions suivantes, ce
// contenu est réinjecté tel quel dans le prompt système comme contexte de
// confiance, permettant à une information fausse ou une instruction cachée
// introduite lors d'une session antérieure de compromettre durablement le
// comportement de l'agent.

namespace App\Services;

class AgentMemoryService
{
    public function processTurn(string $userId, string $agentResponse, LlmClient $llm, MemoryStore $memoryStore): void
    {
        $extractedFacts = $llm->extractFacts($agentResponse);

        // Écriture directe en mémoire persistante, sans confirmation de
        // l'utilisateur ni cloisonnement explicite par tenant.
        $memoryStore->append($userId, $extractedFacts);
    }

    public function nextSession(string $userId, LlmClient $llm, MemoryStore $memoryStore): string
    {
        $memory = $memoryStore->get($userId);

        return $llm->chat([
            ['role' => 'system', 'content' => "Contexte connu : {$memory}"],
        ])->content;
    }
}
