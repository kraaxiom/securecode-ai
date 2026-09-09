// Vulnérable : Formula Injection / CSV Injection (CWE-1236)
// Les valeurs saisies par l'utilisateur sont écrites telles quelles dans un
// export CSV. Si une valeur commence par '=', '+', '-' ou '@', le tableur
// qui ouvrira le fichier (Excel, LibreOffice) l'interprétera comme une
// formule, permettant l'exécution de commandes ou l'exfiltration de données
// chez la victime qui ouvre l'export.
package com.example.security.formula;

import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class FormulaInjectionController {

    public record ExportRow(String nom, String commentaire) {}

    @PostMapping(value = "/api/export/csv", produces = MediaType.TEXT_PLAIN_VALUE)
    public String exportCsv(@RequestBody List<ExportRow> rows) {
        StringBuilder csv = new StringBuilder("nom,commentaire\n");
        for (ExportRow row : rows) {
            // Aucune neutralisation : une valeur du type
            // "=cmd|'/c calc'!A1" est écrite directement dans le fichier.
            csv.append(row.nom()).append(',').append(row.commentaire()).append('\n');
        }
        return csv.toString();
    }
}
