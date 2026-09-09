<?php

// Correction : le system prompt n'est plus la seule barrière. Une couche de
// modération indépendante du modèle principal analyse le message entrant et
// la réponse sortante, et le nombre de tours "suspects" par session est
// borné pour empêcher une érosion progressive des restrictions via des
// tentatives répétées. Aucun contenu de contournement n'est manipulé ici :
// uniquement la structure défensive applicative.

namespace App\Services;

class ChatAssistantService
{
    private const MAX_SUSPICIOUS_TURNS = 3;
    private const REFUSAL_MESSAGE = 'Je ne peux pas répondre à cette demande.';

    private string $systemPrompt = 'Tu es un assistant utile et sûr.';

    public function chat(
        string $userMessage,
        array $history,
        LlmClient $llm,
        SafetyClassifier $safetyClassifier,
        ConversationSession $session,
        AuditLogger $auditLog
    ): string {
        if ($safetyClassifier->flags($userMessage)) {
            $auditLog->record('blocked_input_flagged', $session->id);
            return self::REFUSAL_MESSAGE;
        }

        if ($session->suspiciousTurnCount() > self::MAX_SUSPICIOUS_TURNS) {
            $auditLog->record('session_flagged_repeated_attempts', $session->id);
            return self::REFUSAL_MESSAGE;
        }

        $messages = array_merge(
            [['role' => 'system', 'content' => $this->systemPrompt]],
            $history,
            [['role' => 'user', 'content' => $userMessage]]
        );

        $response = $llm->chat($messages);

        if ($safetyClassifier->flags($response->content)) {
            $auditLog->record('blocked_output_flagged', $session->id);
            return self::REFUSAL_MESSAGE;
        }

        return $response->content;
    }
}
