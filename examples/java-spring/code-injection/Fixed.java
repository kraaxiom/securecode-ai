// Corrigé : Code Injection (CWE-94)
// Aucune évaluation dynamique de code. Les opérations autorisées sont
// définies dans une liste blanche déclarative, et l'entrée utilisateur ne
// sélectionne qu'un nom d'opération, jamais du code exécutable.
package com.example.demo.controller;

import org.springframework.web.bind.annotation.*;

import java.util.Map;
import java.util.function.BinaryOperator;

@RestController
@RequestMapping("/api/calc")
public class CodeInjectionController {

    // Liste blanche des opérations autorisées : aucune exécution de code dynamique.
    private static final Map<String, BinaryOperator<Double>> ALLOWED_OPERATIONS = Map.of(
            "add", (a, b) -> a + b,
            "sub", (a, b) -> a - b,
            "mul", (a, b) -> a * b,
            "div", (a, b) -> a / b
    );

    @GetMapping("/eval")
    public double eval(@RequestParam String op, @RequestParam double a, @RequestParam double b) {
        BinaryOperator<Double> operation = ALLOWED_OPERATIONS.get(op);
        if (operation == null) {
            throw new IllegalArgumentException("Opération non autorisée");
        }
        return operation.apply(a, b);
    }
}
