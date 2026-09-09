<?php

/**
 * VULNÉRABLE — Blind SSRF (SSRF aveugle) — CWE-918
 *
 * Ce job asynchrone valide une URL de webhook fournie par l'utilisateur
 * en effectuant une requête "fire and forget" : le résultat (statut,
 * contenu, erreur) n'est jamais renvoyé au client. Aucune whitelist,
 * aucune validation de l'IP résolue. Même sans retour visible, cette
 * requête sortante permet à un attaquant de scanner le réseau interne
 * ou de déclencher des callbacks hors bande (DNS/HTTP), rendant
 * l'exploitation plus difficile à détecter mais tout aussi dangereuse.
 */

namespace App\Jobs;

use Illuminate\Bus\Queueable;
use Illuminate\Contracts\Queue\ShouldQueue;
use Illuminate\Foundation\Bus\Dispatchable;
use Illuminate\Queue\InteractsWithQueue;
use Illuminate\Queue\SerializesModels;
use Illuminate\Support\Facades\Http;

class ValidateWebhookUrlJob implements ShouldQueue
{
    use Dispatchable, InteractsWithQueue, Queueable, SerializesModels;

    public function __construct(private readonly string $webhookUrl)
    {
    }

    public function handle(): void
    {
        // Requête sortante déclenchée par une entrée utilisateur,
        // sans whitelist ni validation d'IP, dont le résultat n'est
        // jamais exposé à l'appelant.
        try {
            Http::timeout(5)->get($this->webhookUrl);
        } catch (\Throwable $e) {
            // Erreur avalée silencieusement : aucune information
            // n'est renvoyée, mais la requête a bien été émise.
        }
    }
}
