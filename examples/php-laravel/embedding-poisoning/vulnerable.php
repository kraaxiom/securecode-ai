<?php

// Faille : le contenu soumis par un utilisateur est directement transformé
// en embedding puis indexé dans la base vectorielle utilisée pour la
// recherche sémantique, sans modération préalable, sans limite de fréquence
// par source et sans détection de densité anormale. Un attaquant peut ainsi
// fabriquer du contenu dont le vecteur est artificiellement proche de
// requêtes fréquentes pour polluer les résultats retournés à d'autres
// utilisateurs.

namespace App\Services;

class VectorIndexingService
{
    public function indexContent(string $userContent, string $sourceId, EmbeddingModel $embeddingModel, VectorStore $vectorStore): void
    {
        $vector = $embeddingModel->embed($userContent);

        // Indexation directe, sans modération ni contrôle de volume/densité.
        $vectorStore->upsert(uniqid('vec_', true), $vector, ['source' => $sourceId]);
    }
}
