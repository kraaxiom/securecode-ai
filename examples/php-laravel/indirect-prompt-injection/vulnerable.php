<?php

// Faille : le contenu récupéré depuis une URL externe est inséré tel quel
// dans le message envoyé au modèle, sans marquage de provenance ni
// délimiteur distinguant "texte à résumer" et "instructions à exécuter".
// L'appel expose en plus des outils actifs (envoi d'email, exécution de
// code) pendant ce traitement, permettant à des instructions cachées dans
// la page web de détourner l'agent vers des actions non prévues.

namespace App\Services;

class UrlSummarizerService
{
    public function summarize(string $url, LlmClient $llm): string
    {
        $pageContent = file_get_contents($url);

        $response = $llm->chat([
            ['role' => 'user', 'content' => "Résume ce contenu : {$pageContent}"],
        ], [
            'tools' => ['send_email', 'execute_code'],
        ]);

        return $response->content;
    }
}
