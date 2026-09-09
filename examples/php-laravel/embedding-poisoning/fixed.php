<?php

// Correction : le contenu passe par une modération avant génération de
// l'embedding, une limite de débit par source freine les campagnes
// d'indexation automatisées, et tout vecteur présentant une densité
// anormale (proximité suspecte de requêtes fréquentes) est mis en
// quarantaine au lieu d'être indexé directement.

namespace App\Services;

class VectorIndexingService
{
    private const RATE_LIMIT_PER_SOURCE = 50; // vecteurs / heure

    public function indexContent(
        string $userContent,
        string $sourceId,
        EmbeddingModel $embeddingModel,
        VectorStore $vectorStore,
        ContentModeration $contentModeration,
        RateLimiter $rateLimiter,
        AuditLogger $auditLog
    ): void {
        if (!$contentModeration->isAllowed($userContent)) {
            $auditLog->record('rejected_content_moderation', $sourceId);
            return;
        }

        if ($rateLimiter->exceeded($sourceId, self::RATE_LIMIT_PER_SOURCE)) {
            $auditLog->record('rejected_rate_limit', $sourceId);
            return;
        }

        $vector = $embeddingModel->embed($userContent);

        if ($vectorStore->isAnomalousDensity($vector)) {
            $auditLog->record('flagged_anomalous_embedding', $sourceId);
            $vectorStore->quarantine($vector, ['source' => $sourceId]);
            return;
        }

        $vectorStore->upsert(uniqid('vec_', true), $vector, ['source' => $sourceId]);
    }
}
