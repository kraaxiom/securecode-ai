<?php

// Faille : le pipeline de constitution du jeu de fine-tuning récupère et
// intègre les données de chaque source sans vérifier leur provenance, leur
// intégrité, ni détecter d'anomalies statistiques. Un attaquant contrôlant
// une des sources peut y injecter des exemples empoisonnés (biais, portes
// dérobées comportementales) qui seront ensuite appris par le modèle de
// façon persistante et difficile à détecter a posteriori.

namespace App\Services;

class FineTuningDatasetBuilder
{
    public function build(array $sources): array
    {
        $dataset = [];

        foreach ($sources as $source) {
            // Aucune vérification de provenance ni d'intégrité avant
            // intégration au jeu d'entraînement.
            $batch = $this->fetchAndParse($source['url']);
            $dataset = array_merge($dataset, $batch);
        }

        return $dataset;
    }

    private function fetchAndParse(string $url): array
    {
        $raw = file_get_contents($url);

        return json_decode($raw, true) ?? [];
    }
}
