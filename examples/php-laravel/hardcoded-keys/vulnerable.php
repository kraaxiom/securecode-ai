<?php
/**
 * VULNÉRABLE — Clés/secrets codés en dur dans le code source.
 *
 * Faille : une clé API, une clé de chiffrement ou un secret intégré
 * littéralement dans le code devient accessible à quiconque a accès au
 * dépôt, à l'historique Git ou au binaire compilé.
 *
 * CWE-798: Use of Hard-coded Credentials
 * OWASP A02:2021 - Cryptographic Failures
 */

namespace App\Services;

class PaymentGatewayService
{
    // VULNÉRABLE : clé API de paiement codée en dur (valeur illustrative).
    private string $stripeSecretKey = 'sk_live_EXAMPLE_NOT_A_REAL_KEY';

    // VULNÉRABLE : clé de chiffrement symétrique codée en dur.
    private string $encryptionKey = 'CLE_SECRETE_EXEMPLE';

    public function charge(int $amountCents, string $customerToken): array
    {
        // Utilisation directe du secret en dur pour appeler l'API distante.
        $client = new \GuzzleHttp\Client([
            'base_uri' => 'https://api.stripe.com',
            'headers' => ['Authorization' => 'Bearer ' . $this->stripeSecretKey],
        ]);

        return $client->post('/v1/charges', [
            'form_params' => [
                'amount' => $amountCents,
                'currency' => 'xof',
                'source' => $customerToken,
            ],
        ])->getBody()->getContents();
    }
}
