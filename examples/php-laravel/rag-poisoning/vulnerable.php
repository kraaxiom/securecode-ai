<?php

// Faille : le pipeline d'ingestion indexe directement tout document crawlé
// depuis une liste d'URLs fournie, sans vérifier que la source appartient à
// une liste blanche de confiance, sans scanner le contenu à la recherche
// d'instructions cachées, et sans conserver la provenance du document. Un
// contenu malveillant inséré dans une source externe sera plus tard récupéré
// et traité comme une information fiable par le modèle.

namespace App\Services;

class RagIngestionService
{
    public function __construct(private VectorStore $vectorStore, private HttpCrawler $crawler)
    {
    }

    public function ingestDocuments(array $urls): void
    {
        foreach ($urls as $url) {
            $doc = $this->crawler->fetch($url);
            $chunks = $this->splitIntoChunks($doc);

            foreach ($chunks as $chunk) {
                $this->vectorStore->upsert($this->embed($chunk), ['text' => $chunk]);
            }
        }
    }

    private function splitIntoChunks(string $doc): array
    {
        return str_split($doc, 500);
    }

    private function embed(string $chunk): array
    {
        return app(EmbeddingClient::class)->embed($chunk);
    }
}
