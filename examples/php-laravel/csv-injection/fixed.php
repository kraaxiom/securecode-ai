<?php

// Correction : chaque cellule est passée dans sanitizeCsvCell(), qui
// préfixe d'une apostrophe toute valeur commençant par un caractère
// déclencheur de formule ('=', '+', '-', '@', tabulation, retour chariot).
// Le tableur ouvre alors la cellule comme texte brut, pas comme formule.

namespace App\Http\Controllers;

use App\Models\Customer;
use Symfony\Component\HttpFoundation\StreamedResponse;

class CustomerExportController extends Controller
{
    private function sanitizeCsvCell(string $value): string
    {
        if (preg_match('/^[=+\-@\t\r]/', $value)) {
            return "'" . $value;
        }

        return $value;
    }

    public function export()
    {
        $customers = Customer::all(['name', 'comment']);

        $response = new StreamedResponse(function () use ($customers) {
            $handle = fopen('php://output', 'w');
            foreach ($customers as $customer) {
                fputcsv($handle, [
                    $this->sanitizeCsvCell($customer->name),
                    $this->sanitizeCsvCell($customer->comment),
                ]);
            }
            fclose($handle);
        });

        $response->headers->set('Content-Type', 'text/csv');
        $response->headers->set('Content-Disposition', 'attachment; filename="customers.csv"');

        return $response;
    }
}
