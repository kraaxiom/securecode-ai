<?php

// Correction : seules les URLs dont le domaine figure dans une liste blanche
// de sources de confiance sont crawlées, chaque document est scanné à la
// recherche de motifs d'instructions cachées avant indexation, et la
// provenance (URL source, date d'ingestion, niveau de confiance) est
// conservée avec chaque chunk pour permettre une traçabilité complète et un
// retrait rapide en cas de détection ultérieure d'un document malveillant.

namespace App\Services;

use Illuminate\Support\Facades\Log;

class RagIngestionService
{
    private const TRUSTED_SOURCES = [
        'docs.internal.acme.com',
        'verified-partner-wiki.acme.com',
    ];

    public function __construct(
        private VectorStore $vectorStore,
        private HttpCrawler $crawler,
        private ContentScanner $contentScanner,
    ) {
    }

    public function ingestDocuments(array $urls): void
    {
        foreach ($urls as $url) {
            $domain = parse_url($url, PHP_URL_HOST);

            if (!in_array($domain, self::TRUSTED_SOURCES, true)) {
                Log::warning('rejected_untrusted_source', ['url' => $url]);
                continue;
            }

            $doc = $this->crawler->fetch($url);

            if ($this->contentScanner->containsSuspiciousInstructions($doc)) {
                Log::warning('flagged_suspicious_document', ['url' => $url]);
                continue;
            }

            foreach ($this->splitIntoChunks($doc) as $chunk) {
                $this->vectorStore->upsert($this->embed($chunk), [
                    'text' => $chunk,
                    'source_url' => $url,
                    'ingested_at' => now()->toIso8601String(),
                    'trust_level' => 'verified',
                ]);
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
