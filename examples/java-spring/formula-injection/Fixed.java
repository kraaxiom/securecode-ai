// Corrigé : Formula Injection / CSV Injection (CWE-1236)
// Toute valeur exportée dont le premier caractère est un déclencheur de
// formule ('=', '+', '-', '@', tabulation, retour chariot) est préfixée
// d'une apostrophe, ce qui force les tableurs à traiter la cellule comme du
// texte brut et empêche toute exécution de formule côté client.
package com.example.security.formula;

import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
public class FormulaInjectionController {

    public record ExportRow(String nom, String commentaire) {}

    private static final char[] TRIGGER_CHARS = {'=', '+', '-', '@', '\t', '\r'};

    // Préfixe la valeur d'une apostrophe si elle commence par un caractère
    // pouvant déclencher l'interprétation d'une formule par un tableur.
    private String neutraliserFormule(String valeur) {
        if (valeur == null || valeur.isEmpty()) {
            return valeur;
        }
        char premier = valeur.charAt(0);
        for (char trigger : TRIGGER_CHARS) {
            if (premier == trigger) {
                return "'" + valeur;
            }
        }
        return valeur;
    }

    @PostMapping(value = "/api/export/csv", produces = MediaType.TEXT_PLAIN_VALUE)
    public String exportCsv(@RequestBody List<ExportRow> rows) {
        StringBuilder csv = new StringBuilder("nom,commentaire\n");
        for (ExportRow row : rows) {
            csv.append(neutraliserFormule(row.nom()))
               .append(',')
               .append(neutraliserFormule(row.commentaire()))
               .append('\n');
        }
        return csv.toString();
    }
}
