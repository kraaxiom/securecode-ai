<?php

// Faille : les appels d'outils générés par le modèle (function calling)
// sont exécutés directement avec les arguments tels que produits par le
// texte du modèle, sans validation de schéma, sans vérifier que la
// définition de l'outil provient d'une source de confiance, et sans liste
// blanche contextuelle. Un contenu externe non fiable traité par l'agent
// (résultat d'outil précédent, document récupéré) peut ainsi influencer
// quel outil est appelé et avec quels paramètres, sans aucun garde-fou
// applicatif.

namespace App\Services;

class AgentToolExecutor
{
    public function __construct(private ToolRegistry $registry)
    {
    }

    public function executeToolCall(object $call): mixed
    {
        $tool = $this->registry->get($call->name);

        return $tool->execute($call->args); // aucune validation de schéma ni de contexte
    }

    public function loadToolDefinitions(ToolSource $source): array
    {
        return $source->load(); // chargé sans vérifier la confiance de la source
    }
}
