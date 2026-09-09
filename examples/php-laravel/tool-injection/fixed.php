<?php

// Correction : chaque appel d'outil est validé contre un schéma typé strict
// avant exécution, indépendamment du texte généré par le modèle. Seules les
// définitions d'outils provenant de sources explicitement listées comme
// fiables sont chargées, et une liste blanche contextuelle restreint les
// outils disponibles selon le niveau de confiance du contenu en cours de
// traitement. Chaque exécution est journalisée pour permettre l'audit.

namespace App\Services;

use Illuminate\Support\Facades\Log;

class AgentToolExecutor
{
    private const TRUSTED_TOOL_SOURCES = ['internal-registry'];

    private const ALLOWED_TOOLS_BY_TRUST = [
        'low' => ['search_docs'],
        'high' => ['search_docs', 'send_notification'],
    ];

    public function __construct(private ToolRegistry $registry, private ToolSchemaValidator $validator)
    {
    }

    public function executeToolCall(object $call, string $contextTrustLevel): mixed
    {
        $allowed = self::ALLOWED_TOOLS_BY_TRUST[$contextTrustLevel] ?? [];

        if (!in_array($call->name, $allowed, true)) {
            Log::warning('blocked_tool_not_in_allowlist', [
                'tool' => $call->name,
                'trust_level' => $contextTrustLevel,
            ]);
            abort(403, "Outil non autorisé dans ce contexte");
        }

        $validatedArgs = $this->validator->validate($call->name, $call->args);
        $tool = $this->registry->get($call->name);
        $result = $tool->execute($validatedArgs);

        Log::info('tool_executed', ['tool' => $call->name, 'args' => $validatedArgs]);

        return $result;
    }

    public function loadToolDefinitions(ToolSource $source): array
    {
        if (!in_array($source->id, self::TRUSTED_TOOL_SOURCES, true)) {
            Log::warning('rejected_untrusted_tool_source', ['source' => $source->id]);
            abort(403, 'Source d\'outils non approuvée');
        }

        return $source->load();
    }
}
