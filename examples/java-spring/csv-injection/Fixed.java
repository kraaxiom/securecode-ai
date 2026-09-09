// Corrigé : CSV Injection / Formula Injection (CWE-1236)
// Toute cellule dont le premier caractère est `=`, `+`, `-`, `@`, une
// tabulation ou un retour chariot est préfixée par une apostrophe, ce qui
// neutralise son interprétation comme formule par le tableur.
package com.example.demo.controller;

import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/export")
public class CsvInjectionController {

    @GetMapping(value = "/comments.csv", produces = MediaType.TEXT_PLAIN_VALUE)
    public String exportCsv() {
        List<Map<String, String>> rows = List.of(
                Map.of("name", "Alice", "comment", "Bon produit"),
                Map.of("name", "Bob", "comment", "=cmd|'/c calc'!A1")
        );

        StringBuilder csv = new StringBuilder("name,comment\n");
        for (Map<String, String> row : rows) {
            csv.append(sanitizeCsvCell(row.get("name")))
               .append(",")
               .append(sanitizeCsvCell(row.get("comment")))
               .append("\n");
        }
        return csv.toString();
    }

    // Neutralise les caractères déclencheurs de formule dans un tableur.
    private String sanitizeCsvCell(String value) {
        if (value != null && value.matches("^[=+\\-@\\t\\r].*")) {
            return "'" + value;
        }
        return value;
    }
}
