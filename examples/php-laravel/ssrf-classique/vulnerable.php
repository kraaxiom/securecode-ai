<?php

/**
 * VULNÉRABLE — Server-Side Request Forgery (SSRF) classique — CWE-918
 *
 * Le contrôleur récupère une URL fournie par l'utilisateur (paramètre
 * "url") et effectue une requête HTTP sortante sans aucune validation
 * de la destination : pas de whitelist de domaines, pas de résolution
 * DNS ni de vérification que l'IP obtenue n'est pas une adresse privée
 * / loopback / link-local, et les redirections HTTP sont suivies sans
 * revalidation. Un attaquant peut ainsi forcer le serveur à interroger
 * des ressources internes (services d'administration, réseau interne,
 * API cloud de métadonnées) normalement inaccessibles depuis l'extérieur.
 */

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;

class UrlPreviewController extends Controller
{
    public function preview(Request $request)
    {
        $url = $request->input('url');

        // Aucune validation de l'hôte/schéma/IP de destination.
        // Aucune limite sur les redirections suivies.
        $response = Http::get($url);

        return response()->json([
            'status'  => $response->status(),
            'content' => $response->body(),
        ]);
    }
}
