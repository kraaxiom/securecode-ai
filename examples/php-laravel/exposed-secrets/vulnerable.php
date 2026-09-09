<?php

// Faille : la clé secrète Stripe est codée en dur directement dans le
// code source. Elle finit inévitablement versionnée dans Git et visible
// par quiconque a accès au dépôt (y compris son historique), permettant à
// un attaquant d'effectuer des opérations de paiement frauduleuses avec les
// privilèges de cette clé.

namespace App\Services;

use Stripe\StripeClient;

class PaymentService
{
    private StripeClient $stripe;

    public function __construct()
    {
        $this->stripe = new StripeClient('sk_live_EXAMPLE_NOT_A_REAL_KEY');
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
