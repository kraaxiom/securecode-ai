// Vulnérable : Regular Expression Denial of Service - ReDoS (CWE-1333)
// La validation d'email utilise une regex à quantificateurs imbriqués
// ("(...)+)+") appliquée à une entrée de taille non bornée. Une chaîne
// pathologique (ex: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!") provoque
// une explosion combinatoire du backtracking et bloque le thread du serveur.
package com.example.security.redos;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.regex.Pattern;

@RestController
public class RegexDosController {

    // Quantificateurs imbriqués ambigus : complexité exponentielle sur échec.
    private static final Pattern EMAIL_PATTERN =
            Pattern.compile("^([a-zA-Z0-9]+)+@([a-zA-Z0-9]+)+\\.[a-zA-Z]{2,}$");

    @PostMapping("/api/email/validate")
    public boolean validate(@RequestParam String email) {
        // Aucune limite de longueur avant application de la regex.
        return EMAIL_PATTERN.matcher(email).matches();
    }
}
