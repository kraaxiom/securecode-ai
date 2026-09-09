<?php
/**
 * CORRIGÉ — Chargement des secrets depuis l'environnement / gestionnaire de secrets.
 *
 * CWE-798: Use of Hard-coded Credentials
 */

namespace App\Services;

class PaymentGatewayService
{
    private string $stripeSecretKey;
    private string $encryptionKey;

    public function __construct()
    {
        // CORRIGÉ : secrets lus depuis les variables d'environnement Laravel
        // (fichier .env non versionné, ou gestionnaire de secrets en prod).
        $stripeKey = config('services.stripe.secret');
        if (!$stripeKey) {
            throw new \RuntimeException('STRIPE_SECRET_KEY manquant dans la configuration');
        }
        $this->stripeSecretKey = $stripeKey;

        $encKey = env('PAYMENT_ENCRYPTION_KEY');
        if (!$encKey) {
            throw new \RuntimeException('PAYMENT_ENCRYPTION_KEY manquant dans l\'environnement');
        }
        $this->encryptionKey = $encKey;
    }

    public function charge(int $amountCents, string $customerToken): array
    {
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

// config/services.php (extrait) :
// 'stripe' => [
//     'secret' => env('STRIPE_SECRET_KEY'),
// ],
//
// .env (non versionné, listé dans .gitignore) :
// STRIPE_SECRET_KEY=changeme123
// PAYMENT_ENCRYPTION_KEY=changeme123
