// Corrigé : Header Injection (CWE-113)
// Les caractères de contrôle CR/LF sont supprimés avant toute écriture
// d'en-tête, et le nom de fichier est en plus restreint à une liste blanche
// de caractères autorisés, empêchant toute injection d'en-tête ou de
// segment de chemin.
package com.example.security.header;

import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HeaderInjectionController {

    @GetMapping("/api/files/download")
    public void download(@RequestParam String filename, HttpServletResponse response) {
        // 1) Suppression stricte des caractères de contrôle CR/LF.
        String sanitized = filename.replaceAll("[\\r\\n]", "");

        // 2) Liste blanche : uniquement lettres, chiffres, point, tiret,
        //    underscore — aucun séparateur de chemin ni caractère spécial.
        if (!sanitized.matches("^[\\w.-]+$")) {
            response.setStatus(HttpServletResponse.SC_BAD_REQUEST);
            return;
        }

        response.setHeader("Content-Disposition", "attachment; filename=\"" + sanitized + "\"");
        response.setHeader("X-Requested-File", sanitized);
    }
}
