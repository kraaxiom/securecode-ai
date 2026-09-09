<?php

// Correction : la clé API est résolue depuis un gestionnaire de secrets
// (jamais codée en dur ni versionnée), et les journaux applicatifs masquent
// systématiquement les en-têtes d'autorisation avant écriture. Le client ne
// détient jamais la clé : seul ce backend authentifié est autorisé à
// appeler le fournisseur LLM.

namespace App\Services;

use App\Support\SecretManager;
use Illuminate\Support\Facades\Http;
use Illuminate\Support\Facades\Log;

class LlmProxyController
{
    public function ask(string $question): array
    {
        $apiKey = SecretManager::resolve('llm_api_key'); // jamais en dur, jamais versionné

        $response = Http::withHeaders([
            'Authorization' => 'Bearer ' . $apiKey,
        ])->post('https://api.anthropic.com/v1/messages', [
            'model' => 'claude-3',
            'messages' => [['role' => 'user', 'content' => $question]],
        ]);

        Log::debug('llm_response_headers', $this->redactSensitiveHeaders($response->headers()));

        return $response->json();
    }

    private function redactSensitiveHeaders(array $headers): array
    {
        foreach (['authorization', 'x-api-key'] as $sensitive) {
            unset($headers[$sensitive]);
        }

        return $headers;
    }
}
