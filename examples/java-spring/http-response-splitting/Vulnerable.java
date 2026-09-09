// Vulnérable : HTTP Response Splitting (CWE-113)
// L'URL de redirection fournie par l'utilisateur est écrite directement
// dans l'en-tête Location sans validation ni suppression des caractères
// CR/LF. Un attaquant peut injecter %0d%0a pour scinder la réponse HTTP et
// ajouter des en-têtes arbitraires, voire un second corps de réponse
// (cache poisoning, XSS réfléchi via en-tête falsifié).
package com.example.security.responsesplitting;

import jakarta.servlet.http.HttpServletResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.IOException;

@RestController
public class HttpResponseSplittingController {

    @GetMapping("/redirect")
    public void redirect(@RequestParam String next, HttpServletResponse response) throws IOException {
        // Aucune validation : "next" peut contenir des séquences décodées
        // \r\n qui scindent la réponse HTTP et injectent des en-têtes
        // ou un corps de réponse supplémentaires.
        response.setHeader("Location", next);
        response.setStatus(HttpServletResponse.SC_FOUND);
    }
}
