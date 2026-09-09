<?php

// Faille : les champs "name" et "comment" fournis par les utilisateurs sont
// écrits tels quels dans le fichier CSV exporté. Si une valeur commence par
// '=', '+', '-' ou '@', le tableur qui ouvre le fichier peut l'interpréter
// comme une formule active (injection CSV / formule).

namespace App\Http\Controllers;

use App\Models\Customer;
use Symfony\Component\HttpFoundation\StreamedResponse;

class CustomerExportController extends Controller
{
    public function export()
    {
        $customers = Customer::all(['name', 'comment']);

        $response = new StreamedResponse(function () use ($customers) {
            $handle = fopen('php://output', 'w');
            foreach ($customers as $customer) {
                fputcsv($handle, [$customer->name, $customer->comment]);
            }
            fclose($handle);
        });

        $response->headers->set('Content-Type', 'text/csv');
        $response->headers->set('Content-Disposition', 'attachment; filename="customers.csv"');

        return $response;
    }
}
