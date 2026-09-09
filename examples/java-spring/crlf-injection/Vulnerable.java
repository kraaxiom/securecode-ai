// Vulnérable : CRLF Injection (CWE-93)
// La valeur utilisateur est injectée directement dans un en-tête HTTP de
// redirection sans filtrage. Un attaquant peut y insérer des séquences
// \r\n pour injecter des en-têtes supplémentaires ou scinder la réponse
// HTTP (HTTP Response Splitting).
package com.example.demo.controller;

import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/redirect")
public class CrlfInjectionController {

    @GetMapping
    public void redirectTo(@RequestParam String next, HttpServletResponse response) {
        // Écriture directe de l'entrée utilisateur dans l'en-tête Location.
        response.setStatus(302);
        response.setHeader("Location", next);
    }
}
