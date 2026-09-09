<?php

// Correction : suppression totale de l'évaluation dynamique de code.
// L'opération est choisie parmi une liste blanche de fonctions autorisées
// via une structure déclarative (match), et toute valeur non prévue est
// explicitement rejetée.

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use InvalidArgumentException;

class FormulaController extends Controller
{
    private const ALLOWED_OPERATIONS = ['add', 'sub', 'mul', 'div'];

    public function compute(Request $request)
    {
        $op = $request->input('op');
        $a = (float) $request->input('a');
        $b = (float) $request->input('b');

        if (!in_array($op, self::ALLOWED_OPERATIONS, true)) {
            throw new InvalidArgumentException('Opération non autorisée');
        }

        $result = match ($op) {
            'add' => $a + $b,
            'sub' => $a - $b,
            'mul' => $a * $b,
            'div' => $a / $b,
        };

        return response()->json(['result' => $result]);
    }
}
