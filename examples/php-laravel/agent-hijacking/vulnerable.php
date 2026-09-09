<?php

// Faille : l'agent LLM planifie une suite d'étapes puis exécute directement
// les outils à fort impact (paiement, suppression de compte) renvoyés par le
// modèle, sans validation métier indépendante ni confirmation humaine. Si un
// contenu externe non fiable (email, document) a influencé la planification,
// l'attaquant peut faire exécuter des actions irréversibles à l'insu de
// l'utilisateur, avec les mêmes privilèges élevés que l'agent tout au long
// de la session.

namespace App\Services;

class AgentRunnerService
{
    private array $highImpactTools = ['send_payment', 'delete_account'];

    public function run(string $userInput, LlmClient $llm): array
    {
        $plan = $llm->plan($userInput, [
            'send_payment', 'delete_account', 'send_email',
        ]);

        $results = [];
        foreach ($plan->steps as $step) {
            // Exécution directe des outils, y compris ceux à fort impact,
            // sans validation applicative ni confirmation humaine.
            $results[] = $step->tool->execute($step->args);
        }

        return $results;
    }
}
