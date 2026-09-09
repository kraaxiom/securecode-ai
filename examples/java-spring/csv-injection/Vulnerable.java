// Vulnérable : CSV Injection / Formula Injection (CWE-1236)
// Les valeurs utilisateur sont écrites telles quelles dans l'export CSV. Si
// une valeur commence par `=`, `+`, `-` ou `@`, le tableur du destinataire
// peut l'interpréter comme une formule (ex: exécution de commande via
// `=cmd|'/c calc'!A1`).
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
            // Écriture directe sans échappement des caractères déclencheurs de formule.
            csv.append(row.get("name")).append(",").append(row.get("comment")).append("\n");
        }
        return csv.toString();
    }
}
