<?php

/**
 * CORRIGÉ — Server-Side Request Forgery (SSRF) classique — CWE-918
 *
 * Correctifs appliqués :
 *  1. Whitelist stricte des hôtes/domaines autorisés pour les requêtes
 *     sortantes : seule une destination métier connue peut être appelée.
 *  2. Le schéma est restreint à https pour éviter file://, ftp://, gopher://.
 *  3. Résolution DNS explicite du domaine, puis validation de l'adresse IP
 *     obtenue afin d'exclure les plages privées, loopback et link-local
 *     (empêche l'atteinte du réseau interne même via un domaine public
 *     pointant vers une IP interne).
 *  4. Les redirections HTTP automatiques sont désactivées : une réponse
 *     3xx n'est jamais suivie aveuglément, ce qui empêche un serveur
 *     "de confiance" de rediriger la requête vers une cible interne.
 *  5. Timeout court pour limiter l'impact d'un scan de réseau interne.
 */

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Http;
use Illuminate\Validation\ValidationException;

class UrlPreviewController extends Controller
{
    /**
     * Liste blanche des hôtes autorisés pour la fonctionnalité d'aperçu.
     */
    private const ALLOWED_HOSTS = [
        'api.partenaire.example.com',
    ];

    public function preview(Request $request)
    {
        $rawUrl = $request->input('url');

        $parts = parse_url((string) $rawUrl);
        if (!$parts || ($parts['scheme'] ?? '') !== 'https'
            || !in_array($parts['host'] ?? '', self::ALLOWED_HOSTS, true)
        ) {
            throw ValidationException::withMessages([
                'url' => 'Destination non autorisée.',
            ]);
        }

        // Résolution DNS explicite puis validation de l'IP obtenue :
        // on rejette les plages privées, réservées et link-local
        // (ex: 127.0.0.0/8, 10.0.0.0/8, 169.254.0.0/16, ...).
        $ip = gethostbyname($parts['host']);
        if (filter_var(
            $ip,
            FILTER_VALIDATE_IP,
            FILTER_FLAG_NO_PRIV_RANGE | FILTER_FLAG_NO_RES_RANGE
        ) === false) {
            throw ValidationException::withMessages([
                'url' => 'Destination interdite.',
            ]);
        }

        // Redirections désactivées : chaque saut devrait être revalidé
        // séparément si la fonctionnalité en a réellement besoin.
        $response = Http::withOptions([
            'allow_redirects' => false,
            'timeout'         => 5,
        ])->get($rawUrl);

        return response()->json([
            'status'  => $response->status(),
            'content' => $response->body(),
        ]);
    }
}
