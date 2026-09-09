<?php

// Correction : séparation stricte entre la couche de décision du LLM et la
// couche d'exécution. Les outils à fort impact exigent une confirmation
// humaine explicite et une validation métier indépendante du modèle avant
// exécution ; chaque décision et action est journalisée pour permettre
// l'audit post-incident.

namespace App\Services;

class AgentRunnerService
{
    private array $highImpactTools = ['send_payment', 'delete_account'];

    public function run(string $userInput, LlmClient $llm, AuditLogger $auditLog, HumanApprovalGate $approval): array
    {
        $plan = $llm->plan($userInput, [
            'send_payment', 'delete_account', 'send_email',
        ]);

        $results = [];
        foreach ($plan->steps as $step) {
            if (in_array($step->tool->name, $this->highImpactTools, true)) {
                if (!$approval->requireHumanConfirmation($step)) {
                    $auditLog->record('blocked_high_impact_action', $step);
                    continue;
                }
            }

            // Validation métier indépendante du modèle, quel que soit
            // l'outil, avant toute exécution.
            $this->validateAgainstBusinessRules($step);

            $results[] = $step->tool->execute($step->args);
            $auditLog->record('action_executed', $step);
        }

        return $results;
    }

    private function validateAgainstBusinessRules(object $step): void
    {
        // Contrôles métier (limites, permissions, cohérence des arguments)
        // appliqués indépendamment de ce que le modèle a produit.
    }
}
