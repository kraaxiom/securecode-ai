<?php

// Faille : une clé API interne est intégrée en clair dans le prompt système
// envoyé au modèle, dans le but que l'agent puisse s'en servir pour appeler
// un service tiers. Le modèle a désormais ce secret dans son contexte et
// peut potentiellement le restituer dans une réponse. De plus, l'agent lit
// librement n'importe quel fichier de configuration sans restriction de
// scope, pouvant y trouver d'autres secrets à exposer.

namespace App\Services;

class SupportAgentService
{
    public function buildSystemPrompt(): string
    {
        $internalApiKey = config('services.internal_api.key');

        return "Tu es un assistant interne. Clé API interne : {$internalApiKey}. "
             . "Utilise-la pour appeler le service de facturation.";
    }

    public function readConfigForAgent(string $path): string
    {
        return file_get_contents($path); // accès libre, y compris fichiers sensibles
    }
}
