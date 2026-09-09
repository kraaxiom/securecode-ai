// Corrigé : Regular Expression Denial of Service - ReDoS (CWE-1333)
// La regex est réécrite sans quantificateurs imbriqués (complexité linéaire
// garantie), et une limite stricte de longueur est imposée sur l'entrée avant
// toute application du moteur de correspondance, éliminant le risque de
// backtracking catastrophique.
package com.example.security.redos;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.regex.Pattern;

@RestController
public class RegexDosController {

    // Limite de longueur RFC 5321 pour une adresse email.
    private static final int MAX_EMAIL_LENGTH = 254;

    // Regex non ambiguë : pas de quantificateur imbriqué, complexité linéaire.
    private static final Pattern EMAIL_PATTERN =
            Pattern.compile("^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$");

    @PostMapping("/api/email/validate")
    public boolean validate(@RequestParam String email) {
        // Limite de longueur imposée avant l'exécution de la regex.
        if (email == null || email.length() > MAX_EMAIL_LENGTH) {
            return false;
        }
        return EMAIL_PATTERN.matcher(email).matches();
    }
}
