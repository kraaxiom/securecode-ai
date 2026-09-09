<?php

// Correction : seules les sources explicitement approuvées sont acceptées,
// leur intégrité est vérifiée via une empreinte signée, puis les données
// passent par une détection d'anomalies statistiques avant intégration.
// Le modèle résultant est ensuite évalué sur un jeu de référence fixe pour
// détecter une dérive de comportement suspecte après le cycle d'entraînement.

namespace App\Services;

class FineTuningDatasetBuilder
{
    private const ACCEPTABLE_DRIFT_THRESHOLD = 0.85;

    private array $trustedSources = ['internal-labeled-set', 'verified-partner-feed'];

    public function build(array $sources, AuditLogger $auditLog, AnomalyDetector $anomalyDetector): array
    {
        $dataset = [];

        foreach ($sources as $source) {
            if (!in_array($source['id'], $this->trustedSources, true)) {
                $auditLog->record('rejected_untrusted_source', $source['id']);
                continue;
            }

            $this->verifyIntegrity($source['url'], $source['signed_hash']);

            $batch = $this->fetchAndParse($source['url']);
            $batch = $anomalyDetector->filterStatisticalOutliers($batch);

            $dataset = array_merge($dataset, $batch);
        }

        return $dataset;
    }

    public function validateModelAfterTraining(object $model, array $goldenSet, ModelEvaluator $evaluator): void
    {
        $score = $evaluator->evaluate($model, $goldenSet);

        if ($score < self::ACCEPTABLE_DRIFT_THRESHOLD) {
            throw new \RuntimeException('Dérive de modèle détectée après entraînement');
        }
    }

    private function verifyIntegrity(string $url, string $expectedHash): void
    {
        $actualHash = hash('sha256', file_get_contents($url));

        if (!hash_equals($expectedHash, $actualHash)) {
            throw new \RuntimeException('Intégrité de la source invalide');
        }
    }

    private function fetchAndParse(string $url): array
    {
        $raw = file_get_contents($url);

        return json_decode($raw, true) ?? [];
    }
}
