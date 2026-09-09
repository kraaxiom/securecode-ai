<?php

// Correction : la clé secrète n'est plus présente dans le code. Elle est
// injectée au démarrage via la configuration Laravel, elle-même alimentée
// par une variable d'environnement issue d'un gestionnaire de secrets
// (AWS Secrets Manager / Vault), jamais stockée en clair dans un fichier
// versionné.

namespace App\Services;

use RuntimeException;
use Stripe\StripeClient;

class PaymentService
{
    private StripeClient $stripe;

    public function __construct()
    {
        $secretKey = config('services.stripe.secret');

        if (empty($secretKey)) {
            throw new RuntimeException('Clé Stripe manquante : vérifiez la configuration du gestionnaire de secrets.');
        }

        $this->stripe = new StripeClient($secretKey);
    }

    public function charge(int $amountCents, string $customerId): \Stripe\Charge
    {
        return $this->stripe->charges->create([
            'amount'   => $amountCents,
            'currency' => 'xof',
            'customer' => $customerId,
        ]);
    }
}
