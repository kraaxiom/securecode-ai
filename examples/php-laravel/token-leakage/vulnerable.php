<?php

// Faille : la clé API du fournisseur LLM est codée en dur dans le code
// applicatif et journalisée sans masquage dans les logs de débogage, y
// compris l'en-tête d'autorisation complet. Toute personne ayant accès aux
// journaux ou au dépôt de code peut réutiliser cette clé au nom de
// l'application, entraînant un coût financier et un accès non autorisé aux
// données associées.

namespace App\Services;

class LlmProxyController
{
    private const API_KEY = 'sk-ant-api03-XXXXXXXXXXXXXXXXXXXXXXXX'; // clé en dur

    public function ask(string $question): array
    {
        $response = \Illuminate\Support\Facades\Http::withHeaders([
            'Authorization' => 'Bearer ' . self::API_KEY,
        ])->post('https://api.anthropic.com/v1/messages', [
            'model' => 'claude-3',
            'messages' => [['role' => 'user', 'content' => $question]],
        ]);

        error_log(print_r($response->headers(), true)); // journalise l'en-tête d'autorisation en clair

        return $response->json();
    }
}
