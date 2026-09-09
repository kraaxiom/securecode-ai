<?php

// Correction : le prompt système ne contient plus aucun secret en clair,
// seulement une référence indirecte vers un outil dont l'exécution résout
// le secret côté application. L'accès de l'agent au système de fichiers est
// restreint à une liste blanche de répertoires publics, et un filtrage de
// sortie détecte tout motif de secret avant de renvoyer une réponse du
// modèle.

namespace App\Services;

use Illuminate\Support\Facades\Log;

class SupportAgentService
{
    private const ALLOWED_CONFIG_DIRS = ['/app/public-config'];

    public function __construct(private SecretScanner $secretScanner, private HttpClient $httpClient)
    {
    }

    public function buildSystemPrompt(): string
    {
        return 'Tu es un assistant interne. Pour appeler le service de facturation, '
             . "utilise l'outil call_billing_service (le secret est géré par l'application, "
             . 'jamais exposé dans ce contexte).';
    }

    public function readConfigForAgent(string $path): string
    {
        $realPath = realpath($path);
        $isAllowed = $realPath !== false && collect(self::ALLOWED_CONFIG_DIRS)
            ->contains(fn ($dir) => str_starts_with($realPath, realpath($dir) ?: $dir));

        if (!$isAllowed) {
            Log::warning('blocked_sensitive_file_access', ['path' => $path]);
            abort(403, 'Accès refusé à ce chemin');
        }

        return file_get_contents($realPath);
    }

    public function callBillingService(array $args): mixed
    {
        $secret = SecretManager::resolve('billing_service');

        return $this->httpClient->post('/billing', $args, ['Authorization' => "Bearer {$secret}"]);
    }

    public function filterModelOutput(string $content): string
    {
        if ($this->secretScanner->containsSecret($content)) {
            Log::warning('blocked_secret_in_output');
            return '[réponse masquée : secret potentiel détecté]';
        }

        return $content;
    }
}
